package routing

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/m25-lab/lightning-network-node/database/models"
	"github.com/m25-lab/lightning-network-node/rpc/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RoutingGrpcHandler struct {
}

func (server *RoutingServer) RREQ(ctx context.Context, req *pb.RREQRequest) (*pb.RoutingBaseResponse, error) {
	if req == nil {
		return &pb.RoutingBaseResponse{
			ErrorCode: pb.RoutingErrorCode_PARAM_INVALID,
		}, fmt.Errorf("Nil body")
	}

	if !server.CheckSelfIsTargetNode(ctx, req.ToAddress) {
		return &pb.RoutingBaseResponse{
			ErrorCode: pb.RoutingErrorCode_WRONG_NODE,
		}, fmt.Errorf("Wrong node")
	}

	// Try get by broadcast id
	routings, err := server.Node.Repository.Routing.FindRouting(ctx, models.Routing{
		BroadcastID: req.BroadcastID,
		NextHop:     req.ToAddress,
		Owner:       req.ToAddress,
	})
	if err != nil && err.Error() != "NOT_FOUND" {
		return &pb.RoutingBaseResponse{}, err
	}

	// If have --> return error existed
	if len(routings) > 0 {
		return &pb.RoutingBaseResponse{
			ErrorCode: pb.RoutingErrorCode_RREQ_EXISTED,
		}, fmt.Errorf("RREQ existed")
	}

	if !server.CheckSelfIsTargetNode(ctx, req.SourceAddress) {
		err = server.Node.Repository.Routing.InsertOne(ctx, &models.Routing{
			ID:                 primitive.NewObjectID(),
			BroadcastID:        req.BroadcastID,
			DestinationAddress: req.SourceAddress,
			NextHop:            req.ToAddress,
			Owner:              req.ToAddress,
		})
		if err != nil {
			return &pb.RoutingBaseResponse{}, fmt.Errorf("Insert new RREQ error")
		}
	}

	// If not check is destination is selfEndpoint
	if server.CheckIsDestination(ctx, req.DestinationAddress) {
		// If yes --> return OK and go to RREP in background
		go server.StartRREP(req.FromAddress, pb.RREPRequest{
			BroadcastID:        req.BroadcastID,
			ToAddress:          req.FromAddress,
			FromAddress:        req.ToAddress,
			DestinationAddress: req.SourceAddress,
			SourceAddress:      req.DestinationAddress,
		})
	} else {
		// If not --> Forward RREQ
		// Build RREP message
		forwardRREQRequest := *req
		forwardRREQRequest.FromAddress = req.ToAddress

		// Broadcast to all channel opened in background
		neighborNodes, err := server.GetNeighborNodes(forwardRREQRequest.FromAddress)
		if err != nil {
			return &pb.RoutingBaseResponse{}, fmt.Errorf("Get neighbor nodes error")
		}

		if len(neighborNodes) == 0 {
			return &pb.RoutingBaseResponse{
				ErrorCode: pb.RoutingErrorCode_NOT_FOUND_NEIGHBOR_NODE,
			}, fmt.Errorf("Not found any neighbor nodes")
		}

		// fmt.Println("neighborNodes ", neighborNodes)
		for _, neighborNode := range neighborNodes {
			if neighborNode == req.SourceAddress {
				continue
			}
			forwardRREQRequest.ToAddress = neighborNode
			go server.ForwardRREQ(neighborNode, forwardRREQRequest)
		}
	}

	return &pb.RoutingBaseResponse{
		ErrorCode: pb.RoutingErrorCode_OK,
	}, nil
}

func (server *RoutingServer) RREP(ctx context.Context, req *pb.RREPRequest) (*pb.RoutingBaseResponse, error) {
	if req == nil {
		return &pb.RoutingBaseResponse{
			ErrorCode: pb.RoutingErrorCode_PARAM_INVALID,
		}, fmt.Errorf("Nil body")
	}

	if !server.CheckSelfIsTargetNode(ctx, req.ToAddress) {
		return &pb.RoutingBaseResponse{
			ErrorCode: pb.RoutingErrorCode_WRONG_NODE,
		}, fmt.Errorf("Wrong node")
	}

	// Try get by broadcast id
	repRoutings, err := server.Node.Repository.Routing.FindRouting(ctx, models.Routing{
		BroadcastID:        req.BroadcastID,
		DestinationAddress: req.SourceAddress,
		Owner:              req.ToAddress,
	})
	if err != nil && err.Error() != "NOT_FOUND" {
		return &pb.RoutingBaseResponse{}, err
	}

	if len(repRoutings) > 0 {
		return &pb.RoutingBaseResponse{
			ErrorCode: pb.RoutingErrorCode_RREP_EXISTED,
		}, fmt.Errorf("RREP existed")
	}

	if !server.CheckSelfIsTargetNode(ctx, req.SourceAddress) {
		// save record
		err = server.Node.Repository.Routing.InsertOne(ctx, &models.Routing{
			ID:                 primitive.NewObjectID(),
			BroadcastID:        req.BroadcastID,
			DestinationAddress: req.SourceAddress,
			NextHop:            req.FromAddress,
			Owner:              req.ToAddress,
		})
		if err != nil {
			return &pb.RoutingBaseResponse{}, fmt.Errorf("Insert new RREQ error")
		}
	}

	// If have check is source location is selfEndpoint
	if server.CheckIsDestination(ctx, req.DestinationAddress) {
		// If yes get telegram client id --> push a message
		address, err := server.Node.Repository.Address.FindByAddress(ctx, getWalletAddress(req.DestinationAddress))
		if err == nil {
			clientId, err := strconv.ParseInt(address.ClientId, 10, 64)
			if err == nil {
				msg := tgbotapi.NewMessage(clientId, "")
				msg.ParseMode = "Markdown"
				msg.Text = fmt.Sprintf("✅ *Find route for `%s` successfully.*\n", req.BroadcastID)
				_, _ = server.Client.Bot.Send(msg)
			}
		}
	} else {

		reqRoutings, err := server.Node.Repository.Routing.FindRouting(ctx, models.Routing{
			BroadcastID:        req.BroadcastID,
			DestinationAddress: req.DestinationAddress,
			Owner:              req.ToAddress,
		})
		if err != nil && err.Error() != "NOT_FOUND" {
			return &pb.RoutingBaseResponse{}, err
		}

		reqRouting := reqRoutings[0]
		// If not return ok then build next RREPRequest then return OK and RREP next node in background
		forwardRREPRequest := *req
		forwardRREPRequest.FromAddress = req.ToAddress
		forwardRREPRequest.ToAddress = reqRouting.NextHop
		go server.ForwardRREP(reqRouting.NextHop, forwardRREPRequest)
	}

	return &pb.RoutingBaseResponse{
		ErrorCode: pb.RoutingErrorCode_OK,
	}, nil
}

func (server *RoutingServer) CheckSelfIsTargetNode(ctx context.Context, targetAdress string) bool {
	walletAddress, endpoint := extractAdress(targetAdress)
	// check is node have this wallet address
	_, err := server.Node.Repository.Address.FindByAddress(ctx, walletAddress)
	if err != nil {
		return false
	}
	return endpoint == server.GetSelfEndpoint()
}

func (server *RoutingServer) CheckIsDestination(ctx context.Context, destinationAddress string) bool {
	return server.CheckSelfIsTargetNode(ctx, destinationAddress)
}

func getEndpoint(address string) string {
	_, endpoint := extractAdress(address)
	return endpoint
}

func getWalletAddress(address string) string {
	walletAddress, _ := extractAdress(address)
	return walletAddress
}

func extractAdress(address string) (walletAddress, endpoint string) {
	tokens := strings.Split(address, "@")
	if len(tokens) > 0 {
		walletAddress = tokens[0]
	}
	if len(tokens) > 1 {
		endpoint = tokens[1]
	}
	return
}

func (server *RoutingServer) GetSelfEndpoint() string {
	return server.Node.Config.LNode.External
}

func (server *RoutingServer) GetNeighborNodes(address string) ([]string, error) {
	// Query white list
	neighborAdress, err := server.Node.Repository.Whitelist.FindMany(context.Background(), getWalletAddress(address))
	if err != nil {
		return nil, err
	}

	// find endpoint of address
	res := []string{}
	for _, wl := range neighborAdress {
		res = append(res, wl.PartnerAddress)
	}

	return res, nil
}

func (server *RoutingServer) StartRREP(toAddress string, req pb.RREPRequest) error {
	rpcClient := pb.NewRoutingServiceClient(server.Client.CreateConn(getEndpoint(toAddress)))
	// time..Sleep(1 * time.Second)
	response, err := rpcClient.RREP(context.Background(), &req)
	if err != nil {
		log.Println("Resp: ", response)
		log.Println(err.Error())
		return err
	}
	return nil
}

func (server *RoutingServer) ForwardRREQ(toAddress string, req pb.RREQRequest) error {
	rpcClient := pb.NewRoutingServiceClient(server.Client.CreateConn(getEndpoint(toAddress)))
	// time.Sleep(1 * time.Second)
	response, err := rpcClient.RREQ(context.Background(), &req)
	if err != nil {
		log.Println("Resp: ", response)
		log.Println(err.Error())
		return err
	}
	return nil
}

func (server *RoutingServer) ForwardRREP(toAddress string, req pb.RREPRequest) error {
	rpcClient := pb.NewRoutingServiceClient(server.Client.CreateConn(getEndpoint(toAddress)))
	// time.Sleep(1 * time.Second)
	response, err := rpcClient.RREP(context.Background(), &req)
	if err != nil {
		log.Println("Resp: ", response)
		log.Println(err.Error())
		return err
	}
	return nil
}
