package main

import (
	"fmt"
	"github.com/m25-lab/lightning-network-node/internal/client"

	"github.com/m25-lab/lightning-network-node/config"
	"github.com/m25-lab/lightning-network-node/node"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	//test
	client.OpenChannel()

	config, err := config.LoadConfig()
	checkErr(err)

	node, err := node.New(&config)
	checkErr(err)
	defer node.CleanUp()

	err = node.Server.Run()
	checkErr(err)

	fmt.Println("Connected to MongoDB!")
}
