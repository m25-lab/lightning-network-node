package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Whitelist struct {
	ID             primitive.ObjectID `bson:"_id, omitempty"`
	Owner          string             `bson:"owner"`
	PartnerAddress string             `bson:"partner_address"`
	PartnerPubkey  string             `bson:"partner_pubkey"`
	MultiAddress   string             `bson:"multi_address"`
	MultiPubkey    string             `bson:"multi_pubkey"`
}
