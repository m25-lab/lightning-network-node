package repository

import (
	"context"
	"github.com/m25-lab/lightning-network-node/database/models"
)

type FwdCommitmentRepo interface {
	InsertFwdMessage(ctx context.Context, sdC *models.FwdMessage) error
	FindReceiverCommitByDestHash(ctx context.Context, owner string, hash string) (*models.FwdMessage, error)
	FindSenderCommitByDestHash(ctx context.Context, owner string, hash string) (*models.FwdMessage, error)
	DeleteByDestHash(ctx context.Context, s string) error
}
