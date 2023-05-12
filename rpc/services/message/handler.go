package message

import (
	"context"
	"strings"

	"github.com/m25-lab/lightning-network-node/core_chain_sdk/account"
	"github.com/m25-lab/lightning-network-node/database/models"
	"github.com/m25-lab/lightning-network-node/rpc/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (server *MessageServer) SendMessage(ctx context.Context, req *pb.SendMessageRequest) (*pb.SendMessageResponse, error) {
	toAddress := strings.Split(req.To, "@")[0]

	existToAddress, err := server.Node.Repository.Address.FindByAddress(ctx, toAddress)
	if err != nil {
		return &pb.SendMessageResponse{
			Response:  err.Error(),
			ErrorCode: "1005",
		}, nil
	}
	toAccount, _ := account.NewAccount().ImportAccount(existToAddress.Mnemonic)

	parseFromAddress := strings.Split(req.From, "@")
	fromAddress := parseFromAddress[0]
	var fromAccount *account.PKAccount

	//validate with func before add whitelist
	if req.Action == models.AddWhitelist || req.Action == models.AcceptAddWhitelist {
		if req.Action == models.AddWhitelist {
			err := server.ValidateAddWhitelist(ctx, req, fromAddress, toAccount)
			if err != nil {
				return &pb.SendMessageResponse{
					Response:  err.Error(),
					ErrorCode: "1000",
				}, nil
			}
		} else if req.Action == models.AcceptAddWhitelist {
			err := server.ValidateAcceptAddWhitelist(ctx, req, fromAddress, toAccount)
			if err != nil {
				return &pb.SendMessageResponse{
					Response:  err.Error(),
					ErrorCode: "1001",
				}, nil
			}
		}

		messageId, err := primitive.ObjectIDFromHex(req.MessageId)
		if err != nil {
			return nil, err
		}
		msg := &models.Message{
			ID:              primitive.NewObjectID(),
			OriginalID:      messageId,
			ChannelID:       req.ChannelID,
			Action:          req.Action,
			Owner:           toAddress,
			Data:            req.Data,
			Users:           []string{req.From, req.To},
			ReliedMessageId: req.ReliedMessageId,
			IsReplied:       false,
		}

		if err := server.Node.Repository.Message.InsertOne(
			ctx,
			msg,
		); err != nil {
			return nil, err
		}
		server.Client.TelegramMsg(existToAddress.ClientId, msg, toAccount, fromAccount)

		return &pb.SendMessageResponse{
			Response:  "success",
			ErrorCode: "",
		}, nil
	} else {
		fromAddressFromDB, err := server.Client.Node.Repository.Whitelist.FindOneByPartnerAddress(context.Background(), toAddress, req.From)
		if err != nil {
			return &pb.SendMessageResponse{
				Response:  err.Error(),
				ErrorCode: "1004",
			}, nil
		}
		fromAccount = account.NewPKAccount(fromAddressFromDB.PartnerPubkey)

		if req.Action == models.ExchangeHashcode {
			return server.ValidateExchagneHashcode(ctx, req, fromAccount, toAccount)
		} else if req.Action == models.ExchangeCommitment {
			return server.ValidateExchangeCommitment(ctx, req, fromAccount, toAccount, existToAddress.ClientId, req.To)
		} else if req.Action == models.OpenChannel {
			return server.ValidateOpenChannel(ctx, req, fromAccount, toAccount)
		}
	}

	return &pb.SendMessageResponse{
		Response:  "unknown",
		ErrorCode: "1000",
	}, nil
}
