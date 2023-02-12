package message

import (
	"context"
	"strings"

	"github.com/m25-lab/lightning-network-node/database/models"
	"github.com/m25-lab/lightning-network-node/rpc/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (server *MessageServer) SendMessage(ctx context.Context, req *pb.SendMessageRequest) (*pb.SendMessageResponse, error) {
	if (req.Action == "AddWhitelist") && (req.Data == "") {

	}

	if err := server.Node.Repository.Message.InsertOne(
		ctx,
		&models.Message{
			ID:        primitive.NewObjectID(),
			ChannelID: req.ChannelId,
			Action:    req.Action,
			Data:      req.Data,
			Users:     strings.Split(req.Users, ","),
		},
	); err != nil {
		return nil, err
	}

	return &pb.SendMessageResponse{}, nil
}
