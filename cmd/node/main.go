package main

import (
	"context"
	"fmt"
	"time"

	"github.com/m25-lab/lightning-network-node/config"
	"github.com/m25-lab/lightning-network-node/database/models"
	"github.com/m25-lab/lightning-network-node/node"
	"github.com/m25-lab/lightning-network-node/rpc"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	commitment := &models.Commitment{
		ID:            "dafsfadfa",
		ChannelID:     "adfadsf",
		Status:        "1",
		FromAddress:   "a",
		FromHashcode:  "aa",
		FromPayload:   nil,
		FromSignature: "aaa",
		ToAddress:     "b",
		ToHashcode:    "bb",
		ToPayload:     nil,
		ToSignature:   "bbb",
		CreatedAt:     primitive.Timestamp{T: uint32(time.Now().Unix()), I: 0},
		UpdatedAt:     primitive.Timestamp{T: uint32(time.Now().Unix()), I: 0},
	}

	//err = node.Repository.Commitment.Insert(context.Background(), commitment)
	//fmt.Println(err)

	commitment, err = node.Repository.Commitment.FindCommitmentById(context.Background(), "dafsfadfa")
	fmt.Println(commitment)

	err = rpcServer.RunGateway()
	checkErr(err)

	fmt.Println("Connected to MongoDB!")
}
