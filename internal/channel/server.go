package channel

import (
	"github.com/m25-lab/lightning-network-node/internal/pb"
)

type ChannelServer struct {
	pb.UnimplementedChannelServiceServer
}

func NewServer() (*ChannelServer, error) {
	return &ChannelServer{}, nil
}
