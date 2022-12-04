package p2pserver

import (
	"context"
	"fmt"
	"net"
	"strconv"

	"github.com/m25-lab/lightning-network-node/node/p2p"
	"github.com/m25-lab/lightning-network-node/node/rpc/pb"
)

type PeerToPeerHandler struct {
}

/*
layout: |version address:port|
*/
func addrToString(addr *net.TCPAddr) (string, error) {
	ipVer, err := p2p.GetIPVersion(addr)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s %s", strconv.Itoa(int(ipVer)), addr.String()), nil
}

func (server *PeerToPeerServer) GetListPeer(ctx context.Context, req *pb.GetListPeerRequest) (*pb.GetListPeerResponse, error) {
	listPeerProto := make([]*pb.GetListPeerResponse_Peer, len(server.Node.ListPeer))
	for index, peer := range server.Node.ListPeer {
		strAddr, err := addrToString(&peer.Addr)
		if err != nil {
			return nil, err
		}

		listPeerProto[index] = &pb.GetListPeerResponse_Peer{
			Addr: strAddr,
		}
	}

	return &pb.GetListPeerResponse{
		ListPeer: listPeerProto,
	}, nil
}
