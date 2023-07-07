package message

import (
	"context"
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strconv"

	"github.com/m25-lab/lightning-network-node/core_chain_sdk/account"
	"github.com/m25-lab/lightning-network-node/core_chain_sdk/channel"
	"github.com/m25-lab/lightning-network-node/database/models"
	"github.com/m25-lab/lightning-network-node/rpc/pb"
)

func (server *MessageServer) ValidateExchangeCommitment(ctx context.Context, req *pb.SendMessageRequest, fromAccount *account.PKAccount, toAccount *account.PrivateKeySerialized, clientId string, ownAddr string) (*pb.SendMessageResponse, error) {
	selfAddress := toAccount.AccAddress().String() + "@" + server.Node.Config.LNode.External
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

	strSig, err := channelClient.SignMultisigTxFromOneAccount(signCommitmentMsg, toAccount, multiSigPubkey, myCommitmentPayload.IsFirstCommitment)
	if err != nil {
		return &pb.SendMessageResponse{
			Response:  err.Error(),
			ErrorCode: "1006",
		}, nil
	}

	myCommitmentMsg := channelClient.CreateCommitmentMsg(
		multisigAddr,
		toAccount.AccAddress().String(),
		myCommitmentPayload.CoinToCreator,
		fromAccount.AccAddress().String(),
		myCommitmentPayload.CoinToHtlc,
		exchangeHashcodeData.MyHashcode,
	)
	ownsignCommitmentMsg := channel.SignMsgRequest{
		Msg:      myCommitmentMsg,
		GasLimit: 100000,
		GasPrice: "0token",
	}
	ownstrSig, err := channelClient.SignMultisigTxFromOneAccount(ownsignCommitmentMsg, toAccount, multiSigPubkey, myCommitmentPayload.IsFirstCommitment)
	if err != nil {
		return &pb.SendMessageResponse{
			Response:  err.Error(),
			ErrorCode: "1006",
		}, nil
	}
	myCommitmentPayload.OwnSignature = ownstrSig
	savedData, err := json.Marshal(myCommitmentPayload)
	if err != nil {
		return &pb.SendMessageResponse{
			Response:  err.Error(),
			ErrorCode: "1006",
		}, nil
	}
	messageId := primitive.NewObjectID()
	err = server.Node.Repository.Message.InsertOne(context.Background(), &models.Message{
		ID:         messageId,
		OriginalID: messageId,
		ChannelID:  req.ChannelId,
		Action:     models.ExchangeCommitment,
		Owner:      toAccount.AccAddress().String(),
		Data:       string(savedData),
		Users:      []string{req.To, req.From},
		IsReplied:  false,
	})
	if err != nil {
		return &pb.SendMessageResponse{
			Response:  err.Error(),
			ErrorCode: "1006",
		}, nil
	}

	partnerCommitmentPayload, err := json.Marshal(models.CreateCommitmentData{
		Creator:           commitmentMsg.Creator,
		ChannelID:         commitmentMsg.ChannelID,
		From:              commitmentMsg.From,
		Timelock:          commitmentMsg.Timelock,
		ToTimelockAddr:    commitmentMsg.ToTimelockAddr,
		ToHashlockAddr:    commitmentMsg.ToHashlockAddr,
		CoinToCreator:     commitmentMsg.CoinToCreator.Amount.Int64(),
		CoinToHtlc:        commitmentMsg.CoinToHtlc.Amount.Int64(),
		Hashcode:          commitmentMsg.Hashcode,
		PartnerSignature:  strSig,
		IsFirstCommitment: myCommitmentPayload.IsFirstCommitment,
	})
	if err != nil {
		return &pb.SendMessageResponse{
			Response:  err.Error(),
			ErrorCode: "1006",
		}, nil
	}
	if myCommitmentPayload.FwdDest != "" && myCommitmentPayload.HashcodeDest != "" {
		if myCommitmentPayload.FwdDest != ownAddr {
			go func() {
				//find next
				nextHop, err := server.Client.Node.Repository.Routing.FindByDestAndBroadcastId(ctx, selfAddress, myCommitmentPayload.FwdDest, myCommitmentPayload.HashcodeDest)
				if err != nil {
					println("Fwd Commitment: nextHop-FindByDestAndBroadcastId:", err.Error())
					return
				}
				//find receivercommit
				rC, err := server.Node.Repository.FwdCommitment.FindReceiverCommitByDestHash(ctx, myCommitmentPayload.HashcodeDest)
				if err != nil {
					println("Fwd Commitment: rC-FindReceiverCommitByDestHash:", err.Error())
					return
				}
				rCData := models.ReceiverCommitment{}
				err = json.Unmarshal([]byte(rC.Data), &rCData)
				if err != nil {
					println("Fwd Commitment: nextHop-Unmarshal:", err.Error())
					return
				}
				server.Client.LnTransfer(clientId, nextHop.NextHop, rCData.CoinTransfer, &myCommitmentPayload.FwdDest, &myCommitmentPayload.HashcodeDest)
			}()
		}
	}
	clientIdS, err := strconv.ParseInt(clientId, 10, 64)
	if err != nil {
		return &pb.SendMessageResponse{
			Response:  err.Error(),
			ErrorCode: "1006",
		}, nil
	}
	telMsg := tgbotapi.NewMessage(clientIdS, "")
	telMsg.ParseMode = "Markdown"
	telMsg.Text = fmt.Sprintf("*Balance Update* \n Channel ID: `%s` \n Partner: `%s` \n Your balance: `%d` \n Partner balance: `%d` \n Commitment ID: `%s`", myCommitmentPayload.ChannelID, req.From, myCommitmentPayload.CoinToHtlc, myCommitmentPayload.CoinToCreator, messageId.Hex())
	_, err = server.Client.Bot.Send(telMsg)
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
