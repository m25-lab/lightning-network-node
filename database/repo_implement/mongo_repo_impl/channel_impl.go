package mongo_repo_impl

import (
	"context"

	"github.com/m25-lab/lightning-network-node/database/models"
	"github.com/m25-lab/lightning-network-node/database/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ChannelRepoImplMongo struct {
	Db *mongo.Database
}

func NewChannelRepo(db *mongo.Database) repository.ChannelRepo {
	return &ChannelRepoImplMongo{
		Db: db,
	}
}

func (mongo *ChannelRepoImplMongo) InsertOpenChannelRequest(ctx context.Context, channel *models.OpenChannelRequest) error {
	if _, err := mongo.Db.Collection(Channel).InsertOne(ctx, channel); err != nil {
		return err
	}

	return nil
}

func (mongo *ChannelRepoImplMongo) FindChannelById(ctx context.Context, id string) (*models.OpenChannelRequest, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var channel *models.OpenChannelRequest

	response := mongo.Db.Collection(Channel).FindOne(ctx, bson.M{"_id": objectId})
	if err := response.Decode(channel); err != nil {
		return nil, err
	}

	return channel, nil
}
