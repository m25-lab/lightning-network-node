package client

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"math/big"

// 	channelTypes "github.com/AstraProtocol/channel/x/channel/types"
// 	"github.com/cosmos/cosmos-sdk/types"
// 	signingTypes "github.com/cosmos/cosmos-sdk/types/tx/signing"
// 	"github.com/m25-lab/lightning-network-node/core_chain_sdk/bank"
// 	"github.com/m25-lab/lightning-network-node/core_chain_sdk/channel"
// 	"github.com/m25-lab/lightning-network-node/core_chain_sdk/common"
// 	"github.com/m25-lab/lightning-network-node/rpc/pb"
// )

// func CreateCommitmentFromA(amountA int64, amountB int64, secret string, timelock uint64, seq int) (channelTypes.MsgCommitment, string, string) {
// 	cfg := &Config{
// 		ChainId:               "channel",
// 		Endpoint:              "http://0.0.0.0:26657",
// 		LightningNodeEndpoint: "0.0.0.0:2525",
// 		CoinType:              60,
// 		PrefixAddress:         "cosmos",
// 		TokenSymbol:           "token",
// 	}

// 	c := NewClient(cfg)
// 	acc := c.NewAccountClient()
// 	AAccount, _ := acc.ImportAccount("excuse quiz oyster vendor often spray day vanish slice topic pudding crew promote floor shadow best subway slush slender good merit hollow certain repeat")
// 	BAccount, _ := acc.ImportAccount("claim market flip canoe wreck maid recipe bright fuel slender ladder album behind repeat come trophy come vicious frown prefer height unknown thank damp")
// 	multisigAddr, multiSigPubkey, _ := acc.CreateMulSigAccountFromTwoAccount(AAccount.PublicKey(), BAccount.PublicKey(), 2)

// 	partACommitment := channelTypes.MsgCommitment{
// 		ChannelID:      multisigAddr + ":token:" + fmt.Sprint(seq),
// 		Creator:        multisigAddr,
// 		From:           multisigAddr,
// 		Timelock:       uint64(timelock),
// 		ToTimelockAddr: BAccount.AccAddress().String(),
// 		CoinToCreator: &types.Coin{
// 			Denom:  "token",
// 			Amount: types.NewInt(amountA),
// 		},
// 		ToHashlockAddr: AAccount.AccAddress().String(),
// 		Hashcode:       common.ToHashCode(secret),
// 		CoinToHtlc: &types.Coin{
// 			Denom:  "token",
// 			Amount: types.NewInt(60),
// 		},
// 	}

// 	commitmentA := channel.SignMsgRequest{
// 		Msg:      &partACommitment,
// 		GasLimit: 200000,
// 		GasPrice: "0token",
// 	}

// 	channelClient := c.NewChannelClient()

// 	strSig, err := channelClient.SignMultisigTxFromOneAccount(commitmentA, AAccount, multiSigPubkey)
// 	if err != nil {
// 		panic(err)
// 	}

// 	return partACommitment, AAccount.AccAddress().String(), strSig
// }

// func CreateCommitmentFromB(partACommitment channelTypes.MsgCommitment, aAddress string, aSignature string, secret string) (channelTypes.MsgCommitment, string, string) {
// 	cfg := &Config{
// 		ChainId:               "channel",
// 		Endpoint:              "http://0.0.0.0:26657",
// 		LightningNodeEndpoint: "0.0.0.0:2525",
// 		CoinType:              60,
// 		PrefixAddress:         "cosmos",
// 		TokenSymbol:           "token",
// 	}

// 	c := NewClient(cfg)
// 	channelClient := c.NewChannelClient()

// 	acc := c.NewAccountClient()
// 	AAccount, _ := acc.ImportAccount("excuse quiz oyster vendor often spray day vanish slice topic pudding crew promote floor shadow best subway slush slender good merit hollow certain repeat")
// 	BAccount, _ := acc.ImportAccount("claim market flip canoe wreck maid recipe bright fuel slender ladder album behind repeat come trophy come vicious frown prefer height unknown thank damp")

// 	multisigAddr, multiSigPubkey, _ := acc.CreateMulSigAccountFromTwoAccount(AAccount.PublicKey(), BAccount.PublicKey(), 2)

// 	// @Description: Check if the signature is correct
// 	commitmentA := channel.SignMsgRequest{
// 		Msg:      &partACommitment,
// 		GasLimit: 200000,
// 		GasPrice: "0token",
// 	}
// 	strACommitment, _ := json.Marshal(commitmentA)

// 	// @Description: Save commitment A
// 	response, err := c.rpcLightningNode.channel.CreateCommitment(
// 		context.Background(),
// 		&pb.CreateCommitmentRequest{
// 			ChannelId:   partACommitment.ChannelID,
// 			FromAddress: aAddress,
// 			Payload:     string(strACommitment),
// 			Signature:   aSignature,
// 		},
// 	)
// 	fmt.Println(response, err)

// 	// @Description: Create commitment B
// 	partBCommitment := channelTypes.MsgCommitment{
// 		ChannelID:      partACommitment.ChannelID,
// 		Creator:        multisigAddr,
// 		From:           multisigAddr,
// 		Timelock:       partACommitment.Timelock,
// 		ToTimelockAddr: partACommitment.ToHashlockAddr,
// 		CoinToCreator:  partACommitment.CoinToHtlc,
// 		ToHashlockAddr: partACommitment.ToTimelockAddr,
// 		Hashcode:       common.ToHashCode(secret),
// 		CoinToHtlc:     partACommitment.CoinToCreator,
// 	}

// 	// @Description: Sign commitment B
// 	commitmentB := channel.SignMsgRequest{
// 		Msg:      &partBCommitment,
// 		GasLimit: 200000,
// 		GasPrice: "0token",
// 	}

// 	strSig, err := channelClient.SignMultisigTxFromOneAccount(commitmentB, BAccount, multiSigPubkey)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return partBCommitment, BAccount.AccAddress().String(), strSig
// }

// func StoreCommitmentFromA(partBCommitment channelTypes.MsgCommitment, bAddress string, bSignature string) {
// 	cfg := &Config{
// 		ChainId:               "channel",
// 		Endpoint:              "http://0.0.0.0:26657",
// 		LightningNodeEndpoint: "0.0.0.0:2525",
// 		CoinType:              60,
// 		PrefixAddress:         "cosmos",
// 		TokenSymbol:           "token",
// 	}

// 	c := NewClient(cfg)

// 	commitmentB := channel.SignMsgRequest{
// 		Msg:      &partBCommitment,
// 		GasLimit: 200000,
// 		GasPrice: "0token",
// 	}

// 	strBCommitment, _ := json.Marshal(commitmentB)

// 	// @Description: Save commitment B
// 	response, err := c.rpcLightningNode.channel.CreateCommitment(
// 		context.Background(),
// 		&pb.CreateCommitmentRequest{
// 			ChannelId:   partBCommitment.ChannelID,
// 			FromAddress: bAddress,
// 			Payload:     string(strBCommitment),
// 			Signature:   bSignature,
// 		},
// 	)
// 	fmt.Println(response, err)
// }

// func OpenChannelFromA(amountA int64, amountB int64, seq int) string {
// 	cfg := &Config{
// 		ChainId:               "channel",
// 		Endpoint:              "http://0.0.0.0:26657",
// 		LightningNodeEndpoint: "0.0.0.0:2525",
// 		CoinType:              60,
// 		PrefixAddress:         "cosmos",
// 		TokenSymbol:           "token",
// 	}

// 	c := NewClient(cfg)
// 	acc := c.NewAccountClient()
// 	AAccount, _ := acc.ImportAccount("excuse quiz oyster vendor often spray day vanish slice topic pudding crew promote floor shadow best subway slush slender good merit hollow certain repeat")
// 	BAccount, _ := acc.ImportAccount("claim market flip canoe wreck maid recipe bright fuel slender ladder album behind repeat come trophy come vicious frown prefer height unknown thank damp")

// 	fmt.Println("account A:", AAccount.AccAddress().String())
// 	fmt.Println("account B:", BAccount.AccAddress().String())
// 	fmt.Println("PrivateKey", AAccount.PrivateKeyToString())
// 	fmt.Println("PublicKey", AAccount.PublicKey())

// 	multisigAddr, multiSigPubkey, _ := acc.CreateMulSigAccountFromTwoAccount(AAccount.PublicKey(), BAccount.PublicKey(), 2)
// 	transfer(c, AAccount.PrivateKeyToString(), multisigAddr, big.NewInt(1))
// 	fmt.Println("multisigAddr", multisigAddr)

// 	openChannelRequest := channel.SignMsgRequest{
// 		Msg: &channelTypes.MsgOpenChannel{
// 			Creator: multisigAddr,
// 			PartA:   AAccount.AccAddress().String(),
// 			PartB:   BAccount.AccAddress().String(),
// 			CoinA: &types.Coin{
// 				Denom:  "token",
// 				Amount: types.NewInt(amountA),
// 			},
// 			CoinB: &types.Coin{
// 				Denom:  "token",
// 				Amount: types.NewInt(amountB),
// 			},
// 			MultisigAddr: multisigAddr,
// 			Sequence:     fmt.Sprint(seq),
// 		},
// 		GasLimit: 200000,
// 		GasPrice: "0token",
// 	}

// 	channelClient := c.NewChannelClient()

// 	strSig, err := channelClient.SignMultisigTxFromOneAccount(openChannelRequest, AAccount, multiSigPubkey)
// 	if err != nil {
// 		panic(err)
// 	}

// 	strOpenChannelRequest, _ := json.Marshal(openChannelRequest)

// 	response, _ := c.rpcLightningNode.channel.OpenChannel(
// 		context.Background(),
// 		&pb.OpenChannelRequest{
// 			FromAddress: AAccount.AccAddress().String(),
// 			ToAddress:   multisigAddr,
// 			Signature:   strSig,
// 			Payload:     string(strOpenChannelRequest),
// 		})
// 	return response.GetResponse()
// }

// func OpenChannelFromB(channelId string) {
// 	cfg := &Config{
// 		ChainId:               "channel",
// 		Endpoint:              "http://0.0.0.0:26657",
// 		LightningNodeEndpoint: "0.0.0.0:2525",
// 		CoinType:              60,
// 		PrefixAddress:         "cosmos",
// 		TokenSymbol:           "token",
// 	}

// 	c := NewClient(cfg)
// 	acc := c.NewAccountClient()
// 	channelClient := c.NewChannelClient()

// 	AAccount, _ := acc.ImportAccount("excuse quiz oyster vendor often spray day vanish slice topic pudding crew promote floor shadow best subway slush slender good merit hollow certain repeat")
// 	BAccount, _ := acc.ImportAccount("claim market flip canoe wreck maid recipe bright fuel slender ladder album behind repeat come trophy come vicious frown prefer height unknown thank damp")

// 	channelResult, _ := c.rpcLightningNode.channel.GetChannelById(context.Background(), &pb.GetChannelRequest{Id: channelId})

// 	var payload struct {
// 		Msg      channelTypes.MsgOpenChannel
// 		GasLimit uint64
// 		GasPrice string
// 	}

// 	fmt.Println("payload", channelResult.Payload)
// 	json.Unmarshal([]byte(channelResult.Payload), &payload)

// 	tmpMultisigAddr, multiSigPubkey, _ := acc.CreateMulSigAccountFromTwoAccount(AAccount.PublicKey(), BAccount.PublicKey(), 2)
// 	fmt.Println("tmpMultisigAddr", tmpMultisigAddr) // astra1747xvksuc7ecpylckzjvmcqvvlmp6t6ujs3lld

// 	msg := channelTypes.NewMsgOpenChannel(
// 		payload.Msg.Creator,
// 		payload.Msg.PartA,
// 		payload.Msg.PartB,
// 		payload.Msg.CoinA,
// 		payload.Msg.CoinB,
// 		payload.Msg.MultisigAddr,
// 		payload.Msg.Sequence,
// 	)

// 	openChannelRequest := channel.SignMsgRequest{
// 		Msg:      msg,
// 		GasLimit: payload.GasLimit,
// 		GasPrice: payload.GasPrice,
// 	}

// 	fmt.Println("Sig1", channelResult.SignatureA)
// 	fmt.Println("openChannelRequest", openChannelRequest)

// 	signList := make([][]signingTypes.SignatureV2, 0)

// 	// 1. decode signature from A: string -> []byte
// 	signByte1, err := common.SignatureJsonDecoder(c.RpcClient().TxConfig, channelResult.SignatureA)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// 2. append signature to signList
// 	signList = append(signList, signByte1)

// 	// 3. create multisignature from B
// 	strSig2, err := channelClient.SignMultisigTxFromOneAccount(openChannelRequest, BAccount, multiSigPubkey)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("Sig2", strSig2)

// 	// 4. decode signature from B: string -> []byte
// 	signByte2, err := common.SignatureJsonDecoder(c.RpcClient().TxConfig, strSig2)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// 5. append signature to signList
// 	signList = append(signList, signByte2)

// 	fmt.Println("new tx multisign")

// 	// 6. create new transaction with multisignature
// 	newTx := common.NewMultisigTxBuilder(c.RpcClient(),
// 		nil,
// 		openChannelRequest.GasLimit,
// 		openChannelRequest.GasPrice,
// 		0,
// 		2)

// 	// 7. build unsigned transaction
// 	txBuilderMultiSign, err := newTx.BuildUnsignedTx(openChannelRequest.Msg)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// 8. create multisignature
// 	err = newTx.GenerateMultisig(txBuilderMultiSign, multiSigPubkey, signList)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// 9. encode transaction to json
// 	txJson, err := common.TxBuilderJsonEncoder(c.RpcClient().TxConfig, txBuilderMultiSign)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("rawData", string(txJson))

// 	// 10. decode transaction from json to []byte
// 	txByte, err := common.TxBuilderJsonDecoder(c.RpcClient().TxConfig, txJson)
// 	if err != nil {
// 		panic(err)
// 	}

// 	txHash := common.TxHash(txByte)
// 	fmt.Println("txHash", txHash)

// 	// 11. broadcast transaction
// 	txResult, err := c.RpcClient().BroadcastTxCommit(txByte)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("tx openchannel result code", txResult.Code)
// 	fmt.Println(txResult)
// }

// func FundFromA(amount int64, secret string, timelock uint64, seq int) (channelTypes.MsgFund, string, string) {
// 	cfg := &Config{
// 		ChainId:               "channel",
// 		Endpoint:              "http://0.0.0.0:26657",
// 		LightningNodeEndpoint: "0.0.0.0:2525",
// 		CoinType:              60,
// 		PrefixAddress:         "cosmos",
// 		TokenSymbol:           "token",
// 	}

// 	c := NewClient(cfg)
// 	acc := c.NewAccountClient()
// 	AAccount, _ := acc.ImportAccount("excuse quiz oyster vendor often spray day vanish slice topic pudding crew promote floor shadow best subway slush slender good merit hollow certain repeat")
// 	BAccount, _ := acc.ImportAccount("claim market flip canoe wreck maid recipe bright fuel slender ladder album behind repeat come trophy come vicious frown prefer height unknown thank damp")
// 	multisigAddr, multiSigPubkey, _ := acc.CreateMulSigAccountFromTwoAccount(AAccount.PublicKey(), BAccount.PublicKey(), 2)

// 	partACommitment := channelTypes.MsgFund{
// 		ChannelID: multisigAddr + ":cosmos:" + fmt.Sprint(seq),
// 		Creator:   multisigAddr,
// 		From:      AAccount.AccAddress().String(),
// 		CoinToHtlc: &types.Coin{
// 			Denom:  "token",
// 			Amount: types.NewInt(amount),
// 		},
// 		Hashcode: secret,
// 		Timelock: fmt.Sprint(timelock),
// 		Multisig: multisigAddr,
// 	}

// 	commitmentA := channel.SignMsgRequest{
// 		Msg:      &partACommitment,
// 		GasLimit: 200000,
// 		GasPrice: "0token",
// 	}

// 	channelClient := c.NewChannelClient()

// 	strSig, err := channelClient.SignMultisigTxFromOneAccount(commitmentA, AAccount, multiSigPubkey)
// 	if err != nil {
// 		panic(err)
// 	}

// 	return partACommitment, AAccount.AccAddress().String(), strSig
// }

// func AcceptFundFromB(fundCommitment channelTypes.MsgFund, aAddress string, aSignature string, secret string) (channelTypes.MsgAcceptfund, string, string) {
// 	cfg := &Config{
// 		ChainId:               "channel",
// 		Endpoint:              "http://0.0.0.0:26657",
// 		LightningNodeEndpoint: "0.0.0.0:2525",
// 		CoinType:              60,
// 		PrefixAddress:         "cosmos",
// 		TokenSymbol:           "token",
// 	}

// 	c := NewClient(cfg)
// 	acc := c.NewAccountClient()
// 	AAccount, _ := acc.ImportAccount("excuse quiz oyster vendor often spray day vanish slice topic pudding crew promote floor shadow best subway slush slender good merit hollow certain repeat")
// 	BAccount, _ := acc.ImportAccount("claim market flip canoe wreck maid recipe bright fuel slender ladder album behind repeat come trophy come vicious frown prefer height unknown thank damp")
// 	multisigAddr, multiSigPubkey, _ := acc.CreateMulSigAccountFromTwoAccount(AAccount.PublicKey(), BAccount.PublicKey(), 2)

// 	partBCommitment := channelTypes.MsgAcceptfund{
// 		ChannelID:        fundCommitment.ChannelID,
// 		Creator:          fundCommitment.Creator,
// 		From:             fundCommitment.From,
// 		CoinToAcceptSide: fundCommitment.CoinToHtlc,
// 		Hashcode:         secret,
// 		Timelock:         fundCommitment.Timelock,
// 		Multisig:         multisigAddr,
// 	}

// 	commitmentB := channel.SignMsgRequest{
// 		Msg:      &partBCommitment,
// 		GasLimit: 200000,
// 		GasPrice: "0token",
// 	}

// 	channelClient := c.NewChannelClient()

// 	strSig, err := channelClient.SignMultisigTxFromOneAccount(commitmentB, AAccount, multiSigPubkey)
// 	if err != nil {
// 		panic(err)
// 	}

// 	return partBCommitment, BAccount.AccAddress().String(), strSig
// }

// func StoreAcceptFundCommitmentFromA(acceptFundCommitment channelTypes.MsgAcceptfund, bAddress string, bSignature string) {
// 	cfg := &Config{
// 		ChainId:               "channel",
// 		Endpoint:              "http://0.0.0.0:26657",
// 		LightningNodeEndpoint: "0.0.0.0:2525",
// 		CoinType:              60,
// 		PrefixAddress:         "cosmos",
// 		TokenSymbol:           "token",
// 	}

// 	c := NewClient(cfg)
// 	acc := c.NewAccountClient()
// 	AAccount, _ := acc.ImportAccount("excuse quiz oyster vendor often spray day vanish slice topic pudding crew promote floor shadow best subway slush slender good merit hollow certain repeat")
// 	BAccount, _ := acc.ImportAccount("claim market flip canoe wreck maid recipe bright fuel slender ladder album behind repeat come trophy come vicious frown prefer height unknown thank damp")
// 	multisigAddr, _, _ := acc.CreateMulSigAccountFromTwoAccount(AAccount.PublicKey(), BAccount.PublicKey(), 2)

// 	fmt.Println(acceptFundCommitment, bAddress, bSignature)
// 	transfer(c, AAccount.PrivateKeyToString(), multisigAddr, big.NewInt(1))
// }

// func CloseChannel(amountA int64, amoutB int64, seq int) (channelTypes.MsgCloseChannel, string, string) {
// 	cfg := &Config{
// 		ChainId:               "channel",
// 		Endpoint:              "http://0.0.0.0:26657",
// 		LightningNodeEndpoint: "0.0.0.0:2525",
// 		CoinType:              60,
// 		PrefixAddress:         "cosmos",
// 		TokenSymbol:           "token",
// 	}

// 	c := NewClient(cfg)
// 	acc := c.NewAccountClient()
// 	AAccount, _ := acc.ImportAccount("excuse quiz oyster vendor often spray day vanish slice topic pudding crew promote floor shadow best subway slush slender good merit hollow certain repeat")
// 	BAccount, _ := acc.ImportAccount("claim market flip canoe wreck maid recipe bright fuel slender ladder album behind repeat come trophy come vicious frown prefer height unknown thank damp")
// 	multisigAddr, multiSigPubkey, _ := acc.CreateMulSigAccountFromTwoAccount(AAccount.PublicKey(), BAccount.PublicKey(), 2)

// 	closeChannelMsg := channelTypes.MsgCloseChannel{
// 		ChannelID: multisigAddr + ":token:" + fmt.Sprint(seq),
// 		Creator:   multisigAddr,
// 		ToA:       AAccount.AccAddress().String(),
// 		ToB:       BAccount.AccAddress().String(),
// 		CoinA: &types.Coin{
// 			Denom:  "token",
// 			Amount: types.NewInt(amountA),
// 		},
// 		CoinB: &types.Coin{
// 			Denom:  "token",
// 			Amount: types.NewInt(amoutB),
// 		},
// 	}
// 	closeChannelRequest := channel.SignMsgRequest{
// 		Msg:      &closeChannelMsg,
// 		GasLimit: 200000,
// 		GasPrice: "0token",
// 	}

// 	channelClient := c.NewChannelClient()

// 	strSig, err := channelClient.SignMultisigTxFromOneAccount(closeChannelRequest, AAccount, multiSigPubkey)
// 	if err != nil {
// 		panic(err)
// 	}

// 	return closeChannelMsg, BAccount.AccAddress().String(), strSig
// }

// func AcceptCloseChannel(closeChannel channelTypes.MsgCloseChannel, aAddress string, aSignature string) {
// 	cfg := &Config{
// 		ChainId:               "channel",
// 		Endpoint:              "http://0.0.0.0:26657",
// 		LightningNodeEndpoint: "0.0.0.0:2525",
// 		CoinType:              60,
// 		PrefixAddress:         "cosmos",
// 		TokenSymbol:           "token",
// 	}

// 	c := NewClient(cfg)
// 	acc := c.NewAccountClient()
// 	AAccount, _ := acc.ImportAccount("excuse quiz oyster vendor often spray day vanish slice topic pudding crew promote floor shadow best subway slush slender good merit hollow certain repeat")
// 	BAccount, _ := acc.ImportAccount("claim market flip canoe wreck maid recipe bright fuel slender ladder album behind repeat come trophy come vicious frown prefer height unknown thank damp")
// 	_, multiSigPubkey, _ := acc.CreateMulSigAccountFromTwoAccount(AAccount.PublicKey(), BAccount.PublicKey(), 2)

// 	acceptCloseChannelMsg := channelTypes.MsgCloseChannel{
// 		ChannelID: closeChannel.ChannelID,
// 		Creator:   closeChannel.Creator,
// 		ToA:       closeChannel.ToA,
// 		ToB:       closeChannel.ToB,
// 		CoinA:     closeChannel.CoinA,
// 		CoinB:     closeChannel.CoinB,
// 	}

// 	acceptCloseChannelRequest := channel.SignMsgRequest{
// 		Msg:      &acceptCloseChannelMsg,
// 		GasLimit: 200000,
// 		GasPrice: "0token",
// 	}

// 	channelClient := c.NewChannelClient()

// 	signList := make([][]signingTypes.SignatureV2, 0)

// 	signByte1, err := common.SignatureJsonDecoder(c.RpcClient().TxConfig, aSignature)
// 	if err != nil {
// 		panic(err)
// 	}

// 	signList = append(signList, signByte1)

// 	strSig2, err := channelClient.SignMultisigTxFromOneAccount(acceptCloseChannelRequest, BAccount, multiSigPubkey)
// 	if err != nil {
// 		panic(err)
// 	}
// 	signByte2, err := common.SignatureJsonDecoder(c.RpcClient().TxConfig, strSig2)
// 	if err != nil {
// 		panic(err)
// 	}
// 	signList = append(signList, signByte2)

// 	newTx := common.NewMultisigTxBuilder(c.RpcClient(), nil, 200000, "0token", 0, 2)
// 	txBuilderMultiSign, err := newTx.BuildUnsignedTx(acceptCloseChannelRequest.Msg)
// 	if err != nil {
// 		panic(err)
// 	}

// 	err = newTx.GenerateMultisig(txBuilderMultiSign, multiSigPubkey, signList)
// 	if err != nil {
// 		panic(err)
// 	}

// 	txJson, err := common.TxBuilderJsonEncoder(c.RpcClient().TxConfig, txBuilderMultiSign)
// 	if err != nil {
// 		panic(err)
// 	}

// 	txByte, err := common.TxBuilderJsonDecoder(c.RpcClient().TxConfig, txJson)
// 	if err != nil {
// 		panic(err)
// 	}

// 	txHash := common.TxHash(txByte)
// 	fmt.Println(txHash)

// 	txResult, err := c.RpcClient().BroadcastTxSync(txByte)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(txResult)

// }

// func transfer(c *Client, privateKey string, toAddress string, value *big.Int) {
// 	bankClient := c.NewBankClient()
// 	request := &bank.TransferRequest{
// 		PrivateKey: privateKey,
// 		Receiver:   toAddress,
// 		Amount:     value,
// 		GasLimit:   200000,
// 		GasPrice:   "0token",
// 	}

// 	txBuilder, err := bankClient.TransferRawDataWithPrivateKey(request)
// 	if err != nil {
// 		panic(err)
// 	}

// 	txJson, err := common.TxBuilderJsonEncoder(c.RpcClient().TxConfig, txBuilder)
// 	if err != nil {
// 		panic(err)
// 	}

// 	txByte, err := common.TxBuilderJsonDecoder(c.RpcClient().TxConfig, txJson)
// 	if err != nil {
// 		panic(err)
// 	}

// 	txHash := common.TxHash(txByte)
// 	fmt.Println("txHash", txHash)

// 	_, err = c.RpcClient().BroadcastTxCommit(txByte)

// 	if err != nil {
// 		panic(err)
// 	}
// }
