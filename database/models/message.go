package models

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	ExchangeHashcode   string = "exchange_hashcode"
	ExchangeCommitment string = "create_commitment"
	OpenChannel        string = "open_channel"
	CloseChannel       string = "close_channel"
	AddFund            string = "add_fund"
	AddWhitelist       string = "add_whitelist"
	AcceptAddWhitelist string = "accept_add_whitelist"
)

type AddWhitelistData struct {
	Pubkey string `json:"pubkey"`
}

type CreateCommitmentData struct {
	Creator           string `json:"creator"`
	ChannelID         string `json:"channel_id"`
	From              string `json:"from"`
	Timelock          uint64 `json:"timelock"`
	ToTimelockAddr    string `json:"to_timelock_addr"`
	ToHashlockAddr    string `json:"to_hashlock_addr"`
	CoinToCreator     int64  `json:"coin_to_creator"`
	CoinToHtlc        int64  `json:"coin_to_htlc"`
	Hashcode          string `json:"hashcode"`
	PartnerSignature  string `json:"partner_signature"`
	OwnSignature      string `json:"own_signature"`
	FwdDest           string `json:"fwd_dest,omitempty"`
	HashcodeDest      string `json:"hashcode_dest,omitempty"`
	IsFirstCommitment bool   `json:"is_first_commitment"`
}

type OpenChannelData struct {
	StrSig string `json:"str_sig"`
}

type ExchangeHashcodeData struct {
	MySecret        string `bson:"my_secret" json:"my_secret"`
	MyHashcode      string `bson:"my_hashcode" json:"my_hashcode"`
	PartnerHashcode string `bson:"partner_hashcode" json:"partner_hashcode"`
	PartnerSecret   string `bson:"partner_secret,omitempty" json:"partner_secret,omitempty"`
	ChannelID       string `bson:"channel_id,omitempty" json:"channel_id,omitempty"`
}

type Message struct {
	ID              primitive.ObjectID `bson:"_id, omitempty"`
	OriginalID      primitive.ObjectID `bson:"original_id, omitempty"`
	ChannelID       string             `bson:"channel_id"`
	Action          string             `bson:"action"`
	Data            string             `bson:"data"`
	Owner           string             `bson:"owner"`
	Users           []string           `bson:"users"`
	TelegramChatId  int64              `bson:"telegram_chat_id"`
	ReliedMessageId string             `bson:"relied_message_id"`
	IsReplied       bool               `bson:"is_replied"`
}
