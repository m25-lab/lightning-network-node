package client

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/m25-lab/lightning-network-node/core_chain_sdk/account"
	"github.com/m25-lab/lightning-network-node/core_chain_sdk/channel"
	"github.com/m25-lab/lightning-network-node/database/models"
	"github.com/m25-lab/lightning-network-node/rpc/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (client *Client) ExchangeCommitment(clientId string, accountPacked *AccountPacked, amount int64) (*models.Message, error) {
	multisigAddr, multiSigPubkey, _ := account.NewAccount().CreateMulSigAccountFromTwoAccount(accountPacked.fromAccount.PublicKey(), accountPacked.toAccount.PublicKey(), 2)

	//get partner hashcode
	exchangeHashcodeMessage, err := client.Node.Repository.Message.FindOneByChannelID(context.Background(), accountPacked.fromAccount.AccAddress().String(), multisigAddr+":token:1")
	if err != nil {
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

	//create l1 commitment
	channelClient := channel.NewChannel(*client.ClientCtx)
	commitmentMsg := channelClient.CreateCommitmentMsg(
		multisigAddr,
		accountPacked.toAccount.AccAddress().String(),
		0,
		accountPacked.fromAccount.AccAddress().String(),
		amount,
		exchangeHashcodeData.PartnerHashcode,
	)

	//create l1 sign message
	signCommitmentMsg := channel.SignMsgRequest{
		Msg:      commitmentMsg,
		GasLimit: 21000,
		GasPrice: "1token",
	}

	//sign l1 commitment
	strSig, err := channelClient.SignMultisigTxFromOneAccount(signCommitmentMsg, accountPacked.fromAccount, multiSigPubkey)
	if err != nil {
		return nil, err
	}

	//create ln message
	myCommitmentPayload, err := json.Marshal(models.CreateCommitmentData{
		Creator:          commitmentMsg.Creator,
		ChannelID:        commitmentMsg.ChannelID,
		From:             commitmentMsg.From,
		Timelock:         commitmentMsg.Timelock,
		ToTimelockAddr:   commitmentMsg.ToTimelockAddr,
		ToHashlockAddr:   commitmentMsg.ToHashlockAddr,
		CoinToCreator:    commitmentMsg.CoinToCreator.Amount.Int64(),
		CoinToHtlc:       commitmentMsg.CoinToHtlc.Amount.Int64(),
		Hashcode:         commitmentMsg.Hashcode,
		PartnerSignature: strSig,
	})
	if err != nil {
		return nil, err
	}

	messageId := primitive.NewObjectID()
	savedMessage := models.Message{
		ID:         messageId,
		OriginalID: messageId,
		ChannelID:  commitmentMsg.ChannelID,
		Action:     models.ExchangeCommitment,
		Owner:      accountPacked.fromAccount.AccAddress().String(),
		Users: []string{
			accountPacked.fromAccount.AccAddress().String() + "@" + client.Node.Config.LNode.External,
			accountPacked.toAccount.AccAddress().String() + "@" + accountPacked.toEndpoint,
		},
		IsReplied: false,
	}

	rpcClient := pb.NewMessageServiceClient(client.CreateConn(accountPacked.toEndpoint))
	reponse, err := rpcClient.SendMessage(context.Background(), &pb.SendMessageRequest{
		MessageId: messageId.Hex(),
		ChannelID: savedMessage.ChannelID,
		Action:    models.ExchangeCommitment,
		Data:      string(myCommitmentPayload),
		From:      accountPacked.fromAccount.AccAddress().String() + "@" + client.Node.Config.LNode.External,
		To:        accountPacked.toAccount.AccAddress().String() + "@" + accountPacked.toEndpoint,
	})
	if err != nil {
		return nil, err
	}
	if reponse.ErrorCode != "" {
		return nil, errors.New(reponse.ErrorCode)
	}

	partnerCommitmentPayload := models.CreateCommitmentData{}
	err = json.Unmarshal([]byte(reponse.Response), &partnerCommitmentPayload)
	if err != nil {
		return nil, err
	}

	//check partner commitment
	if (partnerCommitmentPayload.Creator != commitmentMsg.Creator) ||
		(partnerCommitmentPayload.ChannelID != commitmentMsg.ChannelID) ||
		(partnerCommitmentPayload.From != commitmentMsg.From) ||
		(partnerCommitmentPayload.Timelock != commitmentMsg.Timelock) ||
		(partnerCommitmentPayload.ToTimelockAddr != commitmentMsg.ToHashlockAddr) ||
		(partnerCommitmentPayload.ToHashlockAddr != commitmentMsg.ToTimelockAddr) ||
		(partnerCommitmentPayload.CoinToCreator != commitmentMsg.CoinToHtlc.Amount.Int64()) ||
		(partnerCommitmentPayload.CoinToHtlc != commitmentMsg.CoinToCreator.Amount.Int64()) ||
		(partnerCommitmentPayload.Hashcode != exchangeHashcodeData.MyHashcode) {
		return nil, errors.New("partner commitment is not match")
	}

	//check partner signature
	//@TODO

	savedMessage.Data = string(reponse.Response)
	err = client.Node.Repository.Message.InsertOne(context.Background(), &savedMessage)
	if err != nil {
		return nil, err
	}

	return &savedMessage, nil
}
