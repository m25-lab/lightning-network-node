package message

import (
	"github.com/m25-lab/lightning-network-node/node"
	"github.com/m25-lab/lightning-network-node/rpc/pb"
)

type MessageServer struct {
	pb.UnimplementedMessageServiceServer
	Node *node.LightningNode
}

func New(node *node.LightningNode) (*MessageServer, error) {
	return &MessageServer{Node: node}, nil
}
