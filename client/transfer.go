package client

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/m25-lab/lightning-network-node/core_chain_sdk/account"
	"github.com/m25-lab/lightning-network-node/core_chain_sdk/bank"
	"github.com/m25-lab/lightning-network-node/core_chain_sdk/common"
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

	multisigAddr, _, _ := account.NewAccount().CreateMulSigAccountFromTwoAccount(accountPacked.fromAccount.PublicKey(), accountPacked.toAccount.PublicKey(), 2)
	err = client.Transfer(fromAccount, multisigAddr, 1)
	if err != nil {
		return err
	}

	_, err = client.ExchangeHashcode(clientId, accountPacked)
	if err != nil {
		return err
	}

	_, err = client.ExchangeCommitment(clientId, accountPacked, amount)
	if err != nil {
		return err
	}

	err = client.OpenChannel(clientId, accountPacked)
	if err != nil {
		return err
	}

	return nil
}

func (client *Client) Transfer(fromAccount *account.PrivateKeySerialized, toAddress string, value int64) error {
	bankClient := bank.NewBank(*client.ClientCtx, "token", 60)
	request := &bank.TransferRequest{
		PrivateKey: fromAccount.PrivateKeyToString(),
		Receiver:   toAddress,
		Amount:     big.NewInt(value),
		GasLimit:   100000,
		GasPrice:   "0token",
	}

	txBuilder, err := bankClient.TransferRawDataWithPrivateKey(request)
	if err != nil {
		return err
	}

	txJson, err := common.TxBuilderJsonEncoder(client.ClientCtx.TxConfig, txBuilder)
	if err != nil {
		return err
	}

	txByte, err := common.TxBuilderJsonDecoder(client.ClientCtx.TxConfig, txJson)
	if err != nil {
		return err
	}

	broadcastResponse, err := client.ClientCtx.BroadcastTxCommit(txByte)
	if err != nil {
		return err
	}
	fmt.Printf("Transfer: %s\n", broadcastResponse.String())

	return nil
}
