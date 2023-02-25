package mongo_repo_impl

import (
	"context"

	"github.com/m25-lab/lightning-network-node/database/models"
	"github.com/m25-lab/lightning-network-node/database/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (mongo *MessageRepoImplMongo) FindOneById(ctx context.Context, owner string, id string) (*models.Message, error) {
	message := models.Message{}

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	response := mongo.Db.Collection(Message).FindOne(ctx, bson.M{"_id": oid, "owner": owner})
	if err := response.Decode(&message); err != nil {
		return nil, err
	}

	return &message, nil
}

func (mongo *MessageRepoImplMongo) FindOneByOriginalID(ctx context.Context, owner string, id string) (*models.Message, error) {
	message := models.Message{}

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	response := mongo.Db.Collection(Message).FindOne(ctx, bson.M{"original_id": oid, "owner": owner})
	if err := response.Decode(&message); err != nil {
		return nil, err
	}

	return &message, nil
}

func (mongo *MessageRepoImplMongo) FindMany(ctx context.Context, owner string, action string) ([]models.Message, error) {
	cur, err := mongo.Db.Collection(Message).Find(ctx, bson.M{"action": action, "owner": owner})
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

func (mongo *MessageRepoImplMongo) UpdateTelegramChatId(ctx context.Context, id primitive.ObjectID, telegramChatId int) error {
	_, err := mongo.Db.Collection(Message).UpdateByID(ctx, id, bson.M{"$set": bson.M{"telegram_chat_id": telegramChatId}})
	if err != nil {
		return err
	}

	return nil
}

func (mongo *MessageRepoImplMongo) Update(ctx context.Context, id primitive.ObjectID, message *models.Message) error {
	updatePayload := bson.M{
		"$set": bson.M{
			"action":           message.Action,
			"telegram_chat_id": message.TelegramChatId,
			"is_replied":       message.IsReplied,
		},
	}
	_, err := mongo.Db.Collection(Message).UpdateByID(ctx, id, updatePayload)
	if err != nil {
		return err
	}

	return nil
}

func (mongo *MessageRepoImplMongo) FindOneByChannelID(ctx context.Context, owner string, ChannelID string) (*models.Message, error) {
	message := models.Message{}

	//get last message
	response := mongo.Db.Collection(Message).FindOne(
		ctx,
		bson.M{"channel_id": ChannelID, "owner": owner},
		options.FindOne().SetSort(bson.M{"$natural": -1}),
	)
	if err := response.Decode(&message); err != nil {
		return nil, err
	}

	return &message, nil
}
