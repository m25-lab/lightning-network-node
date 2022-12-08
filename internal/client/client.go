package client

import (
	"fmt"
	"github.com/AstraProtocol/astra/v2/app"
	sdkClient "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/types"
	authTypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/evmos/ethermint/encoding"
	ethermintTypes "github.com/evmos/ethermint/types"
	"github.com/m25-lab/lightning-network-node/internal/account"
	"github.com/m25-lab/lightning-network-node/internal/bank"
	"github.com/m25-lab/lightning-network-node/internal/channel"
	"github.com/m25-lab/lightning-network-node/node/rpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	coinType         uint32
	prefixAddress    string
	tokenSymbol      string
	rpcLightningNode RpcLightningNode
	rpcClient        sdkClient.Context
}

type Config struct {
	ChainId               string `json:"chain_id,omitempty"`
	Endpoint              string `json:"endpoint,omitempty"`
	CoinType              uint32 `json:"coin_type,omitempty"`
	PrefixAddress         string `json:"prefix_address,omitempty"`
	TokenSymbol           string `json:"token_symbol,omitempty"`
	LightningNodeEndpoint string `json:"lightning_node_endpoint,omitempty"`
}

type RpcLightningNode struct {
	nodeInfo pb.NodeServiceClient
	channel  pb.ChannelServiceClient
}

func NewClient(cfg *Config) *Client {
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

	conn, err := grpc.Dial(cfg.LightningNodeEndpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	c.rpcLightningNode = RpcLightningNode{
		nodeInfo: pb.NewNodeServiceClient(conn),
		channel:  pb.NewChannelServiceClient(conn),
	}
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
