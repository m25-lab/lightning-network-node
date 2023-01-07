package repository

import (
	"context"

	"github.com/m25-lab/lightning-network-node/database/models"
)

type ChannelRepo interface {
	InsertOpenChannelRequest(context.Context, *models.OpenChannelRequest) error
	FindChannelById(context.Context, string) (*models.OpenChannelRequest, error)
}
