package client

import (
	"fmt"
	"github.com/AstraProtocol/astra-go-sdk/account"
	channelTypes "github.com/AstraProtocol/channel/x/channel/types"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/m25-lab/lightning-network-node/internal/channel"
)

func OpenChannel() {
	c := NewClient()
	acc := c.NewAccountClient()
	accountTiki, _ := acc.ImportAccount("gadget final blue appear hero retire wild account message social health hobby decade neglect common egg cruel certain phrase myself alert enlist brother sure")
	privateNewAccount := initAccount(acc, c)
	account2, _ := acc.ImportPrivateKey(privateNewAccount)
	fmt.Println("accountTiki:", accountTiki.AccAddress().String())
	fmt.Println("account2:", account2.AccAddress().String())
	fmt.Println("privateNewAccount", privateNewAccount)

	multisigAddr, multiSigPubkey, _ := acc.CreateMulSignAccountFromTwoAccount(accountTiki.PublicKey(), account2.PublicKey(), 2)
	//transfer(c, accountTiki.PrivateKeyToString(), multisigAddr)
	fmt.Println("multisigAddr", multisigAddr)

	openChannelRequest := channel.SignMsgRequest{
		Msg: channelTypes.NewMsgOpenChannel(
			multisigAddr,
			accountTiki.AccAddress().String(),
			account2.AccAddress().String(),
			&types.Coin{
				Denom:  "cosmos",
				Amount: types.NewInt(10),
			},
			&types.Coin{
				Denom:  "cosmos",
				Amount: types.NewInt(10),
			},
			multisigAddr,
			"1",
		),
		GasLimit: 200000,
		GasPrice: "0cosmos",
	}

	channelClient := c.NewChannelClient()

	msg, strSig, err := channelClient.CreateMultisigMsg(openChannelRequest, accountTiki, multiSigPubkey)
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

func initAccount(acc *account.Account, c *Client) string {
	accountMain, err := acc.ImportPrivateKey("4F4FF288768511CE723C9AC765398C981C2492E72B69971DE5303F7B9D12CF2E")
	if err != nil {
		panic(err)
	}
	fmt.Println("main account", accountMain.AccAddress().String())
	privKey2, err := acc.CreateAccount()
	account2, err := acc.ImportPrivateKey(privKey2.PrivateKeyToString())
	fmt.Println("new account", account2.AccAddress().String())
	return privKey2.PrivateKeyToString()
}
