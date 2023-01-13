package channel

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/m25-lab/lightning-network-node/database/models"
	"github.com/m25-lab/lightning-network-node/rpc/pb"
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

	openChannelRequest := &models.OpenChannelRequest{
		ID:          primitive.NewObjectID(),
		Status:      "pending",
		FromAddress: req.FromAddress,
		ToAddress:   req.ToAddress,
		SignatureA:  req.Signature,
		Payload:     payload,
		CreatedAt:   primitive.DateTime(time.Now().UnixMilli()),
	}

	if err = c.Node.Repository.Channel.InsertOpenChannelRequest(ctx, openChannelRequest); err != nil {
		return &pb.OpenChannelResponse{
			Response: err.Error(),
		}, nil
	}

	return &pb.OpenChannelResponse{
		Response: openChannelRequest.ID.Hex(),
	}, nil
}

func (c *ChannelServer) GetChannelById(ctx context.Context, req *pb.GetChannelRequest) (*pb.GetChannelResponse, error) {
	var channelResult *models.OpenChannelRequest
	channelResult, err := c.Node.Repository.Channel.FindChannelById(ctx, req.Id)
	if err != nil {
		panic(err)
	}

	strPayload, _ := bson.MarshalExtJSON(channelResult.Payload, false, false)

	return &pb.GetChannelResponse{
		Id:          channelResult.ID.Hex(),
		Status:      channelResult.Status,
		FromAddress: channelResult.FromAddress,
		ToAddress:   channelResult.ToAddress,
		SignatureA:  channelResult.SignatureA,
		SignatureB:  channelResult.SignatureB,
		Payload:     string(strPayload),
	}, nil
}

func (c *ChannelServer) CreateCommitment(ctx context.Context, req *pb.CreateCommitmentRequest) (*pb.CreateCommitmentResponse, error) {
	commitmentRequest := &models.Commitment{
		ID:          primitive.NewObjectID(),
		ChannelID:   req.ChannelId,
		FromAddress: req.FromAddress,
		Payload:     req.Payload,
		Signature:   req.Signature,
		CreatedAt:   primitive.Timestamp{T: uint32(time.Now().Unix()), I: 0},
		UpdatedAt:   primitive.Timestamp{T: uint32(time.Now().Unix()), I: 0},
	}

	err := c.Node.Repository.Commitment.Insert(context.Background(), commitmentRequest)
	if err != nil {
		return &pb.CreateCommitmentResponse{
			Response: err.Error(),
		}, nil
	}

	return &pb.CreateCommitmentResponse{
		Response: commitmentRequest.ID.Hex(),
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
