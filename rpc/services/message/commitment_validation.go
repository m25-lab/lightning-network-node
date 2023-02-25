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

	var partnerCommitmentPayload models.CreateCommitmentData
	if err := json.Unmarshal([]byte(req.Data), &partnerCommitmentPayload); err != nil {
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

	if exchangeHashcodeData.MyHashcode != partnerCommitmentPayload.Hashcode ||
		partnerCommitmentPayload.From != multisigAddr ||
		partnerCommitmentPayload.ToTimelockAddr != toAccount.AccAddress().String() ||
		partnerCommitmentPayload.ToHashlockAddr != fromAccount.AccAddress().String() ||
		partnerCommitmentPayload.Creator != multisigAddr ||
		partnerCommitmentPayload.ChannelID != multisigAddr+":token:1" ||
		partnerCommitmentPayload.Timelock != 100 ||
		partnerCommitmentPayload.CoinToCreator != 0 {
		return &pb.SendMessageResponse{
			Response:  "partner hashcode is not correct",
			ErrorCode: "1006",
		}, nil
	}

	//@Todo check signature

	channelClient := channel.NewChannel(*server.Client.ClientCtx)
	commitmentMsg := channelClient.CreateCommitmentMsg(
		multisigAddr,
		fromAccount.AccAddress().String(),
		partnerCommitmentPayload.CoinToHtlc,
		toAccount.AccAddress().String(),
		partnerCommitmentPayload.CoinToCreator,
		exchangeHashcodeData.PartnerHashcode,
	)
	signCommitmentMsg := channel.SignMsgRequest{
		Msg:      commitmentMsg,
		GasLimit: 21000,
		GasPrice: "1token",
	}

	strSig, err := channelClient.SignMultisigTxFromOneAccount(signCommitmentMsg, toAccount, multiSigPubkey)
	if err != nil {
		return &pb.SendMessageResponse{
			Response:  err.Error(),
			ErrorCode: "1006",
		}, nil
	}

	myCommitmentPayload, err := json.Marshal(models.CreateCommitmentData{
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

	// save my commitment
	messageId := primitive.NewObjectID()
	err = server.Node.Repository.Message.InsertOne(context.Background(), &models.Message{
		ID:         messageId,
		OriginalID: messageId,
		ChannelID:  req.ChannelID,
		Action:     models.ExchangeCommitment,
		Owner:      toAccount.AccAddress().String(),
		Data:       string(myCommitmentPayload),
		Users:      []string{req.To, req.From},
		IsReplied:  false,
	})
	if err != nil {
		return &pb.SendMessageResponse{
			Response:  err.Error(),
			ErrorCode: "1006",
		}, nil
	}

	return &pb.SendMessageResponse{
		Response:  string(myCommitmentPayload),
		ErrorCode: "",
	}, nil
}
