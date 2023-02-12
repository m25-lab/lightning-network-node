package repository

import (
	"context"

	"github.com/m25-lab/lightning-network-node/database/models"
)

type WhitelistRepo interface {
	InsertOne(context.Context, *models.Whitelist) error
	FindByAddresses(context.Context, []string) (*models.Whitelist, error)
	FindByMultiAddress(context.Context, string) (*models.Whitelist, error)
}
