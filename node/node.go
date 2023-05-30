package node

import (
	"context"

	"github.com/m25-lab/lightning-network-node/config"
	"github.com/m25-lab/lightning-network-node/database/driver"
	"github.com/m25-lab/lightning-network-node/database/repo_implement/mongo_repo_impl"
	"github.com/m25-lab/lightning-network-node/database/repository"
)

type LightningNode struct {
	Config     *config.Config
	Database   *driver.MongoDB
	Repository *Repository
}

type Repository struct {
	Commitment       repository.CommitmentRepo
	Channel          repository.ChannelRepo
	Message          repository.MessageRepo
	Whitelist        repository.WhitelistRepo
	Address          repository.AddressRepo
	Routing          repository.RoutingRepo
	FwdSecret        repository.FwdSecretRepo
	RoutingEntry     repository.RoutingEntry
	Invoice          repository.InvoiceRepo
	ExchangeHashcode repository.ExchangeHashcodeRepo
	FwdCommitment    repository.FwdCommitmentRepo
}

func New(config *config.Config) (*LightningNode, error) {
	database, err := driver.Connect(context.Background(), &config.Database)
	if err != nil {
		return nil, err
	}

	repository := &Repository{
		Commitment:       mongo_repo_impl.NewCommitmentRepo(database.Client.Database(config.Database.Dbname)),
		Channel:          mongo_repo_impl.NewChannelRepo(database.Client.Database(config.Database.Dbname)),
		Message:          mongo_repo_impl.NewMessageRepo(database.Client.Database(config.Database.Dbname)),
		Whitelist:        mongo_repo_impl.NewWhitelistRepo(database.Client.Database(config.Database.Dbname)),
		Address:          mongo_repo_impl.NewAddressRepo(database.Client.Database(config.Database.Dbname)),
		Routing:          mongo_repo_impl.NewRoutingRepo(database.Client.Database(config.Database.Dbname)),
		ExchangeHashcode: mongo_repo_impl.NewExchangeHashcodeRepo(database.Client.Database(config.Database.Dbname)),
		FwdCommitment:    mongo_repo_impl.NewFwdCommitmentRepo(database.Client.Database(config.Database.Dbname)),
		Invoice:          mongo_repo_impl.NewInvoiceRepo(database.Client.Database(config.Database.Dbname)),
		FwdSecret:        mongo_repo_impl.NewFwdSecretRepo(database.Client.Database(config.Database.Dbname)),
		RoutingEntry:     mongo_repo_impl.NewRoutingEntryRepo(database.Client.Database(config.Database.Dbname)),
	}

	node := &LightningNode{
		config,
		database,
		repository,
	}

	return node, nil
}

func (node *LightningNode) CleanUp() {
	// Disconnect to database
	if err := node.Database.Client.Disconnect(context.Background()); err != nil {
		panic(err)
	}
}
