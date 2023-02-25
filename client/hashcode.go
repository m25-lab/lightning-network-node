package client

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/m25-lab/lightning-network-node/core_chain_sdk/account"
	"github.com/m25-lab/lightning-network-node/core_chain_sdk/common"
	"github.com/m25-lab/lightning-network-node/database/models"
	"github.com/m25-lab/lightning-network-node/rpc/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (client *Client) ExchangeHashcode(clientId string, accountPacked *AccountPacked) (*models.Message, error) {

	multisigAddr, _, _ := account.NewAccount().CreateMulSigAccountFromTwoAccount(accountPacked.fromAccount.PublicKey(), accountPacked.toAccount.PublicKey(), 2)

	//random secret
	secret, err := common.RandomSecret()
	if err != nil {
		return nil, err
	}
	hashCode := common.ToHashCode(secret)

	ID := primitive.NewObjectID()
	savedMessage := models.Message{
		ID:         ID,
		OriginalID: ID,
		ChannelID:  multisigAddr + ":token:1",
		Action:     models.ExchangeHashcode,
		Owner:      accountPacked.fromAccount.AccAddress().String(),
		Users: []string{
			accountPacked.fromAccount.AccAddress().String() + "@" + client.Node.Config.LNode.External,
			accountPacked.toAccount.AccAddress().String() + "@" + accountPacked.toEndpoint,
		},
		IsReplied: false,
	}

	payload, err := json.Marshal(models.ExchangeHashcodeData{
		PartnerHashcode: hashCode,
	})
	if err != nil {
		return nil, err
	}

	rpcClient := pb.NewMessageServiceClient(client.CreateConn(accountPacked.toEndpoint))
	reponse, err := rpcClient.SendMessage(context.Background(), &pb.SendMessageRequest{
		MessageId: ID.Hex(),
		ChannelID: savedMessage.ChannelID,
		Action:    models.ExchangeHashcode,
		Data:      string(payload),
		From:      accountPacked.fromAccount.AccAddress().String() + "@" + client.Node.Config.LNode.External,
		To:        accountPacked.toAccount.AccAddress().String() + "@" + accountPacked.toEndpoint,
	})
	if err != nil {
		return nil, err
	}
	if reponse.ErrorCode != "" {
		return nil, errors.New(reponse.ErrorCode)
	}

	responsePayload := models.ExchangeHashcodeData{}
	err = json.Unmarshal([]byte(reponse.Response), &responsePayload)
	if err != nil {
		return nil, err
	}

	payload, err = json.Marshal(models.ExchangeHashcodeData{
		MySecret:        secret,
		MyHashcode:      hashCode,
		PartnerHashcode: responsePayload.PartnerHashcode,
	})
	if err != nil {
		return nil, err
	}
	savedMessage.Data = string(payload)
	err = client.Node.Repository.Message.InsertOne(context.Background(), &savedMessage)
	if err != nil {
		return nil, err
	}

	return &savedMessage, nil
}
