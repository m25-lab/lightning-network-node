package main

import (
	"fmt"
	"github.com/m25-lab/lightning-network-node/config"
	"github.com/m25-lab/lightning-network-node/node"
	"github.com/m25-lab/lightning-network-node/node/rpc"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	config, err := config.LoadConfig()
	checkErr(err)

	node, err := node.New(&config)
	checkErr(err)
	defer node.CleanUp()

	rpcServer, err := rpc.New(node)
	checkErr(err)

	err = rpcServer.RunGateway()
	checkErr(err)

	fmt.Println("Connected to MongoDB!")
}
