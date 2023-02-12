package mongo_repo_impl

import (
	"context"

	"github.com/m25-lab/lightning-network-node/database/models"
	"github.com/m25-lab/lightning-network-node/database/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AddressRepoImplMongo struct {
	Db *mongo.Database
}

func NewAddressRepo(db *mongo.Database) repository.AddressRepo {
	return &AddressRepoImplMongo{
		Db: db,
	}
}

func (mongo *AddressRepoImplMongo) InsertOne(ctx context.Context, msg *models.Address) error {

	if _, err := mongo.Db.Collection(Address).InsertOne(ctx, msg); err != nil {
		return err
	}
	return nil
}

func (mongo *AddressRepoImplMongo) FindByAddress(ctx context.Context, _address string) (*models.Address, error) {
	address := models.Address{}

	response := mongo.Db.Collection(Address).FindOne(ctx, bson.M{"address": _address})
	if err := response.Decode(&address); err != nil {
		return nil, err
	}

	return &address, nil
}
