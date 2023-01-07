package main

import (
	"github.com/m25-lab/lightning-network-node/internal/client"
)

func main() {
	// channelId := client.OpenChannelFromA()
	// client.OpenChannelFromB(channelId)

	partACommitment, addressA, signature := client.CreateCommitmentFromA()
	partBCommitment, addressB, signature := client.CreateCommitmentFromB(partACommitment, addressA, signature)
	client.StoreCommitmentFromA(partBCommitment, addressB, signature)
}
