package routing

import (
	"context"
	"github.com/m25-lab/lightning-network-node/rpc/pb"
)

func (server *RoutingServer) ProcessInvoiceSecret(ctx context.Context, req *pb.InvoiceSecretMessage) (*pb.RoutingResponse, error) {
	//check hash
	err := server.ValidateInvoiceSecret(ctx, req)
	if err != nil {
		return &pb.RoutingResponse{
			Response:  err.Error(),
			ErrorCode: "ValidateInvoiceSecret",
		}, nil
	}

	//luu DB

	//find dest trong routing table va gui tiep nguyen cuc toi dia chi do

	// if minh la dest, chuyen sang giai doan trade commitment

	return &pb.RoutingResponse{
		Response:  "success",
		ErrorCode: "",
	}, nil
}
