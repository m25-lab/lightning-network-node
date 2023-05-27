package repository

import (
	"context"
	"github.com/m25-lab/lightning-network-node/database/models"
)

type ExchangeHashcodeRepo interface {
	InsertSecret(ctx context.Context, input *models.ExchangeHashcodeData) error
	FindByOwnHash(ctx context.Context, hash string) (*models.ExchangeHashcodeData, error)
	FindByPartnerHash(ctx context.Context, hash string) (*models.ExchangeHashcodeData, error)
}
