package repository

import (
	"context"

	"github.com/m25-lab/lightning-network-node/database/models"
)

type MessageRepo interface {
	InsertOne(context.Context, *models.Message) error
	FindOneById(context.Context, string) (*models.Message, error)
	FindMany(context.Context, string, string) ([]models.Message, error)
}
