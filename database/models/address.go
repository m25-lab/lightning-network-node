package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Address struct {
	ID       primitive.ObjectID `bson:"_id, omitempty"`
	Address  string             `bson:"address"`
	Pubkey   string             `bson:"pubkey"`
	Mnemonic string             `bson:"mnemonic"`
	ClientId string             `bson:"client_id"`
}
