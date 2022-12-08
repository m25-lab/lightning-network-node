package nodeInfo

import (
	"context"

	"github.com/m25-lab/lightning-network-node/node/rpc/pb"
)

type NodeInfoGrpcHandler struct {
}

func (server *NodeInfoServer) NodeInfo(ctx context.Context, req *pb.NodeInfoRequest) (*pb.NodeInfoResponse, error) {
	// return nil, status.Errorf(codVes.NotFound, "method OpenChannel not implemented")
	return &pb.NodeInfoResponse{
		ChainId:       server.Node.Config.Node.ChainId,
		Endpoint:      server.Node.Config.Node.Endpoint,
		CoinType:      server.Node.Config.Node.CoinType,
		PrefixAddress: server.Node.Config.Node.PrefixAddress,
		TokenSymbol:   server.Node.Config.Node.TokenSymbol,
	}, nil
}
