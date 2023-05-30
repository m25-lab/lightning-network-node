package repository

import (
	"context"
	"github.com/m25-lab/lightning-network-node/database/models"
)

type InvoiceRepo interface {
	InsertInvoice(ctx context.Context, input *models.InvoiceData) error
	FindByHash(ctx context.Context, hash string) (*models.InvoiceData, error)
}
