package repository

import (
	"context"

	"github.com/m25-lab/lightning-network-node/database/models"
)

type JobQueueRepo interface {
	Publish(data *models.JobQueueData) error
	Consume(ctx context.Context, topic string) (*models.JobQueueData, error)
	MarkConsumed(ctx context.Context, id string, topic string) error
}
