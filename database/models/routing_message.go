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

type SenderCommitment struct {
	Creator          string `json:"creator"`
	From             string `json:"from"`
	ChannelID        string `json:"channel_id"`
	CoinToSender     int64  `json:"coin_to_sender"`
	CoinToHTLC       int64  `json:"coin_to_htlc"`
	HashcodeHTLC     string `json:"hashcode_htlc"`
	TimelockHTLC     uint64 `json:"timelock_htlc"`
	CoinTransfer     int64  `json:"coin_transfer"`
	HashcodeDest     string `json:"hashcode_dest"`
	TimelockReceiver uint64 `json:"timelock_receiver"`
	Multisig         string `json:"multisig"`
	SenderSignature  string `json:"sender_signature"`
}

type ReceiverCommitment struct {
	Creator           string `json:"creator"`
	From              string `json:"from"`
	ChannelID         string `json:"channel_id"`
	CoinToReceiver    int64  `json:"coin_to_receiver"`
	CoinToHTLC        int64  `json:"coin_to_htlc"`
	HashcodeHTLC      string `json:"hashcode_htlc"`
	TimelockHTLC      uint64 `json:"timelock_htlc"`
	CoinTransfer      int64  `json:"coin_transfer"`
	HashcodeDest      string `json:"hashcode_dest"`
	TimelockSender    uint64 `json:"timelock_sender"`
	Multisig          string `json:"multisig"`
	ReceiverSignature string `json:"receiver_signature"`
}
