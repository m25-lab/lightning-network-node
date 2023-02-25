package client

import (
	"context"
	"strings"

	"github.com/m25-lab/lightning-network-node/core_chain_sdk/account"
)

type AccountPacked struct {
	fromAccount *account.PrivateKeySerialized
	toAccount   *account.PKAccount
	toEndpoint  string
}

func (client *Client) LnTransfer(
	clientId string,
	to string,
	amount int64,
) error {
	fromAccount, err := client.CurrentAccount(clientId)
	if err != nil {
		return err
	}

	existedWhitelist, err := client.Node.Repository.Whitelist.FindOneByPartnerAddress(context.Background(), fromAccount.AccAddress().String(), to)
	if err != nil {
		return err
	}
	toAccount := account.NewPKAccount(existedWhitelist.PartnerPubkey)
	toEndpoint := strings.Split(to, "@")[1]

	accountPacked := &AccountPacked{
		fromAccount: fromAccount,
		toAccount:   toAccount,
		toEndpoint:  toEndpoint,
	}

	// _, err = client.ExchangeHashcode(clientId, accountPacked)
	// if err != nil {
	// 	return err
	// }

	// _, err = client.ExchangeCommitment(clientId, accountPacked, amount)
	// if err != nil {
	// 	return err
	// }

	// multisigAddr, _, _ := account.NewAccount().CreateMulSigAccountFromTwoAccount(accountPacked.fromAccount.PublicKey(), accountPacked.toAccount.PublicKey(), 2)
	// err = client.Transfer(fromAccount, multisigAddr, 1)
	// if err != nil {
	// 	return err
	// }

	err = client.OpenChannel(clientId, accountPacked)
	if err != nil {
		return err
	}

	return nil
}
