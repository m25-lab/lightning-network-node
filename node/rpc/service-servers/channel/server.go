package channel

import "github.com/m25-lab/lightning-network-node/node/rpc/pb"

type ChannelServer struct {
	pb.UnimplementedChannelServiceServer
}

func New() (*ChannelServer, error) {
	return &ChannelServer{}, nil
}
