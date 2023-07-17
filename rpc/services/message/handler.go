package message

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	channelTypes "github.com/m25-lab/channel/x/channel/types"
	"github.com/m25-lab/lightning-network-node/client"
	"github.com/m25-lab/lightning-network-node/core_chain_sdk/channel"
	"github.com/m25-lab/lightning-network-node/core_chain_sdk/common"
	"go.mongodb.org/mongo-driver/mongo"
	"strconv"
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
			ChannelID:       req.ChannelId,
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

func (server *MessageServer) SendSecret(ctx context.Context, req *pb.SendSecretMessage) (*pb.SendSecretResponse, error) {
	//rehash
	if common.ToHashCode(req.MySecret) != req.MyHashcode {
		return &pb.SendSecretResponse{
			Secret:    "",
			ErrorCode: "secret for" + req.MyHashcode + "not match",
		}, nil
	}

	ownHashcode, err := server.Node.Repository.ExchangeHashcode.FindByOwnHash(ctx, req.PartnerHashcode)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &pb.SendSecretResponse{
				Secret:    "",
				ErrorCode: "no record for hash",
			}, nil
		} else {
			return &pb.SendSecretResponse{
				Secret:    "",
				ErrorCode: err.Error(),
			}, nil
		}
	}

	if ownHashcode.PartnerHashcode != req.MyHashcode {
		return &pb.SendSecretResponse{
			Secret:    "",
			ErrorCode: "hashes mismatch",
		}, nil
	}

	input := models.ExchangeHashcodeData{
		MySecret:        ownHashcode.MySecret,
		MyHashcode:      ownHashcode.MyHashcode,
		PartnerHashcode: ownHashcode.PartnerHashcode,
		PartnerSecret:   req.MySecret,
		ChannelID:       ownHashcode.ChannelID,
	}

	err = server.Node.Repository.ExchangeHashcode.UpdateSecret(context.Background(), &input)
	if err != nil {
		println(err.Error(), "; cant update secret")
	}

	return &pb.SendSecretResponse{
		Secret:    ownHashcode.MySecret,
		ErrorCode: "",
	}, nil
}

func (server *MessageServer) BroadcastNoti(ctx context.Context, req *pb.BroadcastNotiMessage) (*pb.BroadcastNotiResponse, error) {
	//Thong bao len Tele
	//khac nhau cho commit cu~/ commit / phan biet bang get secret
	toAddress := strings.Split(req.To, "@")[0]

	existToAddress, err := server.Node.Repository.Address.FindByAddress(ctx, toAddress)
	if err != nil {
		return &pb.BroadcastNotiResponse{
			Response:  err.Error(),
			ErrorCode: "1005",
		}, nil
	}
	toAccount, _ := account.NewAccount().ImportAccount(existToAddress.Mnemonic)

	multisigAddr := strings.Split(req.Index, ":")[0]
	channelID := multisigAddr + ":token:1"

	secretRes, err := server.Node.Repository.ExchangeHashcode.FindByPartnerHash(ctx, req.HashCode)

	clientIdS, err := strconv.ParseInt(existToAddress.ClientId, 10, 64)
	if err != nil {
		return &pb.BroadcastNotiResponse{
			Response:  err.Error(),
			ErrorCode: "1006",
		}, nil
	}

	telMsg := tgbotapi.NewMessage(clientIdS, "")
	telMsg.ParseMode = "Markdown"

	//find last commitment
	lastestCommitment, err := server.Node.Repository.Message.FindOneByChannelIDWithAction(
		context.Background(),
		toAddress,
		channelID,
		models.ExchangeCommitment,
	)

	//CURRENT
	if secretRes.PartnerSecret == "" {

		telMsg.Text = fmt.Sprintf("* Partner broadcasted a CURRENT commitment* \n Channel ID: `%s` \n Partner: `%s` \n Please check your balance.", channelID, req.From)
		_, err = server.Client.Bot.Send(telMsg)
		if err != nil {
			return &pb.BroadcastNotiResponse{
				Response:  err.Error(),
				ErrorCode: "1006",
			}, nil
		}

		//update last commitment
		lastestCommitment.IsReplied = true
		err = server.Node.Repository.Message.Update(context.Background(), lastestCommitment.ID, lastestCommitment)
		if err != nil {
			return &pb.BroadcastNotiResponse{
				Response:  err.Error(),
				ErrorCode: "1006",
			}, nil
		}
		return &pb.BroadcastNotiResponse{
			Response:  "",
			ErrorCode: "",
		}, nil
	}
	//OLD
	// Neu cu~/ Build & broacast withdrawHashlock + thong bao len Tele
	telMsg.Text = fmt.Sprintf("* Partner broadcasted an OLD commitment* \n Channel ID: `%s` \n Partner: `%s` \n Please check your balance. \n Withdrawing Hashlock...", channelID, req.From)
	_, err = server.Client.Bot.Send(telMsg)
	if err != nil {
		return &pb.BroadcastNotiResponse{
			Response:  err.Error(),
			ErrorCode: "1006",
		}, nil
	}

	wdHashlockMsg := channelTypes.MsgWithdrawHashlock{
		Creator: toAddress,
		To:      toAddress,
		Index:   req.Index,
		Secret:  secretRes.PartnerSecret,
	}
	withdrawRequest := channel.SignMsgRequest{
		Msg:      &wdHashlockMsg,
		GasLimit: 100000,
		GasPrice: "0token",
	}
	_, _, err = client.BroadcastTx(server.Client.ClientCtx, toAccount, withdrawRequest)
	if err != nil {
		telMsg.Text = err.Error()
	} else {
		telMsg.Text = fmt.Sprintf("⚡ *Broadcast Withdraw-Hashlock Message successfully.* \n" +
			"Please check your balance.")
	}
	_, err = server.Client.Bot.Send(telMsg)
	if err != nil {
		return &pb.BroadcastNotiResponse{
			Response:  err.Error(),
			ErrorCode: "1006",
		}, nil
	}
	return &pb.BroadcastNotiResponse{
		Response:  "",
		ErrorCode: "",
	}, nil
}
