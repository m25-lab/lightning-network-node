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

func OpenChannel() {
	c := NewClient()
	acc := c.NewAccountClient()
	A_account, _ := acc.ImportAccount("series divide ripple fire person prepare meat smooth source scrap poet quit shoulder choice leaf friend pact fault toddler simple quit popular define jar")
	B_account, _ := acc.ImportAccount("perfect hello crystal august lake giant dutch random season onion acid stable edge reform deposit capable family glow air elegant copper punch student runway")
	fmt.Println("account A:", A_account.AccAddress().String())
	fmt.Println("account B:", B_account.AccAddress().String())

	multisigAddr, multiSigPubkey, _ := acc.CreateMulSignAccountFromTwoAccount(A_account.PublicKey(), B_account.PublicKey(), 2)
	transfer(c, A_account.PrivateKeyToString(), B_account.AccAddress().String())
	fmt.Println("multisigAddr", multisigAddr)

	openChannelRequest := channel.SignMsgRequest{
		Msg: channelTypes.NewMsgOpenChannel(
			multisigAddr,
			A_account.AccAddress().String(),
			B_account.AccAddress().String(),
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

	msg, strSig, err := channelClient.CreateMultisigMsg(openChannelRequest, A_account, multiSigPubkey)
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

	fmt.Println(txBuilder.GetTx().GetMsgs())

	txJson, err := common.TxBuilderJsonEncoder(c.RpcClient().TxConfig, txBuilder)
	if err != nil {
		panic(err)
	}

	fmt.Println(txJson)

	txByte, err := common.TxBuilderJsonDecoder(c.RpcClient().TxConfig, txJson)
	if err != nil {
		panic(err)
	}

	txHash := common.TxHash(txByte)
	fmt.Println("txHash", txHash)

	response, err := c.RpcClient().BroadcastTxCommit(txByte)
	fmt.Println(response)
	if err != nil {
		panic(err)
	}
}
