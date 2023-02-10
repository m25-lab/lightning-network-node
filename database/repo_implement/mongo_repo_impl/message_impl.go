package mongo_repo_impl

import (
	"context"

	"github.com/m25-lab/lightning-network-node/database/models"
	"github.com/m25-lab/lightning-network-node/database/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MessageRepoImplMongo struct {
	Db *mongo.Database
}

func NewMessageRepo(db *mongo.Database) repository.MessageRepo {
	return &MessageRepoImplMongo{
		Db: db,
	}
}

func (mongo *MessageRepoImplMongo) InsertOne(ctx context.Context, msg *models.Message) error {
	if _, err := mongo.Db.Collection(Message).InsertOne(ctx, msg); err != nil {
		return err
	}
	return nil
}

func (mongo *MessageRepoImplMongo) FindOneById(ctx context.Context, id string) (*models.Message, error) {
	message := models.Message{}

	response := mongo.Db.Collection(Message).FindOne(ctx, bson.M{"id": id})
	if err := response.Decode(&message); err != nil {
		return nil, err
	}

	return &message, nil
}

func (mongo *MessageRepoImplMongo) FindMany(ctx context.Context, userId string, action string) ([]models.Message, error) {
	cur, err := mongo.Db.Collection(Message).Find(ctx, bson.M{"action": action, "users": userId})
	if err != nil {
		return nil, err
	}

	var messages []models.Message
	for cur.Next(ctx) {
		message := models.Message{}
		if err := cur.Decode(&message); err != nil {
			return nil, err
		}

		messages = append(messages, message)
	}

	return messages, err
}
