package routing

import (
	"context"
	"errors"
	"log"

	"github.com/m25-lab/lightning-network-node/core_chain_sdk/common"
	"github.com/m25-lab/lightning-network-node/database/models"
	"github.com/m25-lab/lightning-network-node/rpc/pb"
)

func (server *RoutingServer) ValidateInvoiceSecret(ctx context.Context, req *pb.InvoiceSecretMessage) (*models.FwdMessage, error) {
	log.Println("ValidateInvoiceSecret... run")

	rCommit, err := server.Client.Node.Repository.FwdCommitment.FindReceiverCommitByDestHash(ctx, req.To, req.Hashcode)
	if err != nil {
		// log.Println("to ", req.To)
		// log.Println("hash ", req.Hashcode)
		log.Println("FindReceiverCommitByDestHash...")
		log.Println("err ", err.Error())
		return nil, err
	}

	newHash := common.ToHashCode(req.Secret)
	if newHash != req.Hashcode {
		return nil, errors.New("secret not match hash")
	}

	// receiverCommit := models.ReceiverCommitment{}
	// err = json.Unmarshal([]byte(rCommit.Data), &receiverCommit)
	// if err != nil {
	// 	return nil, err
	// }

	return rCommit, nil
}
