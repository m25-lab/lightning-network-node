package client

import (
	"github.com/m25-lab/lightning-network-node/database/models"
)

func (client *Client) ExchangeSecret(clientId string, accountPacked *AccountPacked, hashcodePayload models.ExchangeHashcodeData) (*models.Message, error) {
	//TODO: send rpc to endpoint in accountPacked,
	//TODO: send Model: mySecret, Myhashcode, PartnerHashcode, rep: secret
	//TODO: updateOwn data with partner secret
	return nil, nil
}
