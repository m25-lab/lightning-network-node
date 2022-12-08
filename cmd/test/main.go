package main

import (
	"github.com/m25-lab/lightning-network-node/internal/client"
)

func main() {
	channelId := client.OpenChannelFromA()
	client.OpenChannelFromB(channelId)
}
