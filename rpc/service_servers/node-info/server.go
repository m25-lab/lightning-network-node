package nodeInfo

import (
	"github.com/m25-lab/lightning-network-node/node"
	"github.com/m25-lab/lightning-network-node/node/rpc/pb"
)

type NodeInfoServer struct {
	pb.UnimplementedNodeServiceServer
	Node *node.LightningNode
}

func New(node *node.LightningNode) (*NodeInfoServer, error) {
	return &NodeInfoServer{Node: node}, nil
}
