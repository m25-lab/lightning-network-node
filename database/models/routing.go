package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Routing struct {
	ID                 primitive.ObjectID `json:"_id,omitempty" bson:"_id, omitempty"`
	BroadcastID        string             `json:"broadcast_id" bson:"broadcast_id"`
	DestinationAddress string             `json:"destination_address" bson:"destination_address"`
	NextHop            string             `json:"next_hop" bson:"next_hop"`
	HopCounter         int64              `json:"hop_counter" bson:"hop_counter"`
	SequenceNumber     int64              `json:"sequence_number" bson:"sequence_number"`
	Owner              string             `json:"owner" bson:"owner"`
}

type RREQData struct {
	Amount         int64 `json:"amount"`
	HopCounter     int64 `json:"hopCounter"`
	SequenceNumber int64 `json:"sequenceNumber"`
	RemainReward   int64 `json:"remainReward"`
}

type RREPData struct {
	HopCounter     int64 `json:"hopCounter"`
	SequenceNumber int64 `json:"sequenceNumber"`
}
