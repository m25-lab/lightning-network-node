package routing

import (
	"context"

	"github.com/m25-lab/lightning-network-node/node/rpc/pb"
)

type RoutingGrpcHandler struct {
}

func (server *RoutingServer) GetNodeTable(ctx context.Context, req *pb.GetRoutingNodeRequest) (*pb.GetRoutingNodeResponse, error) {
	return &pb.GetRoutingNodeResponse{}, nil
}
