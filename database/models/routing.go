package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Routing struct {
	ID                 primitive.ObjectID `bson:"_id, omitempty"`
	BroadcastID        string             `bson:"broadcast_id"`
	DestinationAddress string             `bson:"destination_address"`
	NextHop            string             `bson:"next_hop"`
	HopCounter         int64              `bson:"hop_counter"`
	SequenceNumber     int64              `bson:"sequence_number"`
	Owner              string             `bson:"owner"`
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
