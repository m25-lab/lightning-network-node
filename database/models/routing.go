package models

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	RoutingTypeDiscovery string = "RREQ"
	RoutingTypeReply     string = "RREP"
)

type Routing struct {
	ID                 primitive.ObjectID `bson:"_id, omitempty"`
	Type               string             `bson:"type"`
	BroadcastID        string             `bson:"broadcast_id"`
	DestinationAddress string             `bson:"destination_address"`
	NextHop            string             `bson:"next_hop"`
}
