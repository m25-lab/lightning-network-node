package client

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"

	channeltypes "github.com/m25-lab/channel/x/channel/types"
	"github.com/m25-lab/lightning-network-node/core_chain_sdk/account"
	"github.com/m25-lab/lightning-network-node/core_chain_sdk/bank"
	"github.com/m25-lab/lightning-network-node/core_chain_sdk/common"
	"github.com/m25-lab/lightning-network-node/database/models"
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
	fwdDest *string,
	hashcodeDest *string,
) error {
	//create account packed
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

	//check multisigAddr active
	multisigAddr, _, _ := account.NewAccount().CreateMulSigAccountFromTwoAccount(accountPacked.fromAccount.PublicKey(), accountPacked.toAccount.PublicKey(), 2)
	multisigAddrBalance, err := client.Balance(multisigAddr)
	if err != nil {
		return err
	}
	if multisigAddrBalance == 0 {
		err = client.Transfer(clientId, multisigAddr, 1)
		if err != nil {
			return err
		}
	}

	fromAmount := int64(0)
	toAmount := amount

	//check channel open
	isOpenChannel := true
	_, err = client.l1Client.channel.Channel(
		context.Background(),
		&channeltypes.QueryGetChannelRequest{
			Index: multisigAddr + ":token:1",
		},
	)
	if err != nil && err.Error() == "rpc error: code = NotFound desc = not found" {
		isOpenChannel = false
	}
	if !isOpenChannel {
		fromBalance, err := client.Balance(fromAccount.AccAddress().String())
		if err != nil {
			return err
		}
		if fromBalance < amount {
			return fmt.Errorf("not enough balance")
		}
	} else {
		lastestCommitment, err := client.Node.Repository.Message.FindOneByChannelIDWithAction(
			context.Background(),
			fromAccount.AccAddress().String(),
			multisigAddr+":token:1",
			models.ExchangeCommitment,
		)
		if err != nil {
			return err
		}

		payload := models.CreateCommitmentData{}
		err = json.Unmarshal([]byte(lastestCommitment.Data), &payload)
		if err != nil {
			return err
		}

		fromAmount = payload.CoinToHtlc - amount
		if fromAmount < 0 {
			return fmt.Errorf("not enough balance in channel")
		}
		toAmount = payload.CoinToCreator + amount
	}

	//exchange hashcode
	_, err = client.ExchangeHashcode(clientId, accountPacked)
	if err != nil {
		return err
	}

	_, err = client.ExchangeCommitment(clientId, accountPacked, fromAmount, toAmount, fwdDest, hashcodeDest)
	if err != nil {
		return err
	}

	//open channel
	if !isOpenChannel {
		err = client.OpenChannel(clientId, accountPacked)
		if err != nil {
			return err
		}
	}

	return nil
}

func (client *Client) Transfer(clientId string, toAddress string, value int64) error {
	if strings.Contains(toAddress, "@") {
		parsedAddress := strings.Split(toAddress, "@")
		toAddress = parsedAddress[0]
	}
	fmt.Print("toAddress: ", toAddress, "\n")

	fromAccount, err := client.CurrentAccount(clientId)
	if err != nil {
		return err
	}

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
