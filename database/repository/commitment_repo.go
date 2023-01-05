package repository

import (
	"context"

	"github.com/m25-lab/lightning-network-node/database/models"
)

type CommitmentRepo interface {
	FindCommitmentById(context.Context, string) (*models.Commitment, error)
	Insert(context.Context, *models.Commitment) error
}
