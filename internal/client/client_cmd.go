package client

import (
	"fmt"
	channelTypes "github.com/AstraProtocol/channel/x/channel/types"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/m25-lab/lightning-network-node/internal/bank"
	"github.com/m25-lab/lightning-network-node/internal/channel"
	"github.com/m25-lab/lightning-network-node/internal/common"
	"math"
	"math/big"
)

func OpenChannelFromA() {
	c := NewClient()
	acc := c.NewAccountClient()
	AAccount, _ := acc.ImportAccount("series divide ripple fire person prepare meat smooth source scrap poet quit shoulder choice leaf friend pact fault toddler simple quit popular define jar")
	BAccount, _ := acc.ImportAccount("perfect hello crystal august lake giant dutch random season onion acid stable edge reform deposit capable family glow air elegant copper punch student runway")
	fmt.Println("account A:", AAccount.AccAddress().String())
	fmt.Println("account B:", BAccount.AccAddress().String())
	fmt.Println("PrivateKey", AAccount.PrivateKeyToString())
	fmt.Println("ÏÏPublicKey", AAccount.PublicKey())

	multisigAddr, multiSigPubkey, _ := acc.CreateMulSignAccountFromTwoAccount(AAccount.PublicKey(), BAccount.PublicKey(), 2)
	transfer(c, AAccount.PrivateKeyToString(), multisigAddr)
	fmt.Println("multisigAddr", multisigAddr)

	openChannelRequest := channel.SignMsgRequest{
		Msg: channelTypes.NewMsgOpenChannel(
			multisigAddr,
			AAccount.AccAddress().String(),
			BAccount.AccAddress().String(),
			&types.Coin{
				Denom:  "token",
				Amount: types.NewInt(10),
			},
			&types.Coin{
				Denom:  "token",
				Amount: types.NewInt(10),
			},
			multisigAddr,
			"1",
		),
		GasLimit: 200000,
		GasPrice: "0token",
	}

	channelClient := c.NewChannelClient()

	msg, strSig, err := channelClient.CreateMultisigMsg(openChannelRequest, AAccount, multiSigPubkey)
	if err != nil {
		panic(err)
	}

	txJSONBytes, err := c.RpcClient().TxConfig.TxJSONEncoder()(msg)
	if err != nil {
		panic(err)
	}
	txJSON := string(txJSONBytes)
	fmt.Println(txJSON, strSig)

}

func transfer(c *Client, privateKey string, toAddress string) {
	bankClient := c.NewBankClient()
	amount := big.NewInt(0).Mul(big.NewInt(10), big.NewInt(0).SetUint64(uint64(math.Pow10(1))))
	request := &bank.TransferRequest{
		PrivateKey: privateKey,
		Receiver:   toAddress,
		Amount:     amount,
		GasLimit:   200000,
		GasPrice:   "0token",
	}

	txBuilder, err := bankClient.TransferRawDataWithPrivateKey(request)
	if err != nil {
		panic(err)
	}

	txJson, err := common.TxBuilderJsonEncoder(c.RpcClient().TxConfig, txBuilder)
	if err != nil {
		panic(err)
	}

	txByte, err := common.TxBuilderJsonDecoder(c.RpcClient().TxConfig, txJson)
	if err != nil {
		panic(err)
	}

	txHash := common.TxHash(txByte)
	fmt.Println("txHash", txHash)

	_, err = c.RpcClient().BroadcastTxCommit(txByte)

	if err != nil {
		panic(err)
	}
}
