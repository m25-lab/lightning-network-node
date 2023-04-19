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
	rpcHttp, err := sdkClient.NewClientFromNode("http://localhost:26657")
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
		WithChainID("channel").
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
			case models.AcceptRequestInvoice:
				var err error
				invoice, err := client.AcceptRequestInvoice(clientId, messageId)
				if err != nil {
					msg.Text = "Error: " + err.Error()
				} else {
					msg.Text = fmt.Sprintf("✅ *Generate invoice successfully.* \n From %s - Amount %d - Hash %s", invoice.From, invoice.Amount, invoice.HashSecret)
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
				fmt.Println("contact: ", update.Message.CommandArguments())
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
				err = client.LnTransfer(clientId, params[0], amount)
				if err != nil {
					msg.Text = "Error: " + err.Error()
				} else {
					msg.Text = fmt.Sprintf("⚡ *Transfer successfully.* \n Transfer `%d` to `%s`", amount, params[0])
				}
			case "request_invoice":
				params := strings.Split(update.Message.CommandArguments(), " ")
				lastReceiverAddress := params[0]
				if lastReceiverAddress == "" {
					msg.Text = "Error: Required last receiver address"
				}
				amount, err := strconv.ParseInt(params[1], 10, 64)
				if err != nil {
					msg.Text = "Error: " + err.Error()
				}
				invoice, err := client.CreateInvoice(clientId, lastReceiverAddress, amount)
				if err != nil {
					msg.Text = "Error: " + err.Error()
				} else {
					msg.Text = fmt.Sprintf("⬆️ *Create invoice successfully.* \n Invoice %d to %s", invoice.Amount, invoice.To)
				}
			case "invoice_for":
				//params := strings.Split(update.Message.CommandArguments(), " ")
				//firstSenderID, err := strconv.ParseInt(params[0], 10, 64)
				//if err != nil {
				//	msg.Text = "Error: " + err.Error()
				//}
				//amount, err := strconv.ParseInt(params[1], 10, 64)
				//if err != nil {
				//	msg.Text = "Error: " + err.Error()
				//}
				//invoice, err := client.CreateInvoice(clientId, firstSenderID, amount)
				//if err != nil {
				//	msg.Text = "Error: " + err.Error()
				//} else {
				//	msg.Text = fmt.Sprintf("⚡ *Create invoice successfully.* \n Invoice: `%s`", invoice)
				//}
			// case "ln_list_invoices":
			// 	invoices, err := client.LnListInvoices(clientId)
			// 	if err != nil {
			// 		msg.Text = "Error: " + err.Error()
			// 	} else {
			// 		strInvoices := ""
			// 		for i, p := range invoices {
			// 			strInvoices += fmt.Sprintf("%d. `%s` \n", i+1, p.PaymentRequest)
			// 		}
			// 		msg.Text = fmt.Sprintf("*Invoices:* \n %s", strInvoices)
			// 	}
			// case "forward_commit":
			// 	params := strings.Split(update.Message.CommandArguments(), " ")
			// 	amount, err := strconv.ParseInt(params[1], 10, 64)
			// 	if err != nil {
			// 		msg.Text = "Error: " + err.Error()
			// 	}
			// 	err = client.ForwardCommit(clientId, params[0], amount)
			// 	if err != nil {
			// 		msg.Text = "Error: " + err.Error()
			// 	} else {
			// 		msg.Text = fmt.Sprintf("⚡ *Forward commit successfully.* \n Forward commit `%d` to `%s`", amount, params[0])
			// 	}
			// case "ln_close_channel":
			// 	err := client.LnCloseChannel(clientId, update.Message.CommandArguments())
			// 	if err != nil {
			// 		msg.Text = "Error: " + err.Error()
			// 	} else {
			// 		msg.Text = fmt.Sprintf("⚡ *Close channel successfully.* \n Close channel `%s`", update.Message.CommandArguments())
			// 	}
			// case "ln_list_channels":
			// 	channels, err := client.LnListChannels(clientId)
			// 	if err != nil {
			// 		msg.Text = "Error: " + err.Error()
			// 	} else {
			// 		strChannels := ""
			// 		for i, c := range channels {
			// 			strChannels += fmt.Sprintf("%d. `%s` \n", i+1, c.ChannelId)
			// 		}
			// 		msg.Text = fmt.Sprintf("*Channels:* \n %s", strChannels)
			// 	}
			// case "ln_list_peers":
			// 	peers, err := client.LnListPeers(clientId)
			// 	if err != nil {
			// 		msg.Text = "Error: " + err.Error()
			// 	} else {
			// 		strPeers := ""
			// 		for i, p := range peers {
			// 			strPeers += fmt.Sprintf("%d. `%s` \n", i+1, p.PubKey)
			// 		}
			// 		msg.Text = fmt.Sprintf("*Peers:* \n %s", strPeers)
			// 	}
			// case "ln_get_info":
			// 	info, err := client.LnGetInfo(clientId)
			// 	if err != nil {
			// 		msg.Text = "Error: " + err.Error()
			// 	} else {
			// 		msg.Text = fmt.Sprintf("*Info:* \n %s", info)
			// 	}

			// case "ln_get_transactions":
			// 	transactions, err := client.LnGetTransactions(clientId)
			// 	if err != nil {
			// 		msg.Text = "Error: " + err.Error()
			// 	} else {
			// 		msg.Text = fmt.Sprintf("*Transactions:* \n %s", transactions)
			// 	}
			// case "ln_get_node":
			// 	node, err := client.LnGetNode(clientId, update.Message.CommandArguments())
			// 	if err != nil {
			// 		msg.Text = "Error: " + err.Error()
			// 	} else {
			// 		msg.Text = fmt.Sprintf("*Node:* \n %s", node)
			// 	}
			// case "ln_get_route":
			// 	route, err := client.LnGetRoute(clientId, update.Message.CommandArguments())
			// 	if err != nil {
			// 		msg.Text = "Error: " + err.Error()
			// 	} else {
			// 		msg.Text = fmt.Sprintf("*Route:* \n %s", route)
			// 	}
			// case "ln_get_forwarding_history":
			// 	history, err := client.LnGetForwardingHistory(clientId)
			// 	if err != nil {
			// 		msg.Text = "Error: " + err.Error()
			// 	} else {
			// 		msg.Text = fmt.Sprintf("*Forwarding history:* \n %s", history)
			// 	}
			// case "ln_get_network_graph":
			// 	graph, err := client.LnGetNetworkGraph(clientId)
			// 	if err != nil {
			// 		msg.Text = "Error: " + err.Error()
			// 	} else {
			// 		msg.Text = fmt.Sprintf("*Network graph:* \n %s", graph)
			// 	}
			// case "ln_get_channel_graph":
			// 	graph, err := client.LnGetChannelGraph(clientId)
			// 	if err != nil {
			// 		msg.Text = "Error: " + err.Error()
			// 	} else {
			// 		msg.Text = fmt.Sprintf("*Channel graph:* \n %s", graph)
			// 	}
			// case "ln_get_node_info":
			// 	info, err := client.LnGetNodeInfo(clientId, update.Message.CommandArguments())
			// 	if err != nil {
			// 		msg.Text = "Error: " + err.Error()
			// 	} else {
			// 		msg.Text = fmt.Sprintf("*Node info:* \n %s", info)
			// 	}
			// case "ln_get_chan_info":
			// 	info, err := client.LnGetChanInfo(clientId, update.Message.CommandArguments())
			// 	if err != nil {
			// 		msg.Text = "Error: " + err.Error()
			// 	} else {
			// 		msg.Text = fmt.Sprintf("*Channel info:* \n %s", info)
			// 	}
			// case "ln_get_closed_channels":
			// 	channels, err := client.LnGetClosedChannels(clientId)
			// 	if err != nil {
			// 		msg.Text = "Error: " + err.Error()
			// 	} else {
			// 		msg.Text = fmt.Sprintf("*Closed channels:* \n %s", channels)
			// 	}
			// case "ln_get_pending_channels":
			// 	channels, err := client.LnGetPendingChannels(clientId)
			// 	if err != nil {
			// 		msg.Text = "Error: " + err.Error()
			// 	} else {
			// 		msg.Text = fmt.Sprintf("*Pending channels:* \n %s", channels)
			// 	}
			// case "ln_get_info":
			// 	info, err := client.LnGetInfo(clientId)
			// 	if err != nil {
			// 		msg.Text = "Error: " + err.Error()
			// 	} else {
			// 		msg.Text = fmt.Sprintf("*Info:* \n %s", info)
			// 	}
			// case "ln_get_fees":
			// 	fees, err := client.LnGetFees(clientId)
			// 	if err != nil {
			// 		msg.Text = "Error: " + err.Error()
			// 	} else {
			// 		msg.Text = fmt.Sprintf("*Fees:* \n %s", fees)
			// 	}
			// case "ln_get_policy":
			// 	policy, err := client.LnGetPolicy(clientId, update.Message.CommandArguments())
			// 	if err != nil {
			// 		msg.Text = "Error: " + err.Error()
			// 	} else {
			// 		msg.Text = fmt.Sprintf("*Policy:* \n %s", policy)
			// 	}
			// case "ln_get_transactions":
			// 	transactions, err := client.LnGetTransactions(clientId)
			// 	if err != nil {
			// 		msg.Text = "Error: " + err.Error()
			// 	} else {
			// 		msg.Text = fmt.Sprintf("*Transactions:* \n %s", transactions)
			// 	}
			default:
				msg.Text = "I don't know that command"
			}
		} else {
			continue
		}
		msg.ParseMode = "Markdown"
		fmt.Println("Message reply: ", msg)
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
