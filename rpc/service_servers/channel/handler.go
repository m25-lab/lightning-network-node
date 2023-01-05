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

	openChannelRequest := models.OpenChannelRequest{
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
		Response: openChannelRequest.ID.Hex(),
	}, nil
}

func (c *ChannelServer) GetChannelById(ctx context.Context, req *pb.GetChannelRequest) (*pb.GetChannelResponse, error) {
	var channelResult models.OpenChannelRequest
	objectId, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return &pb.GetChannelResponse{}, err
	}

	if err := c.Node.Database.ChannelCollection.FindOne(
		ctx,
		bson.M{
			"_id": objectId,
		}).Decode(&channelResult); err != nil {
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
