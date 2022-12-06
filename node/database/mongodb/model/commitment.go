package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Commitment struct {
	ID            primitive.ObjectID  `bson:"_id, omitempty"`
	ChannelID     string              `bson:"channel_id"`
	Status        string              `bson:"status"`
	FromAddress   string              `bson:"from_address"`
	FromHashcode  string              `bson:"from_hashcode"`
	FromPayload   interface{}         `bson:"from_payload"`
	FromSignature string              `bson:"from_signature"`
	ToAddress     string              `bson:"to_address"`
	ToHashcode    string              `bson:"to_hashcode"`
	ToPayload     string              `bson:"to_payload"`
	ToSignature   string              `bson:"to_signature"`
	CreatedAt     primitive.Timestamp `bson:"created_at"`
	UpdatedAt     primitive.Timestamp `bson:"updated_at"`
}
