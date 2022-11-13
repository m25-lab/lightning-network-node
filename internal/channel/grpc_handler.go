package channel

import (
	"context"

	"github.com/m25-lab/lightning-network-node/internal/pb"
)

type ChannelGrpcHandler struct {
}

func (c *ChannelServer) OpenChannel(ctx context.Context, req *pb.OpenChannelRequest) (*pb.OpenChannelResponse, error) {
	// return nil, status.Errorf(codes.NotFound, "method OpenChannel not implemented")
	return &pb.OpenChannelResponse{
		Response: "Open Channel",
	}, nil
}

func (c *ChannelServer) CreateCommitment(ctx context.Context, req *pb.CreateCommitmentRequest) (*pb.CreateCommitmentResponse, error) {
	// return nil, status.Errorf(codes.NotFound, "method CreateCommitment not implemented")
	return &pb.CreateCommitmentResponse{
		Response: "Create Commitment",
	}, nil
}

func (c *ChannelServer) WithdrawHashlock(ctx context.Context, req *pb.WithdrawHashlockRequest) (*pb.WithdrawHashlockResponse, error) {
	// return nil, status.Errorf(codes.NotFound, "method WithdrawHashlock not implemented")
	return &pb.WithdrawHashlockResponse{
		Response: "Withdraw Hashlock",
	}, nil
}

func (c *ChannelServer) WithdrawTimelock(ctx context.Context, req *pb.WithdrawTimelockRequest) (*pb.WithdrawTimelockResponse, error) {
	// return nil, status.Errorf(codes.NotFound, "method WithdrawTimelock not implemented")
	return &pb.WithdrawTimelockResponse{
		Response: "Withdraw Timelock",
	}, nil
}
