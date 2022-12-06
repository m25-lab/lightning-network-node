package p2pserver

import (
	"context"

	"github.com/m25-lab/lightning-network-node/node/rpc/pb"
	"google.golang.org/grpc/peer"
)

type PeerToPeerHandler struct {
}

func (server *PeerToPeerServer) Connect(ctx context.Context, req *pb.ConnectRequest) (*pb.ConnectResponse, error) {
	p, _ := peer.FromContext(ctx)

	err := server.Node.AddNewPeer(p.Addr)
	if err != nil {
		return nil, err
	}

	listPeer, err := server.Node.ListPeer.Proto()
	if err != nil {
		return nil, err
	}

	listChannel, err := server.Node.ListOpenChannel.Proto()
	if err != nil {
		return nil, err
	}

	return &pb.ConnectResponse{
		ListPeer:    listPeer,
		ListChannel: listChannel,
	}, nil
}

func (server *PeerToPeerServer) GetListPeer(ctx context.Context, req *pb.GetListPeerRequest) (*pb.GetListPeerResponse, error) {
	listPeerProto, err := server.Node.ListPeer.Proto()
	if err != nil {
		return nil, err
	}

	return &pb.GetListPeerResponse{
		ListPeer: listPeerProto,
	}, nil
}
