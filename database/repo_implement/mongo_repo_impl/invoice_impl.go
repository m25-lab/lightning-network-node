package mongo_repo_impl

import (
	"context"

	"github.com/m25-lab/lightning-network-node/database/models"
	repo "github.com/m25-lab/lightning-network-node/database/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type InvoiceRepoImplMongo struct {
	Db *mongo.Database
}

func (mongo *InvoiceRepoImplMongo) InsertInvoice(ctx context.Context, input *models.InvoiceData) error {
	if _, err := mongo.Db.Collection(Invoice).InsertOne(ctx, input); err != nil {
		return err
	}
	return nil
}

func (mongo *InvoiceRepoImplMongo) FindByHash(ctx context.Context, owner string, hash string) (*models.InvoiceData, error) {
	result := models.InvoiceData{}

	response := mongo.Db.Collection(Invoice).FindOne(ctx, bson.M{"hash": hash, "to": owner})
	if err := response.Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (mongo *InvoiceRepoImplMongo) FindByHashFrom(ctx context.Context, owner string, hash string) (*models.InvoiceData, error) {
	result := models.InvoiceData{}

	response := mongo.Db.Collection(Invoice).FindOne(ctx, bson.M{"hash": hash, "from": owner})
	if err := response.Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func NewInvoiceRepo(db *mongo.Database) repo.InvoiceRepo {
	return &InvoiceRepoImplMongo{
		Db: db,
	}
}
