package core

import (
	"github.com/m25-lab/lightning-network-node/node/p2p"
)

// A channel will be connect two node with one of them is our node.
type Channel struct {
	address    Address
	peers      [2]p2p.Peer
	amount     [2]uint64
	sequence   uint32
	signatures [2]string
}
