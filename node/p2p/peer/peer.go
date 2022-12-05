package peer

import (
	"fmt"
	"net"
	"strconv"

	"github.com/m25-lab/lightning-network-node/node/p2p"
	"github.com/m25-lab/lightning-network-node/node/rpc/pb"
)

type Peer struct {
	Addr net.TCPAddr
}

func (peer *Peer) Proto() (*pb.Peer, error) {
	strAddr, err := AddrToString(&peer.Addr)
	if err != nil {
		return nil, err
	}

	return &pb.Peer{
		Addr: strAddr,
	}, nil
}

type ListPeer []Peer

func (listPeer *ListPeer) Proto() ([]*pb.Peer, error) {
	listPeerProto := make([]*pb.Peer, len(*listPeer))

	for index, peer := range *listPeer {
		peerProto, err := peer.Proto()
		if err != nil {
			return nil, err
		}

		listPeerProto[index] = peerProto
	}

	return listPeerProto, nil
}

/*
layout: |version address:port|
*/
func AddrToString(addr *net.TCPAddr) (string, error) {
	ipVer, err := p2p.GetIPVersion(addr)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s %s", strconv.Itoa(int(ipVer)), addr.String()), nil
}
