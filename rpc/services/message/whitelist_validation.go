package message

import (
	"context"
	"encoding/json"
	"errors"
	"strings"

	"github.com/m25-lab/lightning-network-node/core_chain_sdk/account"
	"github.com/m25-lab/lightning-network-node/database/models"
	"github.com/m25-lab/lightning-network-node/rpc/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (server *MessageServer) ValidateAddWhitelist(ctx context.Context, req *pb.SendMessageRequest) error {
	//get Address from Address@IP
	fromAddress := strings.Split(req.From, "@")[0]

	//unmarshal req.data
	var addWhitelist models.AddWhitelistData
	if err := json.Unmarshal([]byte(req.Data), &addWhitelist); err != nil {
		return errors.New("invalid data")
	}

	//get Account from sender
	fromAccount := account.NewPKAccount(addWhitelist.Pubkey)
	if fromAccount.AccAddress().String() != fromAddress {
		return errors.New("invalid data")
	}

	return nil
}

func (server *MessageServer) ValidateAcceptAddWhitelist(ctx context.Context, req *pb.SendMessageRequest) error {
	//get Address from Address@IP
	fromAddress := strings.Split(req.From, "@")[0]
	toAddress := strings.Split(req.To, "@")[0]

	//unmarshal req.data
	var addWhitelist models.AddWhitelistData
	if err := json.Unmarshal([]byte(req.Data), &addWhitelist); err != nil {
		return errors.New("invalid data")
	}

	//get Account from sender
	fromAccount := account.NewPKAccount(addWhitelist.Pubkey)
	if fromAccount.AccAddress().String() != fromAddress {
		return errors.New("invalid data")
	}

	//check reciver account existed
	existToAddress, err := server.Node.Repository.Address.FindByAddress(ctx, toAddress)
	if err != nil {
		return err
	}
	toAccount := account.NewPKAccount(existToAddress.Pubkey)

	//check exist acceptMessageId
	existMessage, err := server.Node.Repository.Message.FindOneById(ctx, toAddress, req.ReliedMessageId)
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
	multisigAddr, _, _ := acc.CreateMulSigAccountFromTwoAccount(fromAccount.PublicKey(), toAccount.PublicKey(), 2)

	server.Node.Repository.Whitelist.InsertOne(ctx,
		&models.Whitelist{
			ID:             primitive.NewObjectID(),
			Owner:          toAddress,
			PartnerAddress: fromAccount.AccAddress().String(),
			PartnerPubkey:  fromAccount.PublicKey().String(),
			MultiAddress:   multisigAddr,
			MultiPubkey:    "",
		})

	existMessage.IsReplied = true
	server.Node.Repository.Message.Update(ctx, existMessage.ID, existMessage)

	return nil
}
