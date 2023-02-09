package main

import (
	"github.com/m25-lab/lightning-network-node/client"
)

func main() {
	partACommitment, addressA, signature := client.CreateCommitmentFromA(50, 50, "secret from A", 10, 1)
	partBCommitment, addressB, signature := client.CreateCommitmentFromB(partACommitment, addressA, signature, "secret from B")
	client.StoreCommitmentFromA(partBCommitment, addressB, signature)

	channelId := client.OpenChannelFromA(50, 50, 1)
	client.OpenChannelFromB(channelId)

	partACommitment, addressA, signature = client.CreateCommitmentFromA(40, 60, "secret from AA", 10, 1)
	partBCommitment, addressB, signature = client.CreateCommitmentFromB(partACommitment, addressA, signature, "secret from BB")
	client.StoreCommitmentFromA(partBCommitment, addressB, signature)

	partACommitment, addressA, signature = client.CreateCommitmentFromA(70, 30, "secret from AA", 10, 1)
	partBCommitment, addressB, signature = client.CreateCommitmentFromB(partACommitment, addressA, signature, "secret from BB")
	client.StoreCommitmentFromA(partBCommitment, addressB, signature)

	partAFundCommitment, addressA, signature := client.FundFromA(10, "secret from AA", 10, 1)
	partBAcceptFundCommitment, addressB, signature := client.AcceptFundFromB(partAFundCommitment, addressA, signature, "secret from BB")
	client.StoreAcceptFundCommitmentFromA(partBAcceptFundCommitment, addressB, signature)

	closeChannel, addressA, signature := client.CloseChannel(10, 10, 1)
	client.AcceptCloseChannel(closeChannel, addressA, signature)

}
