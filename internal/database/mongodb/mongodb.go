package mongodb

import (
	"context"
	"fmt"
	"time"

	"github.com/m25-lab/lightning-network-node/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Client *mongo.Client
}

func Connect(configs *config.DatabaseConfig) (*MongoDB, error) {
	connectionString := fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority", configs.User, configs.Password, configs.Host)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(configs.Timeout))
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		return nil, err
	}

	return &MongoDB{client}, nil
}
