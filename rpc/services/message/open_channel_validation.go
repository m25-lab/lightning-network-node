package message

import (
	"context"
	"encoding/json"

	"github.com/m25-lab/lightning-network-node/core_chain_sdk/account"
	"github.com/m25-lab/lightning-network-node/core_chain_sdk/channel"
	"github.com/m25-lab/lightning-network-node/database/models"
	"github.com/m25-lab/lightning-network-node/rpc/pb"
)

func (server *MessageServer) ValidateOpenChannel(ctx context.Context, req *pb.SendMessageRequest, fromAccount *account.PKAccount, toAccount *account.PrivateKeySerialized) (*pb.SendMessageResponse, error) {
	multisigAddr, multiSigPubkey, _ := account.NewAccount().CreateMulSigAccountFromTwoAccount(fromAccount.PublicKey(), toAccount.PublicKey(), 2)

	exchangeCommitmentMessage, err := server.Node.Repository.Message.FindOneByChannelID(context.Background(), toAccount.AccAddress().String(), multisigAddr+":token:1")
	if err != nil {
		return &pb.SendMessageResponse{
			Response:  err.Error(),
			ErrorCode: "1007",
		}, nil
	}

	if exchangeCommitmentMessage.Action != models.ExchangeCommitment {
		return &pb.SendMessageResponse{
			Response:  "partner has not sent commitment yet",
			ErrorCode: "1007",
		}, nil
	}

	var exchangeCommitmentData models.CreateCommitmentData
	err = json.Unmarshal([]byte(exchangeCommitmentMessage.Data), &exchangeCommitmentData)
	if err != nil {
		return nil, err
	}

	channelClient := channel.NewChannel(*server.Client.ClientCtx)
	openChannelMsg := channelClient.CreateOpenChannelMsg(
		multisigAddr,
		fromAccount.AccAddress().String(),
		toAccount.AccAddress().String(),
		exchangeCommitmentData.CoinToCreator,
		exchangeCommitmentData.CoinToHtlc,
	)
	signOpenChannelMsg := channel.SignMsgRequest{
		Msg:      openChannelMsg,
		GasLimit: 100000,
		GasPrice: "0token",
	}

	//@TODO: check parter sig corret

	strSig, err := channelClient.SignMultisigTxFromOneAccount(signOpenChannelMsg, toAccount, multiSigPubkey)
	if err != nil {
		return &pb.SendMessageResponse{
			Response:  err.Error(),
			ErrorCode: "1007",
		}, nil
	}

	strSigPayload, err := json.Marshal(models.OpenChannelData{
		StrSig: string(strSig),
	})
	if err != nil {
		return &pb.SendMessageResponse{
			Response:  err.Error(),
			ErrorCode: "1007",
		}, nil
	}

	return &pb.SendMessageResponse{
		Response:  string(strSigPayload),
		ErrorCode: "",
	}, nil
}
