package node

import (
	"context"

	"github.com/m25-lab/lightning-network-node/config"
	"github.com/m25-lab/lightning-network-node/node/core"
	"github.com/m25-lab/lightning-network-node/node/database/mongodb"
	"github.com/m25-lab/lightning-network-node/node/p2p/peer"
)

type LightningNode struct {
	Config          *config.Config
	Database        *mongodb.MongoDB
	ListPeer        peer.ListPeer
	ListOpenChannel core.ListOpenChannel
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
		ListOpenChannel: core.ListOpenChannel{},
	}

	return node, nil
}

func (node *LightningNode) CleanUp() {
	// Disconnect to database
	if err := node.Database.Client.Disconnect(context.Background()); err != nil {
		panic(err)
	}
}
