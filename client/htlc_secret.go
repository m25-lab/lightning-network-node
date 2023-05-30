package client

import (
	"context"
	"errors"
	"github.com/m25-lab/lightning-network-node/core_chain_sdk/common"
	"github.com/m25-lab/lightning-network-node/database/models"
	"github.com/m25-lab/lightning-network-node/rpc/pb"
)

func (client *Client) ExchangeSecret(clientId string, accountPacked *AccountPacked, hashcodePayload models.ExchangeHashcodeData) (*models.ExchangeHashcodeData, error) {
	rpcClient := pb.NewMessageServiceClient(client.CreateConn(accountPacked.toEndpoint))
	response, err := rpcClient.SendSecret(context.Background(), &pb.SendSecretMessage{
		MySecret:        hashcodePayload.MySecret,
		MyHashcode:      hashcodePayload.MyHashcode,
		PartnerHashcode: hashcodePayload.PartnerHashcode,
	})
	if err != nil {
		return nil, err
	}
	if response.ErrorCode != "" {
		return nil, errors.New(response.ErrorCode)
	}

	//rehash
	if common.ToHashCode(response.Secret) != hashcodePayload.PartnerHashcode {
		return nil, errors.New("secret from " + accountPacked.toAccount.AccAddress().String() + " not matched")
	}

	input := models.ExchangeHashcodeData{
		MySecret:        hashcodePayload.MySecret,
		MyHashcode:      hashcodePayload.MyHashcode,
		PartnerHashcode: hashcodePayload.PartnerHashcode,
		PartnerSecret:   response.Secret,
		ChannelID:       hashcodePayload.ChannelID,
	}
	//update secret into ExchangeHashcode table
	err = client.Node.Repository.ExchangeHashcode.UpdateSecret(context.Background(), &input)
	if err != nil {
		return nil, err
	}
	return &input, nil
}
