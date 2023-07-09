package repository

import (
	"context"

	"github.com/m25-lab/lightning-network-node/database/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RoutingRepo interface {
	InsertOne(context.Context, *models.Routing) error
	FindRouting(context.Context, models.Routing) ([]*models.Routing, error)
	FindByDestAndBroadcastId(context.Context, string, string, string) (*models.Routing, error)
	DeletedRoutingByNextHop(context.Context, string, string) error
	UpdateRoute(context.Context, primitive.ObjectID, *models.Routing) error
}
