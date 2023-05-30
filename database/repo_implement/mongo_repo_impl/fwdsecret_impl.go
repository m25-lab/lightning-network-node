package mongo_repo_impl

import (
	"context"
	"github.com/m25-lab/lightning-network-node/database/models"
	repo "github.com/m25-lab/lightning-network-node/database/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type FwdSecrettRepoImplMongo struct {
	Db *mongo.Database
}

func (mongo *FwdSecrettRepoImplMongo) InsertSecret(ctx context.Context, input *models.FwdSecret) error {
	if _, err := mongo.Db.Collection(FwdSecret).InsertOne(ctx, input); err != nil {
		return err
	}
	return nil
}

func (mongo *FwdSecrettRepoImplMongo) FindByDestHash(ctx context.Context, hash string) (*models.FwdSecret, error) {
	result := models.FwdSecret{}

	response := mongo.Db.Collection(FwdSecret).FindOne(ctx, bson.M{"hashcode_dest": hash})
	if err := response.Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (mongo *FwdSecrettRepoImplMongo) DeleteByDestHash(ctx context.Context, hash string) error {
	_, err := mongo.Db.Collection(FwdSecret).DeleteMany(ctx, bson.M{"hashcode_dest": hash})
	if err != nil {
		return err
	}
	return nil
}

func NewFwdSecretRepo(db *mongo.Database) repo.FwdSecretRepo {
	return &FwdSecrettRepoImplMongo{
		Db: db,
	}
}
