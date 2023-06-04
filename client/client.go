package client

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	sdkClient "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	authTypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/evmos/ethermint/encoding"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/m25-lab/channel/app"
	channeltypes "github.com/m25-lab/channel/x/channel/types"
	"github.com/m25-lab/lightning-network-node/core_chain_sdk/account"
	"github.com/m25-lab/lightning-network-node/database/models"
	"github.com/m25-lab/lightning-network-node/node"
	"google.golang.org/grpc"
)

type L1RpcClient struct {
	bank    banktypes.QueryClient
	channel channeltypes.QueryClient
}
type Client struct {
	Node      *node.LightningNode
	Bot       *tgbotapi.BotAPI
	l1Client  *L1RpcClient
	ClientCtx *client.Context
}

func New(node *node.LightningNode) (*Client, error) {
	l1Conn, err := grpc.Dial(
		node.Config.Node.Endpoint,
		grpc.WithInsecure(),
	)

	if err != nil {
		return nil, err
	}

	bot, err := tgbotapi.NewBotAPI(node.Config.Telegram.BotId)

	if err != nil {
		panic(err)
	}

	ar := authTypes.AccountRetriever{}

	//github.com/cosmos/cosmos-sdk/simapp/app.go
	//github.com/evmos/ethermint@v0.19.0/app/app.go -> selected
	encodingConfig := encoding.MakeConfig(app.ModuleBasics)
	rpcHttp, err := sdkClient.NewClientFromNode(node.Config.Corechain.Endpoint) //"http://localhost:26657")
	if err != nil {
		panic(err)
	}
	ClientCtx := sdkClient.Context{}
	ClientCtx = ClientCtx.
		WithClient(rpcHttp).
		//WithNodeURI(c.endpoint).
		WithCodec(encodingConfig.Marshaler).
		WithInterfaceRegistry(encodingConfig.InterfaceRegistry).
		WithTxConfig(encodingConfig.TxConfig).
		WithLegacyAmino(encodingConfig.Amino).
		WithAccountRetriever(ar).
		WithChainID(node.Config.Node.ChainId). //"channel").
		WithBroadcastMode(flags.BroadcastSync)

	return &Client{
		node,
		bot,
		&L1RpcClient{
			banktypes.NewQueryClient(l1Conn),
			channeltypes.NewQueryClient(l1Conn),
		},
		&ClientCtx,
	}, nil
}

func (client *Client) RunTelegramBot() error {
	// client.Bot.Debug = true

	log.Printf("Authorized on account %s", client.Bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := client.Bot.GetUpdatesChan(u)

	for update := range updates {
		flagUpdateTelmsg := false
		var message *models.Message
		var msg tgbotapi.MessageConfig

		if update.CallbackQuery != nil {
			msg = tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "")
			clientId := strconv.FormatInt(update.CallbackQuery.Message.Chat.ID, 10)
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
			if _, err := client.Bot.Request(callback); err != nil {
				panic(err)
			}

			action, messageId, _ := ParseCallbackData(update.CallbackQuery.Data)

			switch action {
			case models.AcceptAddWhitelist:
				var err error
				message, err = client.AcceptAddWhitelist(clientId, messageId)
				if err != nil {
					msg.Text = "Error: " + err.Error()
				} else {
					msg.Text = fmt.Sprintf("✅ *Add whitelist successfully.* \n Add `%s` to whitelist", message.Users[1])
					flagUpdateTelmsg = true
				}
			}
		} else if update.Message.IsCommand() {
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "")
			clientId := strconv.FormatInt(update.Message.From.ID, 10)

			switch update.Message.Command() {
			case "start", "create_account":
				account, err := client.CreateAccount(clientId)
				if err != nil {
					msg.Text = "Error: " + err.Error()
				} else {
					msg.Text = fmt.Sprintf("✅ *Account created successfully.* \n 💳 Your address: `%s` \n 🗝️ Your mnemonic: `%s` \n", account.AccAddress().String(), account.Mnemonic())
				}
			case "current_account":
				account, err := client.CurrentAccount(clientId)
				if err != nil {
					msg.Text = "Error: " + err.Error()
				} else {
					msg.Text = fmt.Sprintf("🟢 Current account: `%s`", account.AccAddress().String())
				}
			case "import_account":
				account, err := client.ImportAccount(clientId, update.Message.CommandArguments())
				if err != nil {
					msg.Text = "Error: " + err.Error()
				} else {
					msg.Text = fmt.Sprintf("✅ *Account imported successfully.* \n 💳 Your address: `%s` \n 🗝️ Your mnemonic: `%s` \n", account.AccAddress().String(), account.Mnemonic())
				}
			case "add_whitelist":
				var err error
				message, err = client.AddWhitelist(clientId, update.Message.CommandArguments())
				if err != nil {
					msg.Text = "Error: " + err.Error()
				} else {
					msg.Text = fmt.Sprintf("⬆️ *Request whitelist to* `%s`.", update.Message.CommandArguments())
					flagUpdateTelmsg = true
				}
			case "whitelist":
				whitelist, err := client.ListWhitelist(clientId)
				if err != nil {
					msg.Text = "Error: " + err.Error()
				} else {
					strWhitelist := ""
					for i, w := range whitelist {
						strWhitelist += fmt.Sprintf("%d. `%s`\n", i+1, w.PartnerAddress)
					}
					msg.Text = fmt.Sprintf("*Whitelist:* \n %s", strWhitelist)
				}
			case "balance":
				account, err := client.CurrentAccount(clientId)
				if err != nil {
					msg.Text = "Error: " + err.Error()
				}
				balance, err := client.Balance(account.AccAddress().String())
				if err != nil {
					msg.Text = "Error: " + err.Error()
				} else {
					msg.Text = fmt.Sprintf("💰 *Balance:* `%d`", balance)
				}
			case "channel_balance":
				balance, err := client.ChannelBalance(clientId, update.Message.CommandArguments())
				if err != nil {
					msg.Text = "Error: " + err.Error()
				} else {
					msg.Text = fmt.Sprintf("Channel ID: `%s` \n My balance: `%d` \n Partner balance: `%d`", update.Message.CommandArguments(), balance.MyBalance, balance.PartnerBalance)
				}
			case "transfer":
				params := strings.Split(update.Message.CommandArguments(), " ")
				amount, err := strconv.ParseInt(params[1], 10, 64)
				if err != nil {
					msg.Text = "Error: " + err.Error()
				}
				err = client.Transfer(clientId, params[0], amount)
				if err != nil {
					msg.Text = "Error: " + err.Error()
				} else {
					msg.Text = fmt.Sprintf("💸 *Transfer successfully.* \n Transfer `%d` to `%s`", amount, params[0])
				}
			case "ln_transfer":
				params := strings.Split(update.Message.CommandArguments(), " ")
				amount, err := strconv.ParseInt(params[1], 10, 64)
				if err != nil {
					msg.Text = "Error: " + err.Error()
				}
				err = client.LnTransfer(clientId, params[0], amount, nil, nil)
				if err != nil {
					msg.Text = "Error: " + err.Error()
				} else {
					msg.Text = fmt.Sprintf("⚡ *Transfer successfully.* \n Transfer `%d` to `%s`", amount, params[0])
				}
			case "ln_transfer_multi":
				params := strings.Split(update.Message.CommandArguments(), " ")
				amount, err := strconv.ParseInt(params[1], 10, 64)
				if err != nil {
					msg.Text = "Error: " + err.Error()
				}
				err = client.LnTransferMulti(clientId, params[0], amount, nil, nil)
				if err != nil {
					msg.Text = "Error: " + err.Error()
				} else {
					msg.Text = fmt.Sprintf("⚡ *Transfer successfully.* \n Transfer `%d` to `%s`", amount, params[0])
				}
			//TODO: case for withdraw and broadcast
			//TODO: new broadcast-only Model for fwdCommit
			//

			default:
				msg.Text = "I don't know that command"
			}
		} else {
			continue
		}
		msg.ParseMode = "Markdown"
		telMsg, err := client.Bot.Send(msg)
		if err != nil {
			return err
		}

		if flagUpdateTelmsg {
			client.Node.Repository.Message.UpdateTelegramChatId(context.Background(), message.ID, telMsg.MessageID)
		}
	}

	return nil
}

func (client *Client) TelegramMsg(_clientId string, msg *models.Message, fromAccount *account.PrivateKeySerialized, toAccount *account.PKAccount) error {
	clientId, err := strconv.ParseInt(_clientId, 10, 64)
	if err != nil {
		return err
	}

	switch msg.Action {
	case models.AddWhitelist:
		return client.ResolveAddWhitelist(clientId, msg)
	case models.AcceptAddWhitelist:
		return client.ResolveAcceptAddWhitelist(clientId, msg)
	}

	return nil
}
