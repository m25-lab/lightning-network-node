package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type OpenChannelRequest struct {
	ID          primitive.ObjectID  `bson:"_id, omitempty"`
	Status      string              `bson:"status"`
	FromAddress string              `bson:"from_address"`
	ToAddress   string              `bson:"to_address"`
	Signatures  []string            `bson:"signatures"`
	Payload     interface{}         `bson:"payload"`
	CreatedAt   primitive.Timestamp `bson:"created_at"`
}

type CloseChannelRequest struct {
	ID          primitive.ObjectID  `bson:"_id, omitempty"`
	Status      string              `bson:"status"`
	FromAddress string              `bson:"from_address"`
	ToAddress   string              `bson:"to_address"`
	Signatures  []string            `bson:"signatures"`
	Payload     interface{}         `bson:"payload"`
	CreatedAt   primitive.Timestamp `bson:"created_at"`
}
