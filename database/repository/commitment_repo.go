package repository

import (
	"context"
	"github.com/m25-lab/lightning-network-node/database/models"
)

type CommitmentRepo interface {
	Insert(context.Context, *models.Commitment) error
	FindCommitmentById(context.Context, string) (*models.Commitment, error)
}
