package repository

import (
	"context"

	"github.com/m25-lab/lightning-network-node/database/models"
)

type WhitelistRepo interface {
	InsertOne(context.Context, *models.Whitelist) error
	FindOneByPartnerAddress(context.Context, string, string) (*models.Whitelist, error)
	FindMany(context.Context, string) ([]models.Whitelist, error)
}
