package message

import (
	"context"
	"encoding/json"

	"github.com/m25-lab/lightning-network-node/core_chain_sdk/account"
	"github.com/m25-lab/lightning-network-node/core_chain_sdk/channel"
	"github.com/m25-lab/lightning-network-node/database/models"
	"github.com/m25-lab/lightning-network-node/rpc/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (server *MessageServer) ValidateExchangeCommitment(ctx context.Context, req *pb.SendMessageRequest, fromAccount *account.PKAccount, toAccount *account.PrivateKeySerialized) (*pb.SendMessageResponse, error) {
	multisigAddr, multiSigPubkey, _ := account.NewAccount().CreateMulSigAccountFromTwoAccount(fromAccount.PublicKey(), toAccount.PublicKey(), 2)

	var myCommitmentPayload models.CreateCommitmentData
	if err := json.Unmarshal([]byte(req.Data), &myCommitmentPayload); err != nil {
		return &pb.SendMessageResponse{
			Response:  err.Error(),
			ErrorCode: "1006",
		}, nil
	}

	exchangeHashcodeMessage, err := server.Node.Repository.Message.FindOneByChannelID(context.Background(), toAccount.AccAddress().String(), multisigAddr+":token:1")
	if err != nil {
		return &pb.SendMessageResponse{
			Response:  err.Error(),
			ErrorCode: "1006",
		}, nil
	}
	if exchangeHashcodeMessage.Action != models.ExchangeHashcode {
		return &pb.SendMessageResponse{
			Response:  "partner has not sent hashcode yet",
			ErrorCode: "1006",
		}, nil
	}

	var exchangeHashcodeData models.ExchangeHashcodeData
	err = json.Unmarshal([]byte(exchangeHashcodeMessage.Data), &exchangeHashcodeData)
	if err != nil {
		return nil, err
	}

	if exchangeHashcodeData.MyHashcode != myCommitmentPayload.Hashcode ||
		myCommitmentPayload.From != multisigAddr ||
		myCommitmentPayload.ToTimelockAddr != toAccount.AccAddress().String() ||
		myCommitmentPayload.ToHashlockAddr != fromAccount.AccAddress().String() ||
		myCommitmentPayload.Creator != multisigAddr ||
		myCommitmentPayload.ChannelID != multisigAddr+":token:1" ||
		myCommitmentPayload.Timelock != 100 {
		return &pb.SendMessageResponse{
			Response:  "partner hashcode is not correct",
			ErrorCode: "1006",
		}, nil
	}

	// save my commitment
	//@Todo check signature
	messageId := primitive.NewObjectID()
	err = server.Node.Repository.Message.InsertOne(context.Background(), &models.Message{
		ID:         messageId,
		OriginalID: messageId,
		ChannelID:  req.ChannelID,
		Action:     models.ExchangeCommitment,
		Owner:      toAccount.AccAddress().String(),
		Data:       req.Data,
		Users:      []string{req.To, req.From},
		IsReplied:  false,
	})
	if err != nil {
		return &pb.SendMessageResponse{
			Response:  err.Error(),
			ErrorCode: "1006",
		}, nil
	}

	channelClient := channel.NewChannel(*server.Client.ClientCtx)
	commitmentMsg := channelClient.CreateCommitmentMsg(
		multisigAddr,
		fromAccount.AccAddress().String(),
		myCommitmentPayload.CoinToHtlc,
		toAccount.AccAddress().String(),
		myCommitmentPayload.CoinToCreator,
		exchangeHashcodeData.PartnerHashcode,
	)
	signCommitmentMsg := channel.SignMsgRequest{
		Msg:      commitmentMsg,
		GasLimit: 100000,
		GasPrice: "0token",
	}

	strSig, err := channelClient.SignMultisigTxFromOneAccount(signCommitmentMsg, toAccount, multiSigPubkey)
	if err != nil {
		return &pb.SendMessageResponse{
			Response:  err.Error(),
			ErrorCode: "1006",
		}, nil
	}

	partnerCommitmentPayload, err := json.Marshal(models.CreateCommitmentData{
		Creator:          commitmentMsg.Creator,
		ChannelID:        commitmentMsg.ChannelID,
		From:             commitmentMsg.From,
		Timelock:         commitmentMsg.Timelock,
		ToTimelockAddr:   commitmentMsg.ToTimelockAddr,
		ToHashlockAddr:   commitmentMsg.ToHashlockAddr,
		CoinToCreator:    commitmentMsg.CoinToCreator.Amount.Int64(),
		CoinToHtlc:       commitmentMsg.CoinToHtlc.Amount.Int64(),
		Hashcode:         commitmentMsg.Hashcode,
		PartnerSignature: strSig,
	})
	if err != nil {
		return &pb.SendMessageResponse{
			Response:  err.Error(),
			ErrorCode: "1006",
		}, nil
	}

	return &pb.SendMessageResponse{
		Response:  string(partnerCommitmentPayload),
		ErrorCode: "",
	}, nil
}
