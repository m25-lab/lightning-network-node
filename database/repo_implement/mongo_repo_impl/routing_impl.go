package mongo_repo_impl

import (
	"context"

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

func (mongo *RoutingRepoImplMongo) FindRoutingByBroadcastIDAndType(ctx context.Context, broadcastID string, routingType string) ([]*models.Routing, error) {
	var routings []*models.Routing = []*models.Routing{}

	response := mongo.Db.Collection(Routing).FindOne(ctx, bson.M{
		"broadcast_id": broadcastID,
		"routing_type": routingType,
	})
	if err := response.Decode(routings); err != nil {
		return nil, err
	}

	return routings, nil
}
