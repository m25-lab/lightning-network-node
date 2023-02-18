package models

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	OpenChannel        string = "open_channel"
	CloseChannel       string = "close_channel"
	AddFund            string = "add_fund"
	AddWhitelist       string = "add_whitelist"
	AcceptAddWhitelist string = "accept_add_whitelist"
)

type AddWhitelistData struct {
	Pubkey string `json:"pubkey"`
}

type Message struct {
	ID             primitive.ObjectID `bson:"_id, omitempty"`
	OriginalID     primitive.ObjectID `bson:"original_id, omitempty"`
	ChannelID      string             `bson:"channel_id"`
	Action         string             `bson:"action"`
	Data           string             `bson:"data"`
	Owner          string             `bson:"owner"`
	Users          []string           `bson:"users"`
	TelegramChatId int64              `bson:"telegram_chat_id"`
}
