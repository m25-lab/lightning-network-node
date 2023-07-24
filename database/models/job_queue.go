package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JobQueueData struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id, omitempty"`
	Owner     string             `json:"owner" bson:"owner"`
	ReadyTime time.Time          `json:"readyTime,omitempty" bson:"ready_time,omitempty"`
	Topic     string             `json:"topic,omitempty" bson:"topic,omitempty"`
	Data      interface{}        `json:"data,omitempty" bson:"data,omitempty"`
}

type CheckFindRoute struct {
	From          string     `json:"from" bson:"from"`
	To            string     `json:"to" bson:"to"`
	Hash          string     `json:"hash" bson:"hash"`
	StartFindTime *time.Time `json:"start_find_time" bson:"start_find_time"`
}

var (
	CheckFindRouteJobName = "check_find_route"
)
