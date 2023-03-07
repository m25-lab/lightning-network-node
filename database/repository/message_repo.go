package repository

import (
	"context"

	"github.com/m25-lab/lightning-network-node/database/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MessageRepo interface {
	InsertOne(context.Context, *models.Message) error
	FindOneById(context.Context, string, string) (*models.Message, error)
	FindOneByChannelID(context.Context, string, string) (*models.Message, error)
	FindOneByChannelIDWithAction(context.Context, string, string, string) (*models.Message, error)
	FindOneByOriginalID(context.Context, string, string) (*models.Message, error)
	FindMany(context.Context, string, string) ([]models.Message, error)
	UpdateTelegramChatId(context.Context, primitive.ObjectID, int) error
	Update(context.Context, primitive.ObjectID, *models.Message) error
}