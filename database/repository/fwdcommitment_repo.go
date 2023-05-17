package repository

import (
	"context"
	"github.com/m25-lab/lightning-network-node/database/models"
)

type FwdCommitmentRepo interface {
	InsertFwdMessage(ctx context.Context, sdC *models.FwdMessage) error
	InsertSenderCommit(ctx context.Context, sdC *models.SenderCommitment) error
	InsertReceiverCommit(ctx context.Context, rcC *models.ReceiverCommitment) error
	FindReceiverCommitByDestHash(ctx context.Context, hash string) (*models.ReceiverCommitment, error)
	FindSenderCommitByDestHash(ctx context.Context, hash string) (*models.SenderCommitment, error)
	DeleteByDestHash(ctx context.Context, s string) error
}
