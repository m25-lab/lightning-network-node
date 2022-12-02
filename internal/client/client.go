package client

import (
	"fmt"
	"github.com/AstraProtocol/astra-go-sdk/account"
	"github.com/AstraProtocol/astra-go-sdk/bank"
	"github.com/AstraProtocol/astra/v2/app"
	sdkClient "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/types"
	authTypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/evmos/ethermint/encoding"
	ethermintTypes "github.com/evmos/ethermint/types"
	"github.com/m25-lab/lightning-network-node/internal/channel"
)

type Client struct {
	coinType      uint32
	prefixAddress string
	tokenSymbol   string
	rpcClient     sdkClient.Context
}

type Config struct {
	ChainId       string `json:"chain_id,omitempty"`
	Endpoint      string `json:"endpoint,omitempty"`
	CoinType      uint32 `json:"coin_type,omitempty"`
	PrefixAddress string `json:"prefix_address,omitempty"`
	TokenSymbol   string `json:"token_symbol,omitempty"`
}

func NewClient() *Client {
	cfg := &Config{
		ChainId:       "astra_11110-1",
		Endpoint:      "http://0.0.0.0:26657",
		CoinType:      60,
		PrefixAddress: "cosmos",
		TokenSymbol:   "cosmos",
	}
	fmt.Println(cfg)
	client := new(Client)
	client.Init(cfg)
	return client
}

func (c *Client) Init(cfg *Config) {
	c.coinType = cfg.CoinType
	c.prefixAddress = cfg.PrefixAddress
	c.tokenSymbol = cfg.TokenSymbol

	sdkConfig := types.GetConfig()
	sdkConfig.SetPurpose(44)

	switch cfg.CoinType {
	case 60:
		sdkConfig.SetCoinType(ethermintTypes.Bip44CoinType)
	case 118:
		sdkConfig.SetCoinType(types.CoinType)
	default:
		panic("Coin type invalid!")
	}

	bech32PrefixAccAddr := fmt.Sprintf("%v", c.prefixAddress)
	bech32PrefixAccPub := fmt.Sprintf("%vpub", c.prefixAddress)
	bech32PrefixValAddr := fmt.Sprintf("%vvaloper", c.prefixAddress)
	bech32PrefixValPub := fmt.Sprintf("%vvaloperpub", c.prefixAddress)
	bech32PrefixConsAddr := fmt.Sprintf("%vvalcons", c.prefixAddress)
	bech32PrefixConsPub := fmt.Sprintf("%vvalconspub", c.prefixAddress)

	sdkConfig.SetBech32PrefixForAccount(bech32PrefixAccAddr, bech32PrefixAccPub)
	sdkConfig.SetBech32PrefixForValidator(bech32PrefixValAddr, bech32PrefixValPub)
	sdkConfig.SetBech32PrefixForConsensusNode(bech32PrefixConsAddr, bech32PrefixConsPub)

	ar := authTypes.AccountRetriever{}

	//github.com/cosmos/cosmos-sdk/simapp/app.go
	//github.com/evmos/ethermint@v0.19.0/app/app.go -> selected
	encodingConfig := encoding.MakeConfig(app.ModuleBasics)
	rpcHttp, err := sdkClient.NewClientFromNode(cfg.Endpoint)
	if err != nil {
		panic(err)
	}

	rpcClient := sdkClient.Context{}
	rpcClient = rpcClient.
		WithClient(rpcHttp).
		//WithNodeURI(c.endpoint).
		WithCodec(encodingConfig.Marshaler).
		WithInterfaceRegistry(encodingConfig.InterfaceRegistry).
		WithTxConfig(encodingConfig.TxConfig).
		WithLegacyAmino(encodingConfig.Amino).
		WithAccountRetriever(ar).
		WithChainID(cfg.ChainId).
		WithBroadcastMode(flags.BroadcastSync)

	c.rpcClient = rpcClient
}

func (c *Client) NewAccountClient() *account.Account {
	return account.NewAccount(c.coinType)
}
func (c *Client) NewChannelClient() *channel.Channel {
	return channel.NewChannel(c.rpcClient)
}
func (c *Client) NewBankClient() *bank.Bank {
	return bank.NewBank(c.rpcClient, c.tokenSymbol, c.coinType)
}
func (c *Client) RpcClient() sdkClient.Context {
	return c.rpcClient
}
