package routing

import (
	"context"
	"github.com/m25-lab/lightning-network-node/database/models"
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
	data := models.FwdSecret{
		HashcodeDest: req.Hashcode,
		Secret:       req.Secret,
	}
	if err := server.Node.Repository.FwdSecret.InsertSecret(ctx, &data); err != nil {
		return &pb.RoutingResponse{
			Response:  err.Error(),
			ErrorCode: "InsertSecret",
		}, nil
	}
	//find dest trong routing table va gui tiep nguyen cuc toi dia chi do
	//TODO: call ProcessInvoiceSecret cho thang tiep theo

	// if minh la dest, chuyen sang giai doan trade commitment

	return &pb.RoutingResponse{
		Response:  "success",
		ErrorCode: "",
	}, nil
}
