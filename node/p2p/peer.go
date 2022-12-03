package p2p

import "net"

type Peer struct {
	ip net.TCPAddr
}

type ListPeer []Peer
