package routing

import "github.com/m25-lab/lightning-network-node/node/rpc/pb"

type RoutingServer struct {
	pb.UnimplementedRoutingServiceServer
}

func NewServer() (*RoutingServer, error) {
	return &RoutingServer{}, nil
}
