package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Whitelist struct {
	ID           primitive.ObjectID `bson:"_id, omitempty"`
	Users        []string           `bson:"users"`
	Pubkeys      []string           `bson:"pubkeys"`
	MultiAddress string             `bson:"multi_address"`
	MultiPubkey  string             `bson:"multi_pubkey"`
}
