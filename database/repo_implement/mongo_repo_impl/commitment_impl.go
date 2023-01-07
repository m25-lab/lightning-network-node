package mongo_repo_impl

import (
	"context"

	"github.com/m25-lab/lightning-network-node/database/models"
	repo "github.com/m25-lab/lightning-network-node/database/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CommitmentRepoImplMongo struct {
	Db *mongo.Database
}

func NewCommitmentRepo(db *mongo.Database) repo.CommitmentRepo {
	return &CommitmentRepoImplMongo{
		Db: db,
	}
}

func (mongo *CommitmentRepoImplMongo) Insert(ctx context.Context, commitment *models.Commitment) error {
	bbytes, _ := bson.Marshal(commitment)

	_, err := mongo.Db.Collection(Commitment).InsertOne(ctx, bbytes)
	if err != nil {
		return err
	}

	return nil
}

func (mongo *CommitmentRepoImplMongo) FindCommitmentById(ctx context.Context, id string) (*models.Commitment, error) {
	commitment := models.Commitment{}

	response := mongo.Db.Collection(Commitment).FindOne(ctx, bson.M{"id": id})
	if err := response.Decode(&commitment); err != nil {
		return nil, err
	}

	return &commitment, nil
}
