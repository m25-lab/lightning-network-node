package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	cryptoTypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	signingTypes "github.com/cosmos/cosmos-sdk/types/tx/signing"
	channelTypes "github.com/m25-lab/channel/x/channel/types"
	"github.com/m25-lab/lightning-network-node/core_chain_sdk/account"
	"github.com/m25-lab/lightning-network-node/core_chain_sdk/channel"
	"github.com/m25-lab/lightning-network-node/core_chain_sdk/common"
	"github.com/m25-lab/lightning-network-node/database/models"
	"github.com/m25-lab/lightning-network-node/rpc/pb"
	"strings"
	"time"
)

func (client *Client) BuildAndBroadcastCommitment(clientId string, commitmentId string) error {
	ctx, cc := context.WithTimeout(context.Background(), 15*time.Second)
	defer cc()
	fromAccount, err := client.CurrentAccount(clientId)
	if err != nil {
		return err
	}

	commitMesssage, err := client.Node.Repository.Message.FindOneById(ctx, fromAccount.AccAddress().String(), commitmentId)
	if err != nil {
		return err
	}

	payload := models.CreateCommitmentData{}
	err = json.Unmarshal([]byte(commitMesssage.Data), &payload)
	if err != nil {
		return err
	}

	existedWhitelist, err := client.Node.Repository.Whitelist.FindOneByPartnerAddress(context.Background(), fromAccount.AccAddress().String(), commitMesssage.Users[1])
	if err != nil {
		return err
	}
	toAccount := account.NewPKAccount(existedWhitelist.PartnerPubkey)
	//broadcast
	channelClient := channel.NewChannel(*client.ClientCtx)
	commitmentMsg := channelClient.CreateCommitmentMsg(
		payload.Creator,
		payload.ToTimelockAddr,
		payload.CoinToCreator,
		payload.ToHashlockAddr,
		payload.CoinToHtlc,
		payload.Hashcode,
	)
	fmt.Println(commitmentMsg)

	signCommitmentMsg := channel.SignMsgRequest{
		Msg:      commitmentMsg,
		GasLimit: 100000,
		GasPrice: "0token",
	}
	//fmt.Println(signCommitmentMsg)
	_, multiSigPubkey, _ := account.NewAccount().CreateMulSigAccountFromTwoAccount(fromAccount.PublicKey(), toAccount.PublicKey(), 2)
	txByte, err := client.BuildMultisigMsgReadyForBroadcast(client, multiSigPubkey, payload.OwnSignature, payload.PartnerSignature, signCommitmentMsg)
	if err != nil {
		return err
	}
	broadcastResponse, err := client.ClientCtx.BroadcastTx(txByte)
	if err != nil {
		return nil
	}
	if broadcastResponse.RawLog != "[]" {
		return errors.New(broadcastResponse.RawLog)
	}
	fmt.Println("\n broadcast commitment response: ", broadcastResponse)
	//update isReplied
	commitMesssage.IsReplied = true
	err = client.Node.Repository.Message.Update(context.Background(), commitMesssage.ID, commitMesssage)
	if err != nil {
		return err
	}

	// rpc to partner
	rpcClient := pb.NewMessageServiceClient(client.CreateConn(strings.Split(existedWhitelist.PartnerAddress, "@")[1]))
	response, err := rpcClient.BroadcastNoti(context.Background(), &pb.BroadcastNotiMessage{
		From:     fromAccount.AccAddress().String(),
		To:       existedWhitelist.PartnerAddress,
		HashCode: payload.Hashcode,
		Index:    fmt.Sprintf("%s:%s", payload.Creator, payload.Hashcode),
	})
	if err != nil {
		return err
	}
	if response.ErrorCode != "" {
		return errors.New(response.Response)
	}
	return nil
}

func (client1 *Client) BuildMultisigMsgReadyForBroadcast(client *Client, multiSigPubkey cryptoTypes.PubKey, sig1, sig2 string, msgRequest channel.SignMsgRequest) ([]byte, error) {
	fmt.Println("sig1: ", sig1)
	fmt.Println("sig2: ", sig2)
	signList := make([][]signingTypes.SignatureV2, 0)
	signByte1, err := common.SignatureJsonDecoder(client.ClientCtx.TxConfig, sig1)
	if err != nil {
		return nil, err
	}

	signByte2, err := common.SignatureJsonDecoder(client.ClientCtx.TxConfig, sig2)
	if err != nil {
		return nil, err
	}
	signList = append(signList, signByte1)
	signList = append(signList, signByte2)
	fmt.Println("signList: ", signList)

	newTx := common.NewMultisigTxBuilder(*client.ClientCtx, nil, msgRequest.GasLimit, msgRequest.GasPrice, 0, 2)
	txBuilderMultiSign, err := newTx.BuildUnsignedTx(msgRequest.Msg)
	if err != nil {
		return nil, err
	}
	//fmt.Println("Toi day roi ne")
	err = newTx.GenerateMultisig(txBuilderMultiSign, multiSigPubkey, uint32(118), signList)
	if err != nil {
		fmt.Println("GenerateMultisig: ", err.Error())
		return nil, err
	}
	txJson, err := common.TxBuilderJsonEncoder(client.ClientCtx.TxConfig, txBuilderMultiSign)
	if err != nil {
		return nil, err
	}
	fmt.Println("txJson: ", txJson)

	txByte, err := common.TxBuilderJsonDecoder(client.ClientCtx.TxConfig, txJson)
	if err != nil {
		return nil, err
	}
	fmt.Println("txByte: ", txByte)
	return txByte, nil
}

func (client *Client) BuildAndBroadcastWithdrawTimelock(clientId string, commitmentId string) error {
	ctx, cc := context.WithTimeout(context.Background(), 15*time.Second)
	defer cc()
	fromAccount, err := client.CurrentAccount(clientId)
	if err != nil {
		return err
	}

	commitMesssage, err := client.Node.Repository.Message.FindOneById(ctx, fromAccount.AccAddress().String(), commitmentId)
	if err != nil {
		return err
	}

	payload := models.CreateCommitmentData{}
	err = json.Unmarshal([]byte(commitMesssage.Data), &payload)
	if err != nil {
		return err
	}

	msg := channelTypes.MsgWithdrawTimelock{
		Creator: fromAccount.AccAddress().String(),
		To:      fromAccount.AccAddress().String(),
		Index:   fmt.Sprintf("%v:%v", payload.Creator, payload.Hashcode),
	}
	withdrawRequest := channel.SignMsgRequest{
		Msg:      &msg,
		GasLimit: 100000,
		GasPrice: "0token",
	}

	fmt.Println("BuildAndBroadcastWithdrawTimelock msg: ", msg)

	_, _, err = BroadcastTx(client.ClientCtx, fromAccount, withdrawRequest)

	return err
}

func BroadcastTx(client *client.Context, account *account.PrivateKeySerialized, request channel.SignMsgRequest) (*sdk.TxResponse, string, error) {

	newTx := common.NewTx(
		*client,
		account,
		request.GasLimit,
		request.GasPrice,
	)

	txBuilder, err := newTx.BuildUnsignedTx(request.Msg)
	if err != nil {
		return nil, "", err
	}

	err = newTx.SignTx(txBuilder)
	if err != nil {
		return nil, "", err
	}

	txJson, err := common.TxBuilderJsonEncoder(client.TxConfig, txBuilder)
	if err != nil {
		panic(err)
	}

	fmt.Println("Tx rawData", string(txJson))

	txByte, err := common.TxBuilderJsonDecoder(client.TxConfig, txJson)
	if err != nil {
		panic(err)
	}

	txHash := common.TxHash(txByte)
	fmt.Println("txHash", txHash)

	//fmt.Println(ethCommon.BytesToHash(txByte).String())

	res, err := client.BroadcastTxCommit(txByte)
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
	if strings.Contains(res.RawLog, "Error") ||
		strings.Contains(res.RawLog, "error") ||
		strings.Contains(res.RawLog, "failed") {
		err = errors.New(res.RawLog)
	}

	return res, txHash, err
}
