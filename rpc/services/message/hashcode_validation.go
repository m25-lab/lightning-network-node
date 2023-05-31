package message

import (
	"context"
	"encoding/json"

	"github.com/m25-lab/lightning-network-node/core_chain_sdk/account"
	"github.com/m25-lab/lightning-network-node/core_chain_sdk/common"
	"github.com/m25-lab/lightning-network-node/database/models"
	"github.com/m25-lab/lightning-network-node/rpc/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (server *MessageServer) ValidateExchagneHashcode(ctx context.Context, req *pb.SendMessageRequest, fromAccount *account.PKAccount, toAccount *account.PrivateKeySerialized) (*pb.SendMessageResponse, error) {
	var exchangeHashcodeData models.ExchangeHashcodeData
	if err := json.Unmarshal([]byte(req.Data), &exchangeHashcodeData); err != nil {
		return &pb.SendMessageResponse{
			Response:  err.Error(),
			ErrorCode: "1005",
		}, nil
	}
	if exchangeHashcodeData.PartnerHashcode == "" {
		return &pb.SendMessageResponse{
			Response:  "PartnerHashcode is empty",
			ErrorCode: "1005",
		}, nil
	}

	//random secret
	secret, err := common.RandomSecret()
	if err != nil {
		return nil, err
	}
	hashCode := common.ToHashCode(secret)

	payload, err := json.Marshal(models.ExchangeHashcodeData{
		MySecret:        secret,
		MyHashcode:      hashCode,
		PartnerHashcode: exchangeHashcodeData.PartnerHashcode,
	})
	if err != nil {
		return nil, err
	}

	messageId := primitive.NewObjectID()
	savedMessage := &models.Message{
		ID:         messageId,
		OriginalID: messageId,
		ChannelID:  req.ChannelId,
		Action:     models.ExchangeHashcode,
		Owner:      toAccount.AccAddress().String(),
		Users:      []string{req.To, req.From},
		Data:       string(payload),
		IsReplied:  false,
	}
	if err := server.Node.Repository.Message.InsertOne(context.Background(), savedMessage); err != nil {
		return &pb.SendMessageResponse{
			Response:  err.Error(),
			ErrorCode: "1005",
		}, nil
	}

	exHashcodeData := models.ExchangeHashcodeData{
		MySecret:        secret,
		MyHashcode:      hashCode,
		PartnerHashcode: exchangeHashcodeData.PartnerHashcode,
		ChannelID:       savedMessage.ChannelID,
	}

	err = server.Node.Repository.ExchangeHashcode.InsertSecret(context.Background(), &exHashcodeData)
	if err != nil {
		return nil, err
	}

	payload, err = json.Marshal(models.ExchangeHashcodeData{
		PartnerHashcode: hashCode,
	})
	if err != nil {
		return nil, err
	}

	return &pb.SendMessageResponse{
		Response: string(payload),
	}, nil
}
