package models

type RREQMessage struct {
	Origin      string `bson:"origin"`
	Destination string `bson:"destination"`
	Sequence    int    `bson:"sequence"`
	Hops        int    `bson:"hops"`
}

type RREPMessage struct {
	Origin      string `bson:"origin"`
	Destination string `bson:"destination"`
	Sequence    int    `bson:"sequence"`
	Hops        int    `bson:"hops"`
}

type RoutingEntry struct {
	Dest         string `bson:"dest" json:"dest"`
	HashcodeDest string `bson:"hashcode_dest" json:"hashcode_dest"`
	Next         string `bson:"next" json:"next"`
}

type SenderCommitment struct {
	Creator          string `json:"creator" bson:"creator"`
	From             string `json:"from" bson:"from"`
	ChannelID        string `json:"channel_id" bson:"channel_id"`
	CoinToSender     int64  `json:"coin_to_sender" bson:"coin_to_sender"`
	CoinToHTLC       int64  `json:"coin_to_htlc" bson:"coin_to_htlc"`
	HashcodeHTLC     string `json:"hashcode_htlc" bson:"hashcode_htlc"`
	TimelockHTLC     uint64 `json:"timelock_htlc" bson:"timelock_htlc"`
	CoinTransfer     int64  `json:"coin_transfer" bson:"coin_transfer"`
	HashcodeDest     string `json:"hashcode_dest" bson:"hashcode_dest"`
	TimelockReceiver uint64 `json:"timelock_receiver" bson:"timelock_receiver"`
	Multisig         string `json:"multisig" bson:"multisig"`
}

type ReceiverCommitment struct {
	Creator        string `json:"creator" bson:"creator"`
	From           string `json:"from" bson:"from"`
	ChannelID      string `json:"channel_id" bson:"channel_id"`
	CoinToReceiver int64  `json:"coin_to_receiver" bson:"coin_to_receiver"`
	CoinToHTLC     int64  `json:"coin_to_htlc" bson:"coin_to_htlc"`
	HashcodeHTLC   string `json:"hashcode_htlc" bson:"hashcode_htlc"`
	TimelockHTLC   uint64 `json:"timelock_htlc" bson:"timelock_htlc"`
	CoinTransfer   int64  `json:"coin_transfer" bson:"coin_transfer"`
	HashcodeDest   string `json:"hashcode_dest" bson:"hashcode_dest"`
	TimelockSender uint64 `json:"timelock_sender" bson:"timelock_sender"`
	Multisig       string `json:"multisig" bson:"multisig"`
}

// TODO: Proto
type FwdMessage struct {
	MsgSender string `bson:"msg_sender" json:"msg_sender"`
	Sig       string `bson:"sig"`
	Data      string `bson:"data" json:"data"`
}

type FwdSecret struct {
	HashcodeDest string `bson:"hashcode_dest" json:"hashcode_dest"`
	Secret       string `bson:"secret" json:"secret"`
}
