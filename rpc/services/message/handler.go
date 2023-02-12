package message

import (
	"context"

	"github.com/m25-lab/lightning-network-node/database/models"
	"github.com/m25-lab/lightning-network-node/rpc/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (server *MessageServer) SendMessage(ctx context.Context, req *pb.SendMessageRequest) (*pb.SendMessageResponse, error) {
	if (req.Action == "AddWhitelist") && (req.Data == "") {
		err := server.ValidateAddWhitelist(ctx, req)
		if err != nil {
			return &pb.SendMessageResponse{
				Response:  "",
				ErrorCode: "1000",
			}, nil
		}
	}

	if err := server.Node.Repository.Message.InsertOne(
		ctx,
		&models.Message{
			ID:        primitive.NewObjectID(),
			ChannelID: req.ChannelId,
			Action:    req.Action,
			Data:      req.Data,
			Users:     []string{req.From, req.To},
		},
	); err != nil {
		return nil, err
	}

	return &pb.SendMessageResponse{}, nil
}
