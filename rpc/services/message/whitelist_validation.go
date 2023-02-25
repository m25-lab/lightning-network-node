package message

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/m25-lab/lightning-network-node/core_chain_sdk/account"
	"github.com/m25-lab/lightning-network-node/database/models"
	"github.com/m25-lab/lightning-network-node/rpc/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (server *MessageServer) ValidateAddWhitelist(ctx context.Context, req *pb.SendMessageRequest, fromAddress string, toAccount *account.PrivateKeySerialized) error {
	//unmarshal req.data
	var addWhitelist models.AddWhitelistData
	if err := json.Unmarshal([]byte(req.Data), &addWhitelist); err != nil {
		return errors.New("invalid data")
	}

	//get Account from sender
	fromAccountFromPayload := account.NewPKAccount(addWhitelist.Pubkey)
	if fromAddress != fromAccountFromPayload.AccAddress().String() {
		return errors.New("invalid data")
	}

	return nil
}

func (server *MessageServer) ValidateAcceptAddWhitelist(ctx context.Context, req *pb.SendMessageRequest, fromAddress string, toAccount *account.PrivateKeySerialized) error {
	//unmarshal req.data
	var addWhitelist models.AddWhitelistData
	if err := json.Unmarshal([]byte(req.Data), &addWhitelist); err != nil {
		return errors.New("invalid data")
	}

	//get Account from sender
	fromAccountFromPayload := account.NewPKAccount(addWhitelist.Pubkey)
	if fromAddress != fromAccountFromPayload.AccAddress().String() {
		return errors.New("invalid data")
	}

	//check exist acceptMessageId
	existMessage, err := server.Node.Repository.Message.FindOneById(ctx, toAccount.AccAddress().String(), req.ReliedMessageId)
	if err != nil {
		return err
	}
	if existMessage.IsReplied {
		return errors.New("invalid data")
	}
	if existMessage.Users[0] != req.To || existMessage.Users[1] != req.From {
		return errors.New("invalid data")
	}

	//create multisig
	acc := account.NewAccount()
	multisigAddr, _, _ := acc.CreateMulSigAccountFromTwoAccount(fromAccountFromPayload.PublicKey(), toAccount.PublicKey(), 2)

	server.Node.Repository.Whitelist.InsertOne(ctx,
		&models.Whitelist{
			ID:             primitive.NewObjectID(),
			Owner:          toAccount.AccAddress().String(),
			PartnerAddress: req.From,
			PartnerPubkey:  fromAccountFromPayload.PublicKey().String(),
			MultiAddress:   multisigAddr,
			MultiPubkey:    "",
		})

	existMessage.IsReplied = true
	server.Node.Repository.Message.Update(ctx, existMessage.ID, existMessage)

	return nil
}
