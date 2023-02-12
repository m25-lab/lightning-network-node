package client

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/m25-lab/lightning-network-node/config"
	"github.com/m25-lab/lightning-network-node/node"
)

type Client struct {
	Node   *node.LightningNode
	Config *config.Config
}

func New(node *node.LightningNode, config *config.Config) (*Client, error) {
	return &Client{
		node,
		config,
	}, nil
}

func (client *Client) RunTelegramBot() error {
	bot, err := tgbotapi.NewBotAPI(client.Config.Telegram.BotId)
	if err != nil {
		panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		if !update.Message.IsCommand() { // ignore any non-command Messages
			continue
		}

		// Create a new MessageConfig. We don't have text yet,
		// so we leave it empty.
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		// Extract the command from the Message.
		switch update.Message.Command() {
		case "create_account":
			account, err := client.CreateAccount(strconv.FormatInt(update.Message.From.ID, 10))
			if err != nil {
				msg.Text = "Error: " + err.Error()
			} else {
				msg.Text = fmt.Sprintf("✅ *Account created successfully.* \n 💳 Your address: `%s` \n 🗝️ Your mnemonic: `%s` \n", account.AccAddress().String(), account.Mnemonic())
			}
		case "sayhi":
			msg.Text = "Hi :)"
		case "status":
			msg.Text = "I'm ok."
		default:
			msg.Text = "I don't know that command"
		}

		msg.ParseMode = "Markdown"
		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}

	return nil
}
