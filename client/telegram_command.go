package client

import (
	"context"

	"github.com/m25-lab/lightning-network-node/core_chain_sdk/account"
	"github.com/m25-lab/lightning-network-node/database/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (client *Client) CreateAccount(clientId string) (*account.PrivateKeySerialized, error) {
	acc := account.NewAccount()

	account, err := acc.CreateAccount()
	if err != nil {
		return nil, err
	}

	err = client.Node.Repository.Address.InsertOne(context.Background(), &models.Address{
		ID:       primitive.NewObjectID(),
		Address:  account.AccAddress().String(),
		Pubkey:   account.PublicKey().String(),
		Mnemonic: account.Mnemonic(),
		ClientId: clientId,
	})

	if err != nil {
		return nil, err
	}

	return account, nil
}
