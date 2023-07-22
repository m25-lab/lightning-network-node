package mongo_repo_impl

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/m25-lab/lightning-network-node/database/models"
	repo "github.com/m25-lab/lightning-network-node/database/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type FwdCommitmentRepoImplMongo struct {
	Db *mongo.Database
}

func (mongo *FwdCommitmentRepoImplMongo) InsertFwdMessage(ctx context.Context, sdC *models.FwdMessage) error {
	if _, err := mongo.Db.Collection(FwdMessage).InsertOne(ctx, sdC); err != nil {
		return err
	}
	return nil
}

func (mongo *FwdCommitmentRepoImplMongo) FindOneById(ctx context.Context, owner string, id string) (*models.FwdMessage, error) {
	message := models.FwdMessage{}

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	response := mongo.Db.Collection(FwdMessage).FindOne(ctx, bson.M{"_id": oid, "to": owner})
	if err := response.Decode(&message); err != nil {
		return nil, err
	}

	return &message, nil
}

func (mongo *FwdCommitmentRepoImplMongo) FindReceiverCommitByDestHash(ctx context.Context, owner string, hash string) (*models.FwdMessage, error) {
	rcC := models.FwdMessage{}

	response := mongo.Db.Collection(FwdMessage).FindOne(ctx, bson.M{"action": models.ReceiverCommit, "hashcode_dest": hash, "to": owner})
	if err := response.Decode(&rcC); err != nil {
		return nil, err
	}

	return &rcC, nil
}

func (mongo *FwdCommitmentRepoImplMongo) FindSenderCommitByDestHash(ctx context.Context, owner string, hash string) (*models.FwdMessage, error) {
	rcC := models.FwdMessage{}

	response := mongo.Db.Collection(FwdMessage).FindOne(ctx, bson.M{"action": models.SenderCommit, "hashcode_dest": hash, "to": owner})
	if err := response.Decode(&rcC); err != nil {
		return nil, err
	}

	return &rcC, nil
}

func (mongo *FwdCommitmentRepoImplMongo) DeleteByDestHash(ctx context.Context, s string) error {
	//TODO implement me
	panic("implement me")
}

func NewFwdCommitmentRepo(db *mongo.Database) repo.FwdCommitmentRepo {
	return &FwdCommitmentRepoImplMongo{
		Db: db,
	}
}
