package client

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/m25-lab/lightning-network-node/core_chain_sdk/account"
	"github.com/m25-lab/lightning-network-node/core_chain_sdk/channel"
	"github.com/m25-lab/lightning-network-node/database/models"
)

func (client *Client) ExchangeFwdCommitment(clientId string, accountPacked *AccountPacked, fromAmount int64, toAmount int64, transferAmount int64, fwdDest *string, hashcodeDest *string) (*models.Message, error) {
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

	//create l1 sender commitment
	channelClient := channel.NewChannel(*client.ClientCtx)
	senderCommitmentMsg := channelClient.CreateSenderCommitmentMsg(
		multisigAddr,
		accountPacked.fromAccount.AccAddress().String(),
		fromAmount,
		toAmount,
		transferAmount,
		exchangeHashcodeData.PartnerHashcode,
		*hashcodeDest,
	)
	//TODO: Dang lam!!!!!
	//create l1 sign message
	signCommitmentMsg := channel.SignMsgRequest{
		Msg:      senderCommitmentMsg,
		GasLimit: 200000,
		GasPrice: "0token", //TODO: 0token or 0stake
	}

	//sign l1 sender commitment
	strSig, err := channelClient.SignMultisigTxFromOneAccount(signCommitmentMsg, accountPacked.fromAccount, multiSigPubkey)
	if err != nil {
		return nil, err
	}

	//create ln message
	msg := models.SenderCommitment{
		Creator:          senderCommitmentMsg.Creator,
		From:             senderCommitmentMsg.From,
		Channelid:        senderCommitmentMsg.Channelid,
		CoinToSender:     senderCommitmentMsg.Cointosender.Amount.Int64(),
		CoinToHTLC:       senderCommitmentMsg.Cointohtlc.Amount.Int64(),
		HashcodeHTLC:     senderCommitmentMsg.Hashcodehtlc,
		TimelockHTLC:     senderCommitmentMsg.Timelockhtlc,
		CoinTransfer:     senderCommitmentMsg.Cointransfer.Amount.Int64(),
		HashcodeDest:     senderCommitmentMsg.Hashcodedest,
		TimelockReceiver: senderCommitmentMsg.Timelockreceiver,
		Multisig:         senderCommitmentMsg.Multisig,
	}

	partnerCommitmentPayload, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}

	//TODO: save own commit

	//TODO: send rpc (input: sendercommit with sig; output: receivercommit)

	//TODO: validate save receiverCommit, sign luon if needed

	//TODO: find way to reuse this in hop trung gian, check is dest with hashcodeDest (then to reveal transfer secret)
	return nil, nil
}
