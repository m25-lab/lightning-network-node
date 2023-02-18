package message

import (
	"context"
	"fmt"
	"strings"

	"github.com/m25-lab/lightning-network-node/database/models"
	"github.com/m25-lab/lightning-network-node/rpc/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (server *MessageServer) SendMessage(ctx context.Context, req *pb.SendMessageRequest) (*pb.SendMessageResponse, error) {
	toAddress := strings.Split(req.To, "@")[0]
	existToAddress, err := server.Node.Repository.Address.FindByAddress(ctx, toAddress)

	if err != nil {
		fmt.Printf("err: %v", err)
		return &pb.SendMessageResponse{
			Response:  "",
			ErrorCode: "1004",
		}, nil
	}

	if req.Action == models.AddWhitelist {
		err := server.ValidateAddWhitelist(ctx, req)
		if err != nil {
			return &pb.SendMessageResponse{
				Response:  "",
				ErrorCode: "1000",
			}, nil
		}
	}
	if req.Action == models.AcceptAddWhitelist {
		err := server.ValidateAcceptAddWhitelist(ctx, req)
		if err != nil {
			return &pb.SendMessageResponse{
				Response:  "",
				ErrorCode: "1001",
			}, nil
		}
	}

	messageId, err := primitive.ObjectIDFromHex(req.MessageId)
	if err != nil {
		return nil, err
	}
	msg := &models.Message{
		ID:         primitive.NewObjectID(),
		OriginalID: messageId,
		ChannelID:  req.ChannelId,
		Action:     req.Action,
		Owner:      toAddress,
		Data:       req.Data,
		Users:      []string{req.From, req.To},
	}

	if err := server.Node.Repository.Message.InsertOne(
		ctx,
		msg,
	); err != nil {
		return nil, err
	}
	server.Client.TelegramMsg(existToAddress.ClientId, msg)

	return &pb.SendMessageResponse{
		Response:  "success",
		ErrorCode: "",
	}, nil
}
