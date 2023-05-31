package models

const (
	SenderCommit   string = "sender_commitment"
	ReceiverCommit string = "receiver_commitment"
)

// To Save in Db of Dest
type InvoiceData struct {
	Amount int64  `json:"amount" bson:"amount"`
	From   string `json:"from" bson:"from"`
	To     string `json:"to" bson:"to"`
	Hash   string `json:"hash" bson:"hash"`
	Secret string `json:"secret" bson:"secret"`
}

// TODO: update field name to match in chanel
type SenderCommitment struct {
	Creator          string `json:"creator" bson:"creator"`
	From             string `json:"from" bson:"from"`
	Channelid        string `json:"channel_id" bson:"channel_id"`
	CoinToSender     int64  `json:"coin_to_sender" bson:"coin_to_sender"`
	CoinToHTLC       int64  `json:"coin_to_htlc" bson:"coin_to_htlc"`
	HashcodeHTLC     string `json:"hashcode_htlc" bson:"hashcode_htlc"`
	TimelockHTLC     string `json:"timelock_htlc" bson:"timelock_htlc"`
	CoinTransfer     int64  `json:"coin_transfer" bson:"coin_transfer"`
	HashcodeDest     string `json:"hashcode_dest" bson:"hashcode_dest"`
	TimelockReceiver string `json:"timelock_receiver" bson:"timelock_receiver"`
	Multisig         string `json:"multisig" bson:"multisig"`
}

type ReceiverCommitment struct {
	Creator        string `json:"creator" bson:"creator"`
	From           string `json:"from" bson:"from"`
	ChannelID      string `json:"channel_id" bson:"channel_id"`
	CoinToReceiver int64  `json:"coin_to_receiver" bson:"coin_to_receiver"`
	CoinToHTLC     int64  `json:"coin_to_htlc" bson:"coin_to_htlc"`
	HashcodeHTLC   string `json:"hashcode_htlc" bson:"hashcode_htlc"`
	TimelockHTLC   string `json:"timelock_htlc" bson:"timelock_htlc"`
	CoinTransfer   int64  `json:"coin_transfer" bson:"coin_transfer"`
	HashcodeDest   string `json:"hashcode_dest" bson:"hashcode_dest"`
	TimelockSender string `json:"timelock_sender" bson:"timelock_sender"`
	Multisig       string `json:"multisig" bson:"multisig"`
}

type FwdMessage struct {
	Action       string `bson:"action" json:"action"`
	PartnerSig   string `bson:"partner_sig" json:"partner_sig"`
	OwnSig       string `bson:"own_sig" json:"own_sig"`
	Data         string `bson:"data" json:"data"`
	From         string `bson:"from" json:"from"`
	To           string `bson:"to" json:"to"`
	HashcodeDest string `bson:"hashcode_dest" json:"hashcode_dest"`
}

type FwdSecret struct {
	HashcodeDest string `bson:"hashcode_dest" json:"hashcode_dest"`
	Secret       string `bson:"secret" json:"secret"`
}