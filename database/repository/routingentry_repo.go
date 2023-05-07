package repository

import (
	"context"
	"github.com/m25-lab/lightning-network-node/database/models"
)

type RoutingEntry interface {
	InsertEntry(ctx context.Context, input *models.RoutingEntry) error
	FindByDestAndHash(ctx context.Context, dest string, hash string) (*models.RoutingEntry, error)
}
