package repository

import (
	"context"

	"github.com/m25-lab/lightning-network-node/database/models"
)

type RoutingRepo interface {
	InsertOne(context.Context, *models.Routing) error
	FindRoutingByBroadcastIDAndType(context.Context, string, string) ([]*models.Routing, error)
}
