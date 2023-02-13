package repository

import (
	"context"

	"github.com/m25-lab/lightning-network-node/database/models"
)

type AddressRepo interface {
	InsertOne(context.Context, *models.Address) error
	FindByClientId(context.Context, string) (*models.Address, error)
	FindByAddress(context.Context, string) (*models.Address, error)
	DeleteByClientId(context.Context, string) error
}
