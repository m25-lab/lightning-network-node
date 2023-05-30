package repository

import (
	"context"

	"github.com/m25-lab/lightning-network-node/database/models"
)

type RoutingRepo interface {
	InsertOne(context.Context, *models.Routing) error
	FindRouting(context.Context, models.Routing) ([]*models.Routing, error)
	FindByDestAndBroadcastId(context.Context, string, string, string) (*models.Routing, error)
}
