package p2pserver

import (
	"github.com/m25-lab/lightning-network-node/node"
	"github.com/m25-lab/lightning-network-node/node/rpc/pb"
)

type PeerToPeerServer struct {
	pb.UnimplementedPeerToPeerServer
	Node *node.LightningNode
}

func New(node *node.LightningNode) (*PeerToPeerServer, error) {
	return &PeerToPeerServer{Node: node}, nil
}
