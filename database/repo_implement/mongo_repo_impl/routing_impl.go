package mongo_repo_impl

import (
	"context"
	"fmt"

	"github.com/m25-lab/lightning-network-node/database/models"
	"github.com/m25-lab/lightning-network-node/database/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type RoutingRepoImplMongo struct {
	Db *mongo.Database
}

func NewRoutingRepo(db *mongo.Database) repository.RoutingRepo {
	return &RoutingRepoImplMongo{
		Db: db,
	}
}

func (mongo *RoutingRepoImplMongo) InsertOne(ctx context.Context, routing *models.Routing) error {
	if _, err := mongo.Db.Collection(Routing).InsertOne(ctx, routing); err != nil {
		return err
	}
	return nil
}

func (mongo *RoutingRepoImplMongo) FindRouting(ctx context.Context, input models.Routing) ([]*models.Routing, error) {
	filter := bson.M{}
	if input.Owner == "" {
		return nil, fmt.Errorf("Owner required")
	} else {
		filter["owner"] = input.Owner
	}

	if input.Type == "" {
		return nil, fmt.Errorf("Routing type required")
	} else {
		filter["type"] = input.Type
	}

	if input.BroadcastID != "" {
		filter["broadcast_id"] = input.BroadcastID
	}

	if input.DestinationAddress != "" {
		filter["destination_address"] = input.DestinationAddress
	}

	if input.NextHop != "" {
		filter["next_hop"] = input.NextHop
	}

	// fmt.Println("filter ", filter)
	cur, err := mongo.Db.Collection(Routing).Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	routings := []*models.Routing{}
	for cur.Next(ctx) {
		routing := models.Routing{}
		if err := cur.Decode(&routing); err != nil {
			return nil, err
		}

		routings = append(routings, &routing)
	}

	if len(routings) == 0 {
		return nil, fmt.Errorf("NOT_FOUND")
	}

	return routings, nil
}
