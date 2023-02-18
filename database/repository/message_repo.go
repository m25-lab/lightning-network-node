package repository

import (
	"context"

	"github.com/m25-lab/lightning-network-node/database/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MessageRepo interface {
	InsertOne(context.Context, *models.Message) error
	FindOneById(context.Context, string, string) (*models.Message, error)
	FindOneByOriginalID(context.Context, string, string) (*models.Message, error)
	FindMany(context.Context, string, string) ([]models.Message, error)
	UpdateTelegramChatId(context.Context, primitive.ObjectID, int) error
}
