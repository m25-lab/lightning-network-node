package routing

import (
	"context"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/m25-lab/lightning-network-node/database/models"
	"github.com/m25-lab/lightning-network-node/rpc/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RoutingGrpcHandler struct {
}

func (server *RoutingServer) RREQ(ctx context.Context, req *pb.RREQRequest) (*pb.RoutingBaseResponse, error) {
	// Try get by broadcast id
	routings, err := server.Node.Repository.Routing.FindRoutingByBroadcastIDAndType(ctx, req.BroadcastID, models.RoutingTypeDiscovery)
	if err != nil {
		return &pb.RoutingBaseResponse{}, err
	}

	// If have --> return error existed
	if len(routings) > 0 {
		return &pb.RoutingBaseResponse{
			ErrorCode: pb.RoutingErrorCode_RREQ_EXISTED,
		}, fmt.Errorf("RREQ existed")
	}

	// If not check is destination is here
	here := server.GetSelfAddress()
	if here == req.DestinationAddress {
		// If yes --> return OK and go to RREP in background
		go server.StartRREP(here, &pb.RREPRequest{
			SourceAddress:      req.DestinationAddress,
			DestinationAddress: req.SourceAddress,
			BroadcastID:        req.BroadcastID,
			FromAddress:        here,
		})
		return &pb.RoutingBaseResponse{
			ErrorCode: pb.RoutingErrorCode_OK,
		}, nil
	}

	// If not --> save record
	err = server.Node.Repository.Routing.InsertOne(ctx, &models.Routing{
		ID:                 primitive.NewObjectID(),
		Type:               models.RoutingTypeDiscovery,
		BroadcastID:        req.BroadcastID,
		DestinationAddress: req.SourceAddress,
		NextHop:            req.FromAddress,
	})
	if err != nil {
		return &pb.RoutingBaseResponse{}, fmt.Errorf("Insert new RREQ error")
	}

	// Build RREP message
	forwardRREQRequest := *req
	forwardRREQRequest.FromAddress = here

	// Broadcast to all channel opened in background
	neighborNodes, err := server.GetNeighborNodes(here)
	if err != nil {
		return &pb.RoutingBaseResponse{}, fmt.Errorf("Get neighbor nodes error")
	}

	for _, neighborNode := range neighborNodes {
		go server.ForwardRREQ(neighborNode, forwardRREQRequest)
	}

	return &pb.RoutingBaseResponse{
		ErrorCode: pb.RoutingErrorCode_OK,
	}, nil
}

func (server *RoutingServer) RREP(ctx context.Context, req *pb.RREPRequest) (*pb.RoutingBaseResponse, error) {
	// Try get by broadcast id
	routings, err := server.Node.Repository.Routing.FindRoutingByBroadcastIDAndType(ctx, req.BroadcastID, models.RoutingTypeDiscovery)
	if err != nil {
		return &pb.RoutingBaseResponse{}, err
	}

	if len(routings) > 0 {
		return &pb.RoutingBaseResponse{
			ErrorCode: pb.RoutingErrorCode_MORE_THAN_ONE_RREQ_EXISTED,
		}, fmt.Errorf("Have more than one RREQ existed")
	}

	// If have check is source location is here
	here := server.GetSelfAddress()
	if here == req.DestinationAddress {
		// If yes save record then get telegram client id --> push a message
		err = server.Node.Repository.Routing.InsertOne(ctx, &models.Routing{
			ID:                 primitive.NewObjectID(),
			Type:               models.RoutingTypeReply,
			BroadcastID:        req.BroadcastID,
			DestinationAddress: req.DestinationAddress,
			NextHop:            req.FromAddress,
		})
		if err != nil {
			return &pb.RoutingBaseResponse{}, fmt.Errorf("Insert new RREQ error")
		}
		msg := buildMessageFindRoutingSuccess(req.BroadcastID)
		_, _ = server.Client.Bot.Send(msg)
	} else {
		routing := routings[0]
		// If not return ok then build next RREPRequest then return OK and RREP next node in background
		forwardRREPRequest := *req
		forwardRREPRequest.FromAddress = here
		go server.ForwardRREP(routing.NextHop, forwardRREPRequest)
	}

	return &pb.RoutingBaseResponse{
		ErrorCode: pb.RoutingErrorCode_OK,
	}, nil
}

func buildMessageFindRoutingSuccess(broadcastID string) *tgbotapi.MessageConfig {
	msg := new(tgbotapi.MessageConfig)
	msg.Text = fmt.Sprintf("✅ *Find route for %s successfully.", broadcastID)
	msg.ParseMode = "Markdown"
	return msg
}

func (server *RoutingServer) GetSelfAddress() string {
	return server.Node.Config.LNode.External
}

func (server *RoutingServer) GetNeighborNodes(address string) ([]string, error) {
	// TODO
	return []string{}, nil
}

func (server *RoutingServer) StartRREP(toAddress string, req *pb.RREPRequest) error {
	rpcClient := pb.NewRoutingServiceClient(server.Client.CreateConn(toAddress))
	response, err := rpcClient.RREP(context.Background(), req)
	if err != nil {
		log.Println(err.Error() + "-" + response.ErrorCode.String())
		return err
	}
	return nil
}

func (server *RoutingServer) ForwardRREQ(toAddress string, req pb.RREQRequest) error {
	rpcClient := pb.NewRoutingServiceClient(server.Client.CreateConn(toAddress))
	response, err := rpcClient.RREQ(context.Background(), &req)
	if err != nil {
		log.Println(err.Error() + "-" + response.ErrorCode.String())
		return err
	}
	return nil
}

func (server *RoutingServer) ForwardRREP(toAddress string, req pb.RREPRequest) error {
	rpcClient := pb.NewRoutingServiceClient(server.Client.CreateConn(toAddress))
	response, err := rpcClient.RREP(context.Background(), &req)
	if err != nil {
		log.Println(err.Error() + "-" + response.ErrorCode.String())
		return err
	}
	return nil
}
