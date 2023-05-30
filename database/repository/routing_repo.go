package repository

import (
	"context"

	"github.com/m25-lab/lightning-network-node/database/models"
)

type RoutingRepo interface {
	InsertOne(context.Context, *models.Routing) error
	FindRouting(context.Context, models.Routing) ([]*models.Routing, error)
}
