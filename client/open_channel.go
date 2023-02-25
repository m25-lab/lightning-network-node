package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	signingTypes "github.com/cosmos/cosmos-sdk/types/tx/signing"
	"github.com/m25-lab/lightning-network-node/core_chain_sdk/account"
	"github.com/m25-lab/lightning-network-node/core_chain_sdk/channel"
	"github.com/m25-lab/lightning-network-node/core_chain_sdk/common"
	"github.com/m25-lab/lightning-network-node/database/models"
	"github.com/m25-lab/lightning-network-node/rpc/pb"
)

func (client *Client) OpenChannel(clientId string, accountPacked *AccountPacked) error {
	multisigAddr, multiSigPubkey, _ := account.NewAccount().CreateMulSigAccountFromTwoAccount(accountPacked.fromAccount.PublicKey(), accountPacked.toAccount.PublicKey(), 2)

	exchangeCommitmentMessage, err := client.Node.Repository.Message.FindOneByChannelID(context.Background(), accountPacked.fromAccount.AccAddress().String(), multisigAddr+":token:1")
	if err != nil {
		return err
	}

	if exchangeCommitmentMessage.Action != models.ExchangeCommitment {
		return errors.New("partner has not sent commitment yet")
	}

	var exchangeCommitmentData models.CreateCommitmentData
	err = json.Unmarshal([]byte(exchangeCommitmentMessage.Data), &exchangeCommitmentData)
	if err != nil {
		return err
	}

	//create l1 open channel message
	channelClient := channel.NewChannel(*client.ClientCtx)
	openChannelMsg := channelClient.CreateOpenChannelMsg(
		multisigAddr,
		accountPacked.fromAccount.AccAddress().String(),
		accountPacked.toAccount.AccAddress().String(),
		exchangeCommitmentData.CoinToCreator,
		exchangeCommitmentData.CoinToHtlc,
	)
	signOpenChannelMsg := channel.SignMsgRequest{
		Msg:      openChannelMsg,
		GasLimit: 21000,
		GasPrice: "1token",
	}

	strSig, err := channelClient.SignMultisigTxFromOneAccount(signOpenChannelMsg, accountPacked.fromAccount, multiSigPubkey)
	if err != nil {
		return err
	}

	strSigPayload, err := json.Marshal(models.OpenChannelData{
		StrSig: string(strSig),
	})
	if err != nil {
		return err
	}

	rpcClient := pb.NewMessageServiceClient(client.CreateConn(accountPacked.toEndpoint))

	response, err := rpcClient.SendMessage(context.Background(), &pb.SendMessageRequest{
		MessageId: "",
		ChannelID: multisigAddr + ":token:1",
		Action:    models.OpenChannel,
		Data:      string(strSigPayload),
		From:      accountPacked.fromAccount.AccAddress().String() + "@" + client.Node.Config.LNode.External,
		To:        accountPacked.toAccount.AccAddress().String() + "@" + accountPacked.toEndpoint,
	})
	if err != nil {
		return err
	}
	if response.ErrorCode != "" {
		return errors.New(response.ErrorCode + ": " + response.Response)
	}

	partnerSigPayload := models.OpenChannelData{}
	err = json.Unmarshal([]byte(response.Response), &partnerSigPayload)
	if err != nil {
		return err
	}
	partnerSig := partnerSigPayload.StrSig

	//@Todo check sigature

	signList := make([][]signingTypes.SignatureV2, 0)
	signByte1, err := common.SignatureJsonDecoder(client.ClientCtx.TxConfig, strSig)
	if err != nil {
		return err
	}

	signByte2, err := common.SignatureJsonDecoder(client.ClientCtx.TxConfig, partnerSig)
	if err != nil {
		return err
	}
	signList = append(signList, signByte2)
	signList = append(signList, signByte1)

	newTx := common.NewMultisigTxBuilder(*client.ClientCtx, nil, signOpenChannelMsg.GasLimit, signOpenChannelMsg.GasPrice, 0, 2)
	txBuilderMultiSign, err := newTx.BuildUnsignedTx(signOpenChannelMsg.Msg)
	if err != nil {
		return err
	}
	txJson, _ := common.TxBuilderJsonEncoder(client.ClientCtx.TxConfig, txBuilderMultiSign)
	fmt.Print(txJson)

	err = newTx.GenerateMultisig(txBuilderMultiSign, multiSigPubkey, signList)
	if err != nil {
		return err
	}
	txJson, err = common.TxBuilderJsonEncoder(client.ClientCtx.TxConfig, txBuilderMultiSign)
	if err != nil {
		return err
	}

	txByte, err := common.TxBuilderJsonDecoder(client.ClientCtx.TxConfig, txJson)
	if err != nil {
		return err
	}
	_, err = client.ClientCtx.BroadcastTx(txByte)
	if err != nil {
		return err
	}

	return nil
}
