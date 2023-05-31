package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Routing struct {
	ID                 primitive.ObjectID `bson:"_id, omitempty"`
	BroadcastID        string             `bson:"broadcast_id"`
	DestinationAddress string             `bson:"destination_address"`
	NextHop            string             `bson:"next_hop"`
	Owner              string             `bson:"owner"`
}
