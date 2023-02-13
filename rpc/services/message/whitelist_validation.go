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
	fromAccount := account.NewPKAccount(addWhitelist.Publickey)
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
	fromAccount := account.NewPKAccount(addWhitelist.Publickey)
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
	existMessage, err := server.Node.Repository.Message.FindOneById(ctx, req.AcceptMessageId)
	if err != nil {
		return err
	}
	if existMessage.Users[0] != req.To || existMessage.Users[1] != req.From {
		return errors.New("invalid data")
	}

	//create multisig
	acc := account.NewAccount()
	multisigAddr, _, _ := acc.CreateMulSigAccountFromTwoAccount(fromAccount.PublicKey(), toAccount.PublicKey(), 2)

	server.Node.Repository.Whitelist.InsertOne(ctx,
		&models.Whitelist{
			ID:           primitive.NewObjectID(),
			Users:        []string{req.To, req.From},
			Pubkeys:      []string{toAccount.PublicKey().String(), fromAccount.PublicKey().String()},
			MultiAddress: multisigAddr,
			MultiPubkey:  "",
		})

	return nil
}
