package node

import (
	"context"
	"net"

	"github.com/m25-lab/lightning-network-node/config"
	"github.com/m25-lab/lightning-network-node/node/core"
	"github.com/m25-lab/lightning-network-node/node/database/mongodb"
	"github.com/m25-lab/lightning-network-node/node/p2p/peer"
)

type LightningNode struct {
	Config          *config.Config
	Database        *mongodb.MongoDB
	ListPeer        peer.ListPeer
	ListOpenChannel core.ListChannel
}

func New(config *config.Config) (*LightningNode, error) {
	database, err := mongodb.Connect(&config.Database)
	if err != nil {
		return nil, err
	}

	node := &LightningNode{
		Config:          config,
		Database:        database,
		ListPeer:        peer.ListPeer{},
		ListOpenChannel: core.ListChannel{},
	}

	return node, nil
}

func (node *LightningNode) AddNewPeer(addr net.Addr) error {
	tcpAddr, err := net.ResolveTCPAddr(addr.Network(), addr.String())
	if err != nil {
		return err
	}

	peer := peer.Peer{
		Addr: *tcpAddr,
	}
	node.ListPeer = append(node.ListPeer, peer)

	return nil
}

func (node *LightningNode) CleanUp() {
	// Disconnect to database
	if err := node.Database.Client.Disconnect(context.Background()); err != nil {
		panic(err)
	}
}
