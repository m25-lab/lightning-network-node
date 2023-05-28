package driver

import (
	"context"
	"fmt"
	"time"

	"github.com/m25-lab/lightning-network-node/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Client               *mongo.Client
	ChannelCollection    *mongo.Collection
	CommitmentCollection *mongo.Collection
	MessageCollection    *mongo.Collection
	RoutingCollection    *mongo.Collection
}

func Connect(ctx context.Context, configs *config.Database) (*MongoDB, error) {
	connectionString := fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority",
		configs.User, configs.Password, configs.Host)

	ctx, cancel := context.WithTimeout(ctx, time.Duration(configs.Timeout))
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		return nil, err
	}

	database := client.Database("testing")

	// Indexes
	messageIndexModel := mongo.IndexModel{
		Keys: bson.D{
			{Key: "users", Value: 1},
			{Key: "action", Value: 1},
		},
	}
	database.Collection("messages").Indexes().CreateOne(ctx, messageIndexModel)

	routingIndexModel := mongo.IndexModel{
		Keys: bson.D{
			{Key: "routing_type", Value: 1},
			{Key: "broadcast_id", Value: 1},
		},
	}
	database.Collection("routing").Indexes().CreateOne(ctx, routingIndexModel)

	return &MongoDB{
		client,
		database.Collection("channels"),
		database.Collection("commitments"),
		database.Collection("messages"),
		database.Collection("routing"),
	}, nil
}
