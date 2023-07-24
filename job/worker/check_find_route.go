package worker

import (
	"context"
	"fmt"
	"log"

	"github.com/m25-lab/lightning-network-node/client"
	"github.com/m25-lab/lightning-network-node/database/models"
	"github.com/m25-lab/lightning-network-node/node"
	"go.mongodb.org/mongo-driver/mongo"
)

type CheckFindRoute struct {
	Repository *node.Repository
	Client     *client.Client
}

func NewCheckFindRoute(repo *node.Repository, client *client.Client) (*CheckFindRoute, error) {
	return &CheckFindRoute{
		Repository: repo,
	}, nil
}

func (worker CheckFindRoute) Handler() {
	ctx := context.Background()
	jobData, err := worker.Repository.JobQueue.Consume(ctx, models.CheckFindRouteJobName)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return
		}
		log.Println(models.CheckFindRouteJobName + " consume failed")
		return
	}

	data := jobData.Data.(*models.CheckFindRoute)
	_, err = worker.Repository.Routing.FindByDestAndBroadcastId(context.Background(), data.From, data.To, data.Hash)
	if err != nil {
		address, err := worker.Repository.Address.FindByAddress(context.Background(), data.From)
		if err != nil {
			return
		}
		worker.Client.SendTele(address.ClientId, fmt.Sprintf("System not found route to %s. Please wait a minute and make another ln transfer multi!", data.From))
	}

	err = worker.Repository.JobQueue.MarkConsumed(ctx, jobData.ID.Hex(), models.CheckFindRouteJobName)
	if err != nil {
		return
	}
}
