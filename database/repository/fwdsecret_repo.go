package repository

import (
	"context"
	"github.com/m25-lab/lightning-network-node/database/models"
)

type FwdSecretRepo interface {
	InsertSecret(ctx context.Context, input *models.FwdSecret) error
	FindByDestHash(ctx context.Context, owner string, hash string) (*models.FwdSecret, error)
	DeleteByDestHash(ctx context.Context, hash string) error
}
