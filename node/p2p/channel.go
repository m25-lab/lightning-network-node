package p2p

import (
	"math/big"
)

type OpenChannel struct {
	linkedNode Node
	capacity   big.Int
}
