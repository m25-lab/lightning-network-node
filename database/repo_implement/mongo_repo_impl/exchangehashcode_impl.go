package mongo_repo_impl

import (
	"context"
	"github.com/m25-lab/lightning-network-node/database/models"
	repo "github.com/m25-lab/lightning-network-node/database/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ExchangeHashcodeRepoImplMongo struct {
	Db *mongo.Database
}

func (mongo *ExchangeHashcodeRepoImplMongo) InsertSecret(ctx context.Context, input *models.ExchangeHashcodeData) error {
	if _, err := mongo.Db.Collection(ExchangeHashcode).InsertOne(ctx, input); err != nil {
		return err
	}
	return nil
}

func (mongo *ExchangeHashcodeRepoImplMongo) FindByOwnHash(ctx context.Context, hash string) (*models.ExchangeHashcodeData, error) {
	result := models.ExchangeHashcodeData{}

	response := mongo.Db.Collection(ExchangeHashcode).FindOne(ctx, bson.M{"my_hashcode": hash})
	if err := response.Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (mongo *ExchangeHashcodeRepoImplMongo) FindByPartnerHash(ctx context.Context, hash string) (*models.ExchangeHashcodeData, error) {
	result := models.ExchangeHashcodeData{}

	response := mongo.Db.Collection(ExchangeHashcode).FindOne(ctx, bson.M{"partner_hashcode": hash})
	if err := response.Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (mongo *ExchangeHashcodeRepoImplMongo) UpdateSecret(ctx context.Context, input *models.ExchangeHashcodeData) error {
	filter := bson.D{{"partner_hashcode", input.PartnerHashcode}}
	update := bson.D{{"$set", bson.D{{"partner_secret", input.PartnerSecret}}}}
	_, err := mongo.Db.Collection(ExchangeHashcode).UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func NewExchangeHashcodeRepo(db *mongo.Database) repo.ExchangeHashcodeRepo {
	return &ExchangeHashcodeRepoImplMongo{
		Db: db,
	}
}
