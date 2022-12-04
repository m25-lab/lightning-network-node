package peer

import (
	"net"
)

type Peer struct {
	Addr net.TCPAddr
}

type ListPeer []Peer
