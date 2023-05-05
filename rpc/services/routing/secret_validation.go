package routing

import (
	"context"
	"github.com/m25-lab/lightning-network-node/rpc/pb"
)

func (server *RoutingServer) ValidateInvoiceSecret(ctx context.Context, req *pb.InvoiceSecretMessage) error {
	// find 1 receivercommit with destHashcode
	_, err := server.Client.Node.Repository.FwdCommitment.FindReceiverCommitByDestHash(ctx, req.Hashcode)
	if err != nil {
		return err
	}
	//TODO:Re hash the secret into destHashcode

	return nil
}
