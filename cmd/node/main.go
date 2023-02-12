package main

import (
	"fmt"
	"sync"

	"github.com/m25-lab/lightning-network-node/client"
	"github.com/m25-lab/lightning-network-node/config"
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
	config, err := config.LoadConfig()
	checkErr(err)

	node, err := node.New(&config)
	checkErr(err)
	defer node.CleanUp()

	fmt.Printf("Running RPC Server...\n")
	rpcServer, err := rpc.New(node)
	checkErr(err)

	go func() {
		rpcServer.RunGateway()
		wg.Done()
	}()

	fmt.Printf("Running Telegram Bot...\n")
	client, err := client.New(node, &config)
	checkErr(err)

	go func() {
		client.RunTelegramBot()
		wg.Done()
	}()

	fmt.Println("Connected to MongoDB!")
	wg.Wait()
}
