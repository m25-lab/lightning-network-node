package main

import (
	"github.com/m25-lab/lightning-network-node/client"
)

func main() {
	partACommitment, addressA, signature := client.CreateCommitmentFromA(50, 50, "secret from A")
	partBCommitment, addressB, signature := client.CreateCommitmentFromB(partACommitment, addressA, signature, "secret from B")
	client.StoreCommitmentFromA(partBCommitment, addressB, signature)

	channelId := client.OpenChannelFromA(50, 50)
	client.OpenChannelFromB(channelId)

	partACommitment, addressA, signature = client.CreateCommitmentFromA(40, 60, "secret from AA")
	partBCommitment, addressB, signature = client.CreateCommitmentFromB(partACommitment, addressA, signature, "secret from BB")
	client.StoreCommitmentFromA(partBCommitment, addressB, signature)

	partACommitment, addressA, signature = client.CreateCommitmentFromA(70, 30, "secret from AA")
	partBCommitment, addressB, signature = client.CreateCommitmentFromB(partACommitment, addressA, signature, "secret from BB")
	client.StoreCommitmentFromA(partBCommitment, addressB, signature)
}
