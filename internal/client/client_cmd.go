package client

import (
	"context"
	"encoding/json"
	"fmt"
	channelTypes "github.com/AstraProtocol/channel/x/channel/types"
	"github.com/cosmos/cosmos-sdk/types"
	signingTypes "github.com/cosmos/cosmos-sdk/types/tx/signing"
	"github.com/m25-lab/lightning-network-node/internal/bank"
	"github.com/m25-lab/lightning-network-node/internal/channel"
	"github.com/m25-lab/lightning-network-node/internal/common"
	"github.com/m25-lab/lightning-network-node/node/rpc/pb"
	"math"
	"math/big"
)

func OpenChannelFromA() string {
	cfg := &Config{
		ChainId:               "channel",
		Endpoint:              "http://0.0.0.0:26657",
		LightningNodeEndpoint: "0.0.0.0:2525",
		CoinType:              60,
		PrefixAddress:         "cosmos",
		TokenSymbol:           "token",
	}

	c := NewClient(cfg)
	acc := c.NewAccountClient()
	AAccount, _ := acc.ImportAccount("series divide ripple fire person prepare meat smooth source scrap poet quit shoulder choice leaf friend pact fault toddler simple quit popular define jar")
	BAccount, _ := acc.ImportAccount("perfect hello crystal august lake giant dutch random season onion acid stable edge reform deposit capable family glow air elegant copper punch student runway")
	fmt.Println("account A:", AAccount.AccAddress().String())
	fmt.Println("account B:", BAccount.AccAddress().String())
	fmt.Println("PrivateKey", AAccount.PrivateKeyToString())
	fmt.Println("PublicKey", AAccount.PublicKey())

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
				Amount: types.NewInt(100),
			},
			&types.Coin{
				Denom:  "token",
				Amount: types.NewInt(100),
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

	strOpenChannelRequest, _ := json.Marshal(openChannelRequest)

	response, _ := c.rpcLightningNode.channel.OpenChannel(
		context.Background(),
		&pb.OpenChannelRequest{
			FromAddress: AAccount.AccAddress().String(),
			ToAddress:   multisigAddr,
			Signature:   strSig,
			Payload:     string(strOpenChannelRequest),
		})
	return response.GetResponse()
}

func OpenChannelFromB(channelId string) {
	cfg := &Config{
		ChainId:               "channel",
		Endpoint:              "http://0.0.0.0:26657",
		LightningNodeEndpoint: "0.0.0.0:2525",
		CoinType:              60,
		PrefixAddress:         "cosmos",
		TokenSymbol:           "token",
	}

	c := NewClient(cfg)
	acc := c.NewAccountClient()
	channelClient := c.NewChannelClient()

	AAccount, _ := acc.ImportAccount("series divide ripple fire person prepare meat smooth source scrap poet quit shoulder choice leaf friend pact fault toddler simple quit popular define jar")
	BAccount, _ := acc.ImportAccount("perfect hello crystal august lake giant dutch random season onion acid stable edge reform deposit capable family glow air elegant copper punch student runway")

	channelResult, _ := c.rpcLightningNode.channel.GetChannelById(context.Background(), &pb.GetChannelRequest{Id: channelId})

	var payload struct {
		Msg      channelTypes.MsgOpenChannel
		GasLimit uint64
		GasPrice string
	}

	fmt.Println("payload", channelResult.Payload)
	json.Unmarshal([]byte(channelResult.Payload), &payload)

	tmpMultisigAddr, multiSigPubkey, _ := acc.CreateMulSignAccountFromTwoAccount(AAccount.PublicKey(), BAccount.PublicKey(), 2)
	fmt.Println("tmpMultisigAddr", tmpMultisigAddr) // astra1747xvksuc7ecpylckzjvmcqvvlmp6t6ujs3lld

	msg := channelTypes.NewMsgOpenChannel(
		payload.Msg.Creator,
		payload.Msg.PartA,
		payload.Msg.PartB,
		payload.Msg.CoinA,
		payload.Msg.CoinB,
		payload.Msg.MultisigAddr,
		payload.Msg.Sequence,
	)

	openChannelRequest := channel.SignMsgRequest{
		Msg:      msg,
		GasLimit: payload.GasLimit,
		GasPrice: payload.GasPrice,
	}

	fmt.Println("Sig1", channelResult.SignatureA)
	fmt.Println("openChannelRequest", openChannelRequest)

	signList := make([][]signingTypes.SignatureV2, 0)
	signByte1, err := common.TxBuilderSignatureJsonDecoder(c.RpcClient().TxConfig, channelResult.SignatureA)
	if err != nil {
		panic(err)
	}
	signList = append(signList, signByte1)
	_, strSig2, err := channelClient.CreateMultisigMsg(openChannelRequest, BAccount, multiSigPubkey)
	if err != nil {
		panic(err)
	}
	fmt.Println("Sig2", strSig2)

	signByte2, err := common.TxBuilderSignatureJsonDecoder(c.RpcClient().TxConfig, strSig2)
	if err != nil {
		panic(err)
	}

	signList = append(signList, signByte2)

	fmt.Println("new tx multisign")

	newTx := common.NewTxMulSign(c.RpcClient(),
		nil,
		openChannelRequest.GasLimit,
		openChannelRequest.GasPrice,
		0,
		2)

	txBuilderMultiSign, err := newTx.BuildUnsignedTx(openChannelRequest.Msg)
	if err != nil {
		panic(err)
	}

	err = newTx.CreateTxMulSign(txBuilderMultiSign, multiSigPubkey, signList)
	if err != nil {
		panic(err)
	}

	txJson, err := common.TxBuilderJsonEncoder(c.RpcClient().TxConfig, txBuilderMultiSign)
	if err != nil {
		panic(err)
	}
	fmt.Println("rawData", string(txJson))

	txByte, err := common.TxBuilderJsonDecoder(c.RpcClient().TxConfig, txJson)
	if err != nil {
		panic(err)
	}

	txHash := common.TxHash(txByte)
	fmt.Println("txHash", txHash)

	txResult, err := c.RpcClient().BroadcastTxCommit(txByte)
	if err != nil {
		panic(err)
	}
	fmt.Println("tx openchannel result code", txResult.Code)
	fmt.Println(txResult)

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
