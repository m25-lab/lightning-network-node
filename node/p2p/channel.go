package p2p

import (
	"math/big"
)

// A channel will be connect two node with one of them is our node.
type OpenChannel struct {
	linkedNode Node
	capacity   big.Int
}
