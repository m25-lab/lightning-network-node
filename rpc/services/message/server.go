package message

import (
	"github.com/m25-lab/lightning-network-node/client"
	"github.com/m25-lab/lightning-network-node/node"
	"github.com/m25-lab/lightning-network-node/rpc/pb"
)

type MessageServer struct {
	pb.UnimplementedMessageServiceServer
	Node   *node.LightningNode
	Client *client.Client
}

func New(node *node.LightningNode, client *client.Client) (*MessageServer, error) {
	return &MessageServer{Node: node, Client: client}, nil
}
