package node

import (
	"context"
	"net"

	"github.com/m25-lab/lightning-network-node/config"
	"github.com/m25-lab/lightning-network-node/database/driver"
	"github.com/m25-lab/lightning-network-node/database/repo_implement/mongo_repo_impl"
	"github.com/m25-lab/lightning-network-node/database/repository"
	"github.com/m25-lab/lightning-network-node/node/core"
	"github.com/m25-lab/lightning-network-node/node/p2p/peer"
)

type LightningNode struct {
	Config          *config.Config
	Database        *driver.MongoDB
	Repository      *Repository
	ListPeer        peer.ListPeer
	ListOpenChannel core.ListChannel
}

type Repository struct {
	Commitment repository.CommitmentRepo
}

func New(config *config.Config) (*LightningNode, error) {
	database, err := driver.Connect(&config.Database)
	if err != nil {
		return nil, err
	}

	repository := &Repository{}
	repository.Commitment = mongo_repo_impl.NewCommitmentRepo(database.Client.Database(config.Database.Dbname))

	node := &LightningNode{
		config,
		database,
		repository,
		peer.ListPeer{},
		core.ListChannel{},
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
