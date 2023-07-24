package main

import (
	"flag"
	"fmt"
	"log"
	"sync"

	"github.com/m25-lab/lightning-network-node/client"
	"github.com/m25-lab/lightning-network-node/config"
	"github.com/m25-lab/lightning-network-node/job"
	"github.com/m25-lab/lightning-network-node/node"
	"github.com/m25-lab/lightning-network-node/rpc"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(2)

	fmt.Printf("Starting Lightning Network Node...\n")
	configType := flag.String("config", "", "Configuration option")
	flag.Parse()

	config, err := config.LoadConfig(configType)
	checkErr(err)

	node, err := node.New(config)
	checkErr(err)
	defer node.CleanUp()

	fmt.Printf("Running RPC Server...\n")
	rpcServer, err := rpc.New(node)
	checkErr(err)

	go func() {
		rpcServer.RunGateway(config.LNode.External)
		wg.Done()
	}()

	fmt.Printf("Running Telegram Bot...\n")
	client, err := client.New(node)
	checkErr(err)

	go func() {
		client.RunTelegramBot()
		wg.Done()
	}()

	log.Println("Running jobs...")
	manager, err := job.New(node, client)
	checkErr(err)
	go func() {
		manager.Run()
		wg.Done()
	}()

	log.Println("Connected to MongoDB!")
	wg.Wait()
}
