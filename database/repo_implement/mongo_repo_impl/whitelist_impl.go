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
	existedWhitelist, err := mongo.FindOneByPartnerAddress(ctx, msg.Owner, msg.PartnerAddress)
	if (err == nil) && (existedWhitelist != nil) {
		return errors.New("Whitelist already existed")
	}
	if _, err := mongo.Db.Collection(Whitelist).InsertOne(ctx, msg); err != nil {
		return err
	}
	return nil
}

func (mongo *WhitelistRepoImplMongo) FindOneByPartnerAddress(ctx context.Context, owner string, partnerAddress string) (*models.Whitelist, error) {
	whitelist := models.Whitelist{}
	response := mongo.Db.Collection(Whitelist).FindOne(ctx, bson.M{"owner": owner, "partner_address": partnerAddress})
	if err := response.Decode(&whitelist); err != nil {
		return nil, err
	}

	return &whitelist, nil
}

func (mongo *WhitelistRepoImplMongo) FindMany(ctx context.Context, owner string) ([]models.Whitelist, error) {
	cur, err := mongo.Db.Collection(Whitelist).Find(ctx, bson.M{"owner": owner})
	if err != nil {
		return nil, err
	}

	var whitelists []models.Whitelist
	for cur.Next(ctx) {
		whitelist := models.Whitelist{}
		if err := cur.Decode(&whitelist); err != nil {
			return nil, err
		}

		whitelists = append(whitelists, whitelist)
	}

	return whitelists, nil
}
