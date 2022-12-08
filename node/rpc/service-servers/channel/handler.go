package channel

import (
	"context"
	"github.com/m25-lab/lightning-network-node/node/database/mongodb/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"

	"github.com/m25-lab/lightning-network-node/node/rpc/pb"
)

type ChannelGrpcHandler struct {
}

func (c *ChannelServer) OpenChannel(ctx context.Context, req *pb.OpenChannelRequest) (*pb.OpenChannelResponse, error) {
	var payload interface{}
	err := bson.UnmarshalExtJSON([]byte(req.Payload), true, &payload)
	if err != nil {
		return &pb.OpenChannelResponse{
			Response: err.Error(),
		}, nil
	}

	openChannelRequest := model.OpenChannelRequest{
		ID:          primitive.NewObjectID(),
		Status:      "pending",
		FromAddress: req.FromAddress,
		ToAddress:   req.ToAddress,
		SignatureA:  req.Signature,
		Payload:     payload,
		CreatedAt:   primitive.DateTime(time.Now().UnixMilli()),
	}
	_, err = c.Node.Database.ChannelCollection.InsertOne(ctx, openChannelRequest)

	if err != nil {
		return &pb.OpenChannelResponse{
			Response: err.Error(),
		}, nil
	}

	return &pb.OpenChannelResponse{
		Response: "OK",
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
