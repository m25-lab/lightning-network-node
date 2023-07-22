package client

import (
	"context"
	"encoding/json"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"strconv"

	"github.com/m25-lab/lightning-network-node/rpc/pb"

	"github.com/m25-lab/lightning-network-node/core_chain_sdk/account"
	"github.com/m25-lab/lightning-network-node/core_chain_sdk/channel"
	"github.com/m25-lab/lightning-network-node/database/models"
)

func (client *Client) ExchangeFwdCommitment(clientId string, accountPacked *AccountPacked, fromAmount int64, toAmount int64, transferAmount int64, fwdDest string, hashcodeDest *string, hops int64) (*primitive.ObjectID, error) {
	log.Println("ExchangeFwdCommitment... run")
	multisigAddr, multiSigPubkey, _ := account.NewAccount().CreateMulSigAccountFromTwoAccount(accountPacked.fromAccount.PublicKey(), accountPacked.toAccount.PublicKey(), 2)

	//get partner hashcode
	exchangeHashcodeMessage, err := client.Node.Repository.Message.FindOneByChannelID(context.Background(), accountPacked.fromAccount.AccAddress().String(), multisigAddr+":token:1")
	if err != nil {
		log.Println("FindOneByChannelID...")
		return nil, err
	}

	if exchangeHashcodeMessage.Action != models.ExchangeHashcode {
		return nil, errors.New("partner has not sent hashcode yet")
	}

	var exchangeHashcodeData models.ExchangeHashcodeData
	err = json.Unmarshal([]byte(exchangeHashcodeMessage.Data), &exchangeHashcodeData)
	if err != nil {
		return nil, err
	}

	//create l1 sender commitment
	blockheight, err := client.ClientCtx.Client.Status(context.Background())
	if err != nil {
		println("BlockHeight-Status: ", err.Error())
	}
	senderLock := strconv.FormatInt(blockheight.SyncInfo.LatestBlockHeight+1000*(hops-1), 10)

	channelClient := channel.NewChannel(*client.ClientCtx)
	senderCommitmentMsg := channelClient.CreateSenderCommitmentMsg(
		multisigAddr,
		accountPacked.fromAccount.AccAddress().String(),
		fromAmount,
		toAmount,
		transferAmount,
		exchangeHashcodeData.PartnerHashcode,
		*hashcodeDest,
		senderLock,
	)
	//create l1 sign message
	signCommitmentMsg := channel.SignMsgRequest{
		Msg:      senderCommitmentMsg,
		GasLimit: 200000,
		GasPrice: "0token",
	}

	//sign l1 sender commitment
	strSig, err := channelClient.SignMultisigTxFromOneAccount(signCommitmentMsg, accountPacked.fromAccount, multiSigPubkey, false)
	if err != nil {
		log.Println("SignMultisigTxFromOneAccount...")
		return nil, err
	}

	//create ln message
	msg := models.SenderCommitment{
		Creator:          senderCommitmentMsg.Creator,
		From:             senderCommitmentMsg.From,
		Channelid:        senderCommitmentMsg.ChannelID,
		CoinToSender:     senderCommitmentMsg.CoinToSender.Amount.Int64(),
		CoinToHTLC:       senderCommitmentMsg.CoinToHtlc.Amount.Int64(),
		HashcodeHTLC:     senderCommitmentMsg.HashcodeHtlc,
		TimelockHTLC:     senderCommitmentMsg.TimelockHtlc,
		CoinTransfer:     senderCommitmentMsg.CoinTransfer.Amount.Int64(),
		HashcodeDest:     senderCommitmentMsg.HashcodeDest,
		TimelockReceiver: senderCommitmentMsg.TimelockReceiver,
		Multisig:         senderCommitmentMsg.Multisig,
		Hops:             hops,
		TimelockSender:   senderCommitmentMsg.TimelockSender,
	}

	partnerCommitmentPayload, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}
	//send rpc (input: sendercommit with sig; output: receivercommit)

	rpcClient := pb.NewRoutingServiceClient(client.CreateConn(accountPacked.toEndpoint))
	response, err := rpcClient.ProcessFwdMessage(context.Background(), &pb.FwdMessage{
		Action:       models.SenderCommit,
		Data:         string(partnerCommitmentPayload),
		From:         accountPacked.fromAccount.AccAddress().String() + "@" + client.Node.Config.LNode.External,
		To:           accountPacked.toAccount.AccAddress().String() + "@" + accountPacked.toEndpoint,
		HashcodeDest: msg.HashcodeDest,
		Sig:          strSig,
		Dest:         fwdDest,
	})
	if err != nil {
		log.Println("ProcessFwdMessage...", err.Error())
		return nil, err
	}

	if response.ErrorCode != "" {
		return nil, errors.New(response.ErrorCode + ":" + response.Response)
	}
	myCommitmentPayload := models.ReceiverCommitment{}
	err = json.Unmarshal([]byte(response.Response), &myCommitmentPayload)
	if err != nil {
		return nil, err
	}

	//Build and sign receiver commit
	receiverCMsg := channelClient.CreateReceiverCommitmentMsg(
		multisigAddr,
		accountPacked.toAccount.AccAddress().String(),
		myCommitmentPayload.CoinToReceiver,
		myCommitmentPayload.CoinToHTLC,
		myCommitmentPayload.CoinTransfer,
		myCommitmentPayload.HashcodeHTLC,
		myCommitmentPayload.HashcodeDest,
		myCommitmentPayload.TimelockSender,
	)

	signReceiverCommitmentMsg := channel.SignMsgRequest{
		Msg:      receiverCMsg,
		GasLimit: 200000,
		GasPrice: "0token",
	}

	strSigReceiver, err := channelClient.SignMultisigTxFromOneAccount(signReceiverCommitmentMsg, accountPacked.fromAccount, multiSigPubkey, false)
	if err != nil {
		log.Println("SignMultisigTxFromOneAccount...")
		return nil, err
	}
	messageId := primitive.NewObjectID()
	err = client.Node.Repository.FwdCommitment.InsertFwdMessage(context.Background(), &models.FwdMessage{
		ID:           messageId,
		Action:       models.ReceiverCommit,
		PartnerSig:   response.PartnerSig,
		OwnSig:       strSigReceiver,
		Data:         response.Response,
		From:         accountPacked.toAccount.AccAddress().String() + "@" + accountPacked.toEndpoint,
		To:           accountPacked.fromAccount.AccAddress().String() + "@" + client.Node.Config.LNode.External,
		HashcodeDest: myCommitmentPayload.HashcodeDest,
	})
	if err != nil {
		log.Println("InsertFwdMessage...")
		return nil, err
	}

	return &messageId, nil
}
