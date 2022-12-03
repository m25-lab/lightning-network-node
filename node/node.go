package node

import (
	"context"

	"github.com/m25-lab/lightning-network-node/config"
	"github.com/m25-lab/lightning-network-node/node/core"
	"github.com/m25-lab/lightning-network-node/node/database/mongodb"
	"github.com/m25-lab/lightning-network-node/node/p2p"
	"github.com/m25-lab/lightning-network-node/node/rpc"
)

type LightningNode struct {
	Config          *config.Config
	Database        *mongodb.MongoDB
	Server          *rpc.NodeServer
	ListPeer        p2p.ListPeer
	ListOpenChannel core.ListOpenChannel
}

func New(config *config.Config) (*LightningNode, error) {
	database, err := mongodb.Connect(&config.Database)
	if err != nil {
		return nil, err
	}

	server, err := rpc.New()
	if err != nil {
		return nil, err
	}

	node := &LightningNode{
		Config:          config,
		Database:        database,
		Server:          server,
		ListPeer:        p2p.ListPeer{},
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
