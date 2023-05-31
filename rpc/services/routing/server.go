package routing

import (
	"github.com/m25-lab/lightning-network-node/client"
	"github.com/m25-lab/lightning-network-node/node"
	"github.com/m25-lab/lightning-network-node/rpc/pb"
)

type RoutingServer struct {
	pb.UnimplementedRoutingServiceServer
	Node   *node.LightningNode
	Client *client.Client
}

func New(node *node.LightningNode, client *client.Client) (*RoutingServer, error) {
	return &RoutingServer{Node: node, Client: client}, nil
}
