package client

import (
	"context"
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/m25-lab/lightning-network-node/database/models"
	"github.com/m25-lab/lightning-network-node/node"
)

type Client struct {
	Node *node.LightningNode
	Bot  *tgbotapi.BotAPI
}

func New(node *node.LightningNode) (*Client, error) {
	bot, err := tgbotapi.NewBotAPI(node.Config.Telegram.BotId)

	if err != nil {
		panic(err)
	}

	return &Client{
		node,
		bot,
	}, nil
}

func (client *Client) RunTelegramBot() error {
	// client.Bot.Debug = true

	log.Printf("Authorized on account %s", client.Bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := client.Bot.GetUpdatesChan(u)

	for update := range updates {
		if update.CallbackQuery != nil {
			flagUpdateTelmsg := false
			var message *models.Message

			clientId := strconv.FormatInt(update.CallbackQuery.Message.Chat.ID, 10)
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
			if _, err := client.Bot.Request(callback); err != nil {
				panic(err)
			}

			msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "")

			action, messageId, err := client.ParseCallbackData(update.CallbackQuery.Data)
			if err != nil {
				msg.Text = "Error: " + err.Error()
			} else {
				switch action {
				case models.AcceptAddWhitelist:
					message, err = client.AcceptAddWhitelist(clientId, messageId)
					if err != nil {
						msg.Text = "Error: " + err.Error()
					} else {
						msg.Text = fmt.Sprintf("✅ *Add whitelist successfully.* \n Add `%s` to whitelist", message.Users[0])
						flagUpdateTelmsg = true
					}
				}
			}

			msg.ParseMode = "Markdown"
			telMsg, err := client.Bot.Send(msg)
			if err != nil {
				panic(err)
			}
			if flagUpdateTelmsg {
				client.Node.Repository.Message.UpdateTelegramChatId(context.Background(), message.ID, telMsg.MessageID)
			}
		} else if update.Message.IsCommand() {
			flagUpdateTelmsg := false
			var message *models.Message

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
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
			default:
				msg.Text = "I don't know that command"
			}

			msg.ParseMode = "Markdown"
			telMsg, err := client.Bot.Send(msg)
			if err != nil {
				log.Panic(err)
			}

			if flagUpdateTelmsg {
				client.Node.Repository.Message.UpdateTelegramChatId(context.Background(), message.ID, telMsg.MessageID)
			}
		}
	}

	return nil
}

func (client *Client) TelegramMsg(_clientId string, msg *models.Message) error {
	clientId, err := strconv.ParseInt(_clientId, 10, 64)
	if err != nil {
		return err
	}

	if msg.Action == models.AddWhitelist {
		return client.ResolveAddWhitelist(clientId, msg)
	} else if msg.Action == models.AcceptAddWhitelist {
		return client.ResolveAcceptAddWhitelist(clientId, msg)
	}

	return nil
}
