package mongo_repo_impl

import (
	"context"
	"fmt"

	"github.com/m25-lab/lightning-network-node/database/models"
	"github.com/m25-lab/lightning-network-node/database/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (mongo *RoutingRepoImplMongo) FindByDestAndBroadcastId(ctx context.Context, owner, destinationAdrr, broadcastId string) (*models.Routing, error) {
	nextHop := models.Routing{}
	response := mongo.Db.Collection(Routing).FindOne(ctx, bson.M{
		"owner":               owner,
		"destination_address": destinationAdrr,
		"broadcast_id":        broadcastId,
	})
	if err := response.Decode(&nextHop); err != nil {
		return nil, err
	}

	return &nextHop, nil
}

func (mongo *RoutingRepoImplMongo) DeletedRoutingByNextHop(ctx context.Context, nextHop string, owner string) error {
	_, err := mongo.Db.Collection(Address).DeleteMany(ctx, bson.M{
		"owner":    owner,
		"next_hop": nextHop,
	})
	if err != nil {
		return err
	}

	return nil
}

func (mongo *RoutingRepoImplMongo) UpdateRoute(ctx context.Context, id primitive.ObjectID, routing *models.Routing) error {
	if _, err := mongo.Db.Collection(Routing).UpdateByID(ctx, id, routing); err != nil {
		return err
	}
	return nil
}
