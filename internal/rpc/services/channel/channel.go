package channel

import (
	"context"
	"github.com/AstraProtocol/astra-go-sdk/account"
	channelTypes "github.com/AstraProtocol/channel/x/channel/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	cryptoTypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
)

type Channel struct {
	rpcClient client.Context
}

type SignMsgRequest struct {
	Msg      sdk.Msg
	GasLimit uint64
	GasPrice string
}

type TxMulSign struct {
	txf              tx.Factory
	signerPrivateKey *account.PrivateKeySerialized
	rpcClient        client.Context
}

func NewChannel(rpcClient client.Context) *Channel {
	return &Channel{rpcClient}
}

func (cn *Channel) CreateMultisigMsg(req SignMsgRequest,
	account *account.PrivateKeySerialized,
	multiSigPubkey cryptoTypes.PubKey) (sdk.Tx, string, error) {

	err := req.Msg.ValidateBasic()
	if err != nil {
		return nil, "", err
	}

	newTx := NewTxMulSign(cn.rpcClient, account, req.GasLimit, req.GasPrice, 0, 2)
	txBuilder, err := newTx.BuildUnsignedTx(req.Msg)

	if err != nil {
		return nil, "", err
	}

	return txBuilder.GetTx().(sdk.Tx), "", nil
}

func (cn *Channel) ListChannel() (*channelTypes.QueryAllChannelResponse, error) {
	channelClient := channelTypes.NewQueryClient(cn.rpcClient)
	return channelClient.ChannelAll(context.Background(), &channelTypes.QueryAllChannelRequest{})
}

func NewTxMulSign(rpcClient client.Context, privateKey *account.PrivateKeySerialized, gasLimit uint64, gasPrice string, sequenNum, accNum uint64) *TxMulSign {
	txf := tx.Factory{}.
		WithChainID(rpcClient.ChainID).
		WithTxConfig(rpcClient.TxConfig).
		WithGasPrices(gasPrice).
		WithGas(gasLimit).
		WithSequence(sequenNum).
		WithAccountNumber(accNum).
		WithSignMode(signing.SignMode_SIGN_MODE_LEGACY_AMINO_JSON)
	//.SetTimeoutHeight(txf.TimeoutHeight())

	return &TxMulSign{txf: txf, signerPrivateKey: privateKey, rpcClient: rpcClient}
}

func (t *TxMulSign) BuildUnsignedTx(msgs sdk.Msg) (client.TxBuilder, error) {
	return t.txf.BuildUnsignedTx(msgs)
}
