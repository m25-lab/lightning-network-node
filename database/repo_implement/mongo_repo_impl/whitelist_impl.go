package mongo_repo_impl

import (
	"context"
	"errors"

	"github.com/m25-lab/lightning-network-node/database/models"
	"github.com/m25-lab/lightning-network-node/database/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type WhitelistRepoImplMongo struct {
	Db *mongo.Database
}

func NewWhitelistRepo(db *mongo.Database) repository.WhitelistRepo {
	return &WhitelistRepoImplMongo{
		Db: db,
	}
}

func (mongo *WhitelistRepoImplMongo) InsertOne(ctx context.Context, msg *models.Whitelist) error {
	existdWhitelist, err := mongo.FindOneByMultiAddress(ctx, msg.MultiAddress)
	if err == nil && existdWhitelist.Users[0] == msg.Users[0] {
		return errors.New("address already in whitelist")
	}

	if _, err := mongo.Db.Collection(Whitelist).InsertOne(ctx, msg); err != nil {
		return err
	}
	return nil
}

func (mongo *WhitelistRepoImplMongo) FindByMultiAddress(ctx context.Context, address string) (*models.Whitelist, error) {
	whitelist := models.Whitelist{}

	response := mongo.Db.Collection(Whitelist).FindOne(ctx, bson.M{"multi_pubkey": address})
	if err := response.Decode(&whitelist); err != nil {
		return nil, err
	}

	return &whitelist, nil
}

func (mongo *WhitelistRepoImplMongo) FindByAddresses(ctx context.Context, addresses []string) (*models.Whitelist, error) {
	whitelist := models.Whitelist{}

	response := mongo.Db.Collection(Whitelist).FindOne(ctx, bson.M{"users": addresses})
	if err := response.Decode(&whitelist); err != nil {
		return nil, err
	}

	return &whitelist, nil
}

func (mongo *WhitelistRepoImplMongo) FindOneByMultiAddress(ctx context.Context, multiAddress string) (*models.Whitelist, error) {
	whitelist := models.Whitelist{}

	response := mongo.Db.Collection(Whitelist).FindOne(ctx, bson.M{"multi_address": multiAddress})
	if err := response.Decode(&whitelist); err != nil {
		return nil, err
	}

	return &whitelist, nil
}

func (mongo *WhitelistRepoImplMongo) FindManyByAddress(ctx context.Context, address string) ([]*models.Whitelist, error) {
	whitelist := []*models.Whitelist{}

	cursor, err := mongo.Db.Collection(Whitelist).Find(ctx, bson.M{"users.0": address})
	if err != nil {
		return nil, err
	}

	if err := cursor.All(ctx, &whitelist); err != nil {
		return nil, err
	}

	return whitelist, nil
}
