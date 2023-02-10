package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Commitment struct {
	ID          primitive.ObjectID  `bson:"_id, omitempty"`
	ChannelID   string              `bson:"channel_id"`
	FromAddress string              `bson:"from_address"`
	Payload     string              `bson:"payload"`
	Signature   string              `bson:"to_signature"`
	CreatedAt   primitive.Timestamp `bson:"created_at"`
	UpdatedAt   primitive.Timestamp `bson:"updated_at"`
}
