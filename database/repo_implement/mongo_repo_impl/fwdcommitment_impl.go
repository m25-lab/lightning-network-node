package mongo_repo_impl

import (
	"context"
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

func (mongo *FwdCommitmentRepoImplMongo) InsertSenderCommit(ctx context.Context, sdC *models.SenderCommitment) error {
	if _, err := mongo.Db.Collection(SenderCommitment).InsertOne(ctx, sdC); err != nil {
		return err
	}
	return nil
}

func (mongo *FwdCommitmentRepoImplMongo) InsertReceiverCommit(ctx context.Context, rcC *models.ReceiverCommitment) error {
	if _, err := mongo.Db.Collection(ReceiverCommitment).InsertOne(ctx, rcC); err != nil {
		return err
	}
	return nil
}

func (mongo *FwdCommitmentRepoImplMongo) FindReceiverCommitByDestHash(ctx context.Context, hash string) (*models.ReceiverCommitment, error) {
	rcC := models.ReceiverCommitment{}

	response := mongo.Db.Collection(ReceiverCommitment).FindOne(ctx, bson.M{"hashcode_dest": hash})
	if err := response.Decode(&rcC); err != nil {
		return nil, err
	}

	return &rcC, nil
}

func (mongo *FwdCommitmentRepoImplMongo) FindSenderCommitByDestHash(ctx context.Context, hash string) (*models.SenderCommitment, error) {
	sdC := models.SenderCommitment{}

	response := mongo.Db.Collection(SenderCommitment).FindOne(ctx, bson.M{"hashcode_dest": hash})
	if err := response.Decode(&sdC); err != nil {
		return nil, err
	}

	return &sdC, nil
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
