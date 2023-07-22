package channel

import (
	"context"
	"github.com/cosmos/cosmos-sdk/client"
	cryptoTypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	channelTypes "github.com/m25-lab/channel/x/channel/types"
	"github.com/m25-lab/lightning-network-node/core_chain_sdk/account"
	"github.com/m25-lab/lightning-network-node/core_chain_sdk/common"
	"github.com/pkg/errors"
	"log"
)

type Channel struct {
	rpcClient client.Context
}

type SignMsgRequest struct {
	Msg      sdk.Msg
	GasLimit uint64
	GasPrice string
}

func NewChannel(rpcClient client.Context) *Channel {
	return &Channel{rpcClient}
}

func (cn *Channel) SignMultisigTxFromOneAccount(req SignMsgRequest,
	account *account.PrivateKeySerialized,
	multiSigPubkey cryptoTypes.PubKey,
	isFirstCommit bool) (string, error) {

	err := req.Msg.ValidateBasic()
	if err != nil {
		return "", err
	}

	from := types.AccAddress(multiSigPubkey.Address())
	accNum, accSeq, err := cn.rpcClient.AccountRetriever.GetAccountNumberSequence(cn.rpcClient, from)
	if isFirstCommit {
		accSeq += 1
	}
	newTx := common.NewMultisigTxBuilder(cn.rpcClient, account, req.GasLimit, req.GasPrice, accSeq, accNum)
	txBuilder, err := newTx.BuildUnsignedTx(req.Msg)

	if err != nil {
		log.Println(err)
		return "", err
	}

	err = newTx.SignTxWithSignerAddress(txBuilder, multiSigPubkey)
	if err != nil {
		return "", errors.Wrap(err, "SignTx")
	}

	sign, err := common.TxBuilderSignatureJsonEncoder(cn.rpcClient.TxConfig, txBuilder)
	if err != nil {
		return "", errors.Wrap(err, "GetSign")
	}

	return sign, nil
}

func (cn *Channel) ListChannel() (*channelTypes.QueryAllChannelResponse, error) {
	channelClient := channelTypes.NewQueryClient(cn.rpcClient)
	return channelClient.ChannelAll(context.Background(), &channelTypes.QueryAllChannelRequest{})
}

func (cn *Channel) CreateCommitmentMsg(
	multisigAddr string,
	toTimelockAddr string,
	coinToCreator int64,
	toHashlockAddr string,
	coinToHtlc int64,
	hashCode string,
) *channelTypes.MsgCommitment {
	return &channelTypes.MsgCommitment{
		ChannelID:      multisigAddr + ":token:1",
		Creator:        multisigAddr,
		From:           multisigAddr,
		Timelock:       100,
		ToTimelockAddr: toTimelockAddr,
		CoinToCreator: &types.Coin{
			Denom:  "token",
			Amount: types.NewInt(coinToCreator),
		},
		ToHashlockAddr: toHashlockAddr,
		Hashcode:       hashCode,
		CoinToHtlc: &types.Coin{
			Denom:  "token",
			Amount: types.NewInt(coinToHtlc),
		},
	}
}

func (cn *Channel) CreateSenderCommitmentMsg(
	multisigAddr string,
	fromAddr string,
	coinToOriginal int64,
	coinToHtlc int64,
	coinTransfer int64,
	hashCode string,
	hashCodeDest string,
	senderLock string,
) *channelTypes.MsgSendercommit {
	return &channelTypes.MsgSendercommit{
		Creator:   multisigAddr,
		From:      fromAddr,
		ChannelID: multisigAddr + ":token:1",
		CoinToSender: &types.Coin{
			Denom:  "token",
			Amount: types.NewInt(coinToOriginal),
		},
		CoinToHtlc: &types.Coin{
			Denom:  "token",
			Amount: types.NewInt(coinToHtlc),
		},
		HashcodeHtlc: hashCode,
		TimelockHtlc: "100",
		CoinTransfer: &types.Coin{
			Denom:  "token",
			Amount: types.NewInt(coinTransfer),
		},
		HashcodeDest:     hashCodeDest,
		TimelockReceiver: "100",
		TimelockSender:   senderLock,
		Multisig:         multisigAddr,
	}
}

func (cn *Channel) CreateReceiverCommitmentMsg(
	multisigAddr string,
	fromAddr string,
	coinToOriginal int64,
	coinToHtlc int64,
	coinTransfer int64,
	hashCode string,
	hashCodeDest string,
	senderLock string,
) *channelTypes.MsgReceivercommit {
	return &channelTypes.MsgReceivercommit{
		Creator:   multisigAddr,
		From:      fromAddr,
		ChannelID: multisigAddr + ":token:1",
		CoinToReceiver: &types.Coin{
			Denom:  "token",
			Amount: types.NewInt(coinToOriginal),
		},
		CoinToHtlc: &types.Coin{
			Denom:  "token",
			Amount: types.NewInt(coinToHtlc),
		},
		HashcodeHtlc: hashCode,
		TimelockHtlc: "100",
		CoinTransfer: &types.Coin{
			Denom:  "token",
			Amount: types.NewInt(coinTransfer),
		},
		HashcodeDest:   hashCodeDest,
		TimelockSender: senderLock,
		Multisig:       multisigAddr,
	}
}

func (cn *Channel) CreateOpenChannelMsg(
	multisigAddr string,
	partA string,
	partB string,
	coinA int64,
	coinB int64,
) *channelTypes.MsgOpenChannel {
	return &channelTypes.MsgOpenChannel{
		Creator:      multisigAddr,
		MultisigAddr: multisigAddr,
		Sequence:     "1",
		PartA:        partA,
		PartB:        partB,
		CoinA: &types.Coin{
			Denom:  "token",
			Amount: types.NewInt(coinB),
		},
		CoinB: &types.Coin{
			Denom:  "token",
			Amount: types.NewInt(coinA),
		},
	}
}
