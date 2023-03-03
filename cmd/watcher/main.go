package main

import (
	"context"

	"github.com/m25-lab/lightning-network-node/config"
	"github.com/m25-lab/lightning-network-node/watcher"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	config, err := config.LoadConfig()
	if err != nil {
		log.Panic(err)
	}

	watcher, err := watcher.NewWatcher(config)
	if err != nil {
		log.Panic(err)
	}

	if err := watcher.Start(context.Background()); err != nil {
		log.Panic(err)
	}
}
