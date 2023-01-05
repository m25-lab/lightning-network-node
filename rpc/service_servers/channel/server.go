package channel

import (
	"github.com/m25-lab/lightning-network-node/node"
	"github.com/m25-lab/lightning-network-node/rpc/pb"
)

type ChannelServer struct {
	pb.UnimplementedChannelServiceServer
	Node *node.LightningNode
}

func New(node *node.LightningNode) (*ChannelServer, error) {
	return &ChannelServer{Node: node}, nil
}
