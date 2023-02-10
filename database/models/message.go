package models

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	OpenChannel  string = "open_channel"
	CloseChannel string = "close_channel"
	AddFund      string = "add_fund"
)

type Message struct {
	ID     primitive.ObjectID `bson:"_id, omitempty"`
	Action string             `bson:"action"`
	Data   string             `bson:"data"`
	Users  []string           `bson:"users"`
}
