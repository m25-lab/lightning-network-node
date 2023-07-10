package repository

import (
	"context"

	"github.com/m25-lab/lightning-network-node/database/models"
)

type RoutingRepo interface {
	InsertOne(context.Context, *models.Routing) error
	FindRouting(context.Context, models.Routing) ([]*models.Routing, error)
	FindByDestAndBroadcastId(ctx context.Context, owner string, destAddr string, brId string) (*models.Routing, error)
}
