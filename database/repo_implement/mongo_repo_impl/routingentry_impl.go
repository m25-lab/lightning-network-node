package mongo_repo_impl

import (
	"context"
	"github.com/m25-lab/lightning-network-node/database/models"
	repo "github.com/m25-lab/lightning-network-node/database/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type RoutingEntryRepoImplMongo struct {
	Db *mongo.Database
}

func (mongo RoutingEntryRepoImplMongo) FindDestByHash(ctx context.Context, hash string) (*string, error) {
	result := models.RoutingEntry{}

	response := mongo.Db.Collection(RoutingEntry).FindOne(ctx, bson.M{"hashcode_dest": hash})
	if err := response.Decode(&result); err != nil {
		return nil, err
	}

	return &result.Dest, nil
}

func (mongo RoutingEntryRepoImplMongo) InsertEntry(ctx context.Context, input *models.RoutingEntry) error {
	if _, err := mongo.Db.Collection(RoutingEntry).InsertOne(ctx, input); err != nil {
		return err
	}
	return nil
}

func (mongo RoutingEntryRepoImplMongo) FindByDestAndHash(ctx context.Context, dest string, hash string) (*models.RoutingEntry, error) {
	result := models.RoutingEntry{}

	response := mongo.Db.Collection(RoutingEntry).FindOne(ctx, bson.M{"dest": dest, "hashcode_dest": hash})
	if err := response.Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func NewRoutingEntryRepo(db *mongo.Database) repo.RoutingEntry {
	return &RoutingEntryRepoImplMongo{
		Db: db,
	}
}
