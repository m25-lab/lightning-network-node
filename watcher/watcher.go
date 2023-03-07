package watcher

import (
	"context"

	"github.com/m25-lab/lightning-network-node/config"
	"github.com/tendermint/tendermint/rpc/client/http"
	coretypes "github.com/tendermint/tendermint/rpc/core/types"
)

type Watcher struct {
	client    *http.HTTP
	publisher *Publisher
}

func NewWatcher(config *config.Config) (*Watcher, error) {
	client, err := getClient(config.Corechain.Endpoint)
	if err != nil {
		return nil, err
	}

	publisher, err := NewPublisher(config, client)
	if err != nil {
		return nil, err
	}

	return &Watcher{client, publisher}, nil
}

func (w *Watcher) Start(ctx context.Context) error {
	w.client.Start()
	go w.publisher.Start()

	channel, err := w.subscribeNewBlockEvent(ctx)
	if err != nil {
		return err
	}

	for {
		event := <-channel
		res, err := ParseEventDataNewBlock(event)
		if err != nil {
			return nil
		}

		if err := w.publisher.PushBlock(res.Block); err != nil {
			return err
		}
	}
}

func (w *Watcher) subscribeNewBlockEvent(ctx context.Context) (<-chan coretypes.ResultEvent, error) {
	query := "tm.event = 'NewBlock'"
	// Subscriber will be set to ip of watcher by default
	channel, err := w.client.Subscribe(ctx, "", query, 10)
	if err != nil {
		return nil, err
	}

	return channel, err
}
