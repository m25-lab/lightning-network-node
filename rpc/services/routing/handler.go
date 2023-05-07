package routing

import (
	"context"
	"github.com/m25-lab/lightning-network-node/database/models"
	"github.com/m25-lab/lightning-network-node/rpc/pb"
	"go.mongodb.org/mongo-driver/mongo"
	"strings"
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

	split := strings.Split(req.Dest, "@")
	destAddr := split[0]
	toEndpoint := split[1]

	_, err = server.Node.Repository.Address.FindByAddress(ctx, destAddr)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			go func() {
				nextEntry, err := server.Node.Repository.RoutingEntry.FindByDestAndHash(ctx, req.Dest, req.Hashcode)
				if err != nil {
					println("nextEntry_FindByDestAndHash", err.Error())
				}
				rpcClient := pb.NewRoutingClient(server.Client.CreateConn(toEndpoint))
				_, err = rpcClient.ProcessInvoiceSecret(ctx, &pb.InvoiceSecretMessage{
					Hashcode: req.Hashcode,
					Secret:   req.Secret,
					Dest:     nextEntry.Dest,
				})
				if err != nil {
					println("ProcessInvoiceSecret", err.Error())
				}
			}()
			return &pb.RoutingResponse{
				Response:  "success",
				ErrorCode: "",
			}, nil
		} else {
			return &pb.RoutingResponse{
				Response:  err.Error(),
				ErrorCode: "destAddr_FindByAddress",
			}, nil
		}
	}
	// if minh la dest, chuyen sang giai doan trade commitment
	go func() {

	}()

	return &pb.RoutingResponse{
		Response:  "success",
		ErrorCode: "",
	}, nil
}
