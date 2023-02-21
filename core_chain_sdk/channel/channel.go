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
	multiSigPubkey cryptoTypes.PubKey) (string, error) {

	err := req.Msg.ValidateBasic()
	if err != nil {
		return "", err
	}

	newTx := common.NewMultisigTxBuilder(cn.rpcClient, account, req.GasLimit, req.GasPrice, 0, 2)
	txBuilder, err := newTx.BuildUnsignedTx(req.Msg)

	if err != nil {
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
