package routing

import (
	"context"
	"github.com/m25-lab/lightning-network-node/database/models"
	"github.com/m25-lab/lightning-network-node/rpc/pb"
)

func (server *RoutingServer) ValidateInvoiceSecret(ctx context.Context, req *pb.InvoiceSecretMessage) (*models.ReceiverCommitment, error) {
	rCommit, err := server.Client.Node.Repository.FwdCommitment.FindReceiverCommitByDestHash(ctx, req.Hashcode)
	if err != nil {
		return nil, err
	}
	//TODO:Re hash the secret into destHashcode

	//TODO: Returned value need to contain whole dest of Receiver
	return rCommit, nil
}
