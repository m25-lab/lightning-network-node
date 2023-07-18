package routing

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/m25-lab/lightning-network-node/core_chain_sdk/account"
	"github.com/m25-lab/lightning-network-node/core_chain_sdk/channel"
	"github.com/m25-lab/lightning-network-node/core_chain_sdk/common"
	"github.com/m25-lab/lightning-network-node/database/models"
	"github.com/m25-lab/lightning-network-node/rpc/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (server *RoutingServer) RREQ(ctx context.Context, req *pb.RREQRequest) (*pb.RoutingBaseResponse, error) {
	log.Println("RREQ...")
	if req == nil {
		return &pb.RoutingBaseResponse{
			ErrorCode: pb.RoutingErrorCode_PARAM_INVALID,
			Response:  "Empty request",
		}, nil
	}

	if !server.CheckSelfIsTargetNode(ctx, req.ToAddress) {
		return &pb.RoutingBaseResponse{
			ErrorCode: pb.RoutingErrorCode_WRONG_NODE,
			Response:  "Wrong node",
		}, nil
	}

	rreqData := ExtractDataFromRREQ(req.Data)
	if rreqData == nil {
		return &pb.RoutingBaseResponse{
			ErrorCode: pb.RoutingErrorCode_PARAM_INVALID,
			Response:  "Empty data",
		}, nil
	}

	// Try get by broadcast id
	routings, err := server.Node.Repository.Routing.FindRouting(ctx, models.Routing{
		BroadcastID: req.BroadcastID,
		Owner:       req.ToAddress,
	})
	if err != nil && err.Error() != "NOT_FOUND" {
		return &pb.RoutingBaseResponse{
			ErrorCode: pb.RoutingErrorCode_SOME_THING_WENT_WRONG,
			Response:  fmt.Sprint("FindRouting err: ", err.Error()),
		}, nil
	}

	if len(routings) > 0 {
		for _, routing := range routings {
			// check is existed backward path gen from RREP
			if routing.DestinationAddress == req.DestinationAddress {
				// start RREP in background and response OK
				go server.StartRREP(req.FromAddress, BuildRREPFromRREQ(req))

				return &pb.RoutingBaseResponse{
					ErrorCode: pb.RoutingErrorCode_OK,
				}, nil
			}
		}

		return &pb.RoutingBaseResponse{
			ErrorCode: pb.RoutingErrorCode_RREQ_EXISTED,
		}, nil
	}

	newSequenceNumber := time.Now().Unix()
	if !server.CheckSelfIsTargetNode(ctx, req.SourceAddress) {
		// log.Println("RREQ")
		err = server.Node.Repository.Routing.InsertOne(ctx, &models.Routing{
			ID:                 primitive.NewObjectID(),
			BroadcastID:        req.BroadcastID,
			DestinationAddress: req.SourceAddress,
			NextHop:            req.FromAddress,
			HopCounter:         rreqData.HopCounter + 1,
			Owner:              req.ToAddress,
			SequenceNumber:     newSequenceNumber,
		})
		if err != nil {
			return &pb.RoutingBaseResponse{
				ErrorCode: pb.RoutingErrorCode_INSERT_ROUTING_ERROR,
				Response:  "Insert routing gen from RREQ error",
			}, nil
		}
	}

	// If not check is destination is selfEndpoint
	if server.CheckIsDestination(ctx, req.DestinationAddress) {
		// If yes --> return OK and go to RREP in background
		go server.StartRREP(req.FromAddress, BuildRREPFromRREQ(req))
	} else {
		// If not --> Forward RREQ
		// Build RREP message
		// forwardRREQRequest := *req
		rreqData.HopCounter++
		rreqData.SequenceNumber = newSequenceNumber
		rreqData.RemainReward = rreqData.RemainReward - GenerateRewardFromRREQ(req)
		if rreqData.RemainReward < 0 {
			return &pb.RoutingBaseResponse{
				ErrorCode: pb.RoutingErrorCode_NOT_ENOUGH_REWARD,
				Response:  "Not enough reward",
			}, nil
		}
		byteData, _ := json.Marshal(rreqData)
		forwardRREQRequest := pb.RREQRequest{
			SourceAddress:      req.SourceAddress,
			DestinationAddress: req.DestinationAddress,
			BroadcastID:        req.BroadcastID,
			FromAddress:        req.ToAddress,
			Data:               string(byteData),
		}

		// Broadcast to all channel opened in background
		neighborNodes, err := server.GetNeighborNodes(forwardRREQRequest.FromAddress)
		if err != nil {
			return &pb.RoutingBaseResponse{
				ErrorCode: pb.RoutingErrorCode_SOME_THING_WENT_WRONG,
				Response:  fmt.Sprint("GetNeighborNodes err: ", err.Error()),
			}, nil
		}

		if len(neighborNodes) == 0 {
			return &pb.RoutingBaseResponse{
				ErrorCode: pb.RoutingErrorCode_NOT_FOUND_NEIGHBOR_NODE,
				Response:  "Not found neighbor node",
			}, nil
		}

		// log.Println("neighborNodes ", neighborNodes)
		hasBroadcast := false
		for _, neighborNode := range neighborNodes {

			neighborNodeAddress := neighborNode.PartnerAddress
			if neighborNodeAddress == req.SourceAddress {
				continue
			}

			if neighborNodeAddress == req.FromAddress {
				continue
			}

			a, err := server.GetChannelBalance(ctx, getWalletAddress(req.ToAddress), neighborNode.MultiAddress)
			if err != nil {
				continue
			}
			println("CointToHTLC: ", a.CoinToHtlc)
			println("rreqData.Amount", rreqData.Amount)
			if a.CoinToHtlc > rreqData.Amount {
				forwardRREQRequest.ToAddress = neighborNodeAddress
				go server.ForwardRREQ(neighborNodeAddress, &forwardRREQRequest)
				hasBroadcast = true
			}
		}
		if !hasBroadcast {
			return &pb.RoutingBaseResponse{
				ErrorCode: pb.RoutingErrorCode_NOT_EXIST_NEIGHBOR_SATISFY,
			}, nil
		}
	}

	return &pb.RoutingBaseResponse{
		ErrorCode: pb.RoutingErrorCode_OK,
	}, nil
}

func (server *RoutingServer) RREP(ctx context.Context, req *pb.RREPRequest) (*pb.RoutingBaseResponse, error) {
	log.Println("RREP...")
	if req == nil {
		return &pb.RoutingBaseResponse{
			ErrorCode: pb.RoutingErrorCode_PARAM_INVALID,
			Response:  "Empty request",
		}, nil
	}

	if !server.CheckSelfIsTargetNode(ctx, req.ToAddress) {
		return &pb.RoutingBaseResponse{
			ErrorCode: pb.RoutingErrorCode_WRONG_NODE,
			Response:  "Wrong node",
		}, nil
	}

	rrepData := ExtractDataFromRREP(req.Data)
	if rrepData == nil {
		return &pb.RoutingBaseResponse{
			ErrorCode: pb.RoutingErrorCode_PARAM_INVALID,
			Response:  "Empty data",
		}, nil
	}

	// Try get by broadcast id
	routings, err := server.Node.Repository.Routing.FindRouting(ctx, models.Routing{
		BroadcastID: req.BroadcastID,
		Owner:       req.ToAddress,
	})
	if err != nil && err.Error() != "NOT_FOUND" {
		return &pb.RoutingBaseResponse{
			ErrorCode: pb.RoutingErrorCode_SOME_THING_WENT_WRONG,
			Response:  fmt.Sprint("FindRouting err: ", err.Error()),
		}, nil
	}

	var matchRoutingGenFromRREQ *models.Routing
	newSequenceNumber := time.Now().Unix()
	skipInsert := false
	if len(routings) > 0 {
		for _, routing := range routings {
			// check is existed forward path gen from other RREP
			if routing.DestinationAddress == req.SourceAddress {
				if rrepData.SequenceNumber > routing.SequenceNumber {
					// upsert routing gen from RREP
					routing.SequenceNumber = rrepData.SequenceNumber
					routing.HopCounter = rrepData.HopCounter
					err = server.Node.Repository.Routing.UpdateRoute(ctx, routing.ID, routing)
					if err != nil {
						skipInsert = true
					}
				}

				return &pb.RoutingBaseResponse{
					ErrorCode: pb.RoutingErrorCode_RREP_EXISTED,
					Response:  "Existed forward path gen from other RREP",
				}, nil
			} else if routing.DestinationAddress == req.DestinationAddress {
				// check is existed forward path gen from other RREQ
				matchRoutingGenFromRREQ = routing
				break
			}
		}
	}

	if !skipInsert && !server.CheckSelfIsTargetNode(ctx, req.SourceAddress) {
		// save record
		// log.Println("RREP")
		err = server.Node.Repository.Routing.InsertOne(ctx, &models.Routing{
			ID:                 primitive.NewObjectID(),
			BroadcastID:        req.BroadcastID,
			DestinationAddress: req.SourceAddress,
			NextHop:            req.FromAddress,
			HopCounter:         rrepData.HopCounter + 1,
			Owner:              req.ToAddress,
		})
		if err != nil {
			return &pb.RoutingBaseResponse{
				ErrorCode: pb.RoutingErrorCode_INSERT_ROUTING_ERROR,
				Response:  "Insert routing gen from RREP error",
			}, nil
		}
	}

	// If have check is source location is selfEndpoint
	if server.CheckIsDestination(ctx, req.DestinationAddress) {
		log.Println("YEAH")
		// If yes get telegram client id --> push a message
		address, err := server.Node.Repository.Address.FindByAddress(ctx, getWalletAddress(req.DestinationAddress))
		if err == nil {
			clientId, err := strconv.ParseInt(address.ClientId, 10, 64)
			if err == nil {
				msg := tgbotapi.NewMessage(clientId, "")
				msg.ParseMode = "Markdown"
				if skipInsert {
					msg.Text = fmt.Sprintf("✅ Update new route for `%s` successfully with `%d` hops. Will you start lightning transfer multi hops ?\n", req.BroadcastID, rrepData.HopCounter+1)
				} else {
					msg.Text = fmt.Sprintf("✅ Found route for `%s` successfully with `%d` hops. Will you start lightning transfer multi hops ?\n", req.BroadcastID, rrepData.HopCounter+1)
				}
				msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
					tgbotapi.NewInlineKeyboardRow(
						tgbotapi.NewInlineKeyboardButtonData("Start", fmt.Sprintf("%s:%s", models.StartLnTransferMultiHop, req.BroadcastID)),
					),
				)
				_, err := server.Client.Bot.Send(msg)
				if err != nil {
					log.Println("err :", err.Error())
				}
			}
		}
	} else {
		// If not return ok then build next RREPRequest then return OK and RREP next node in background
		if matchRoutingGenFromRREQ == nil {
			log.Println("Routing gen from RREQ not existed")
			return &pb.RoutingBaseResponse{
				ErrorCode: pb.RoutingErrorCode_ROUTE_NOT_EXISTED,
				Response:  "Routing gen from RREQ not existed",
			}, nil
		}

		rrepData.HopCounter++
		rrepData.SequenceNumber = newSequenceNumber
		byteData, _ := json.Marshal(rrepData)
		forwardRREPRequest := pb.RREPRequest{
			SourceAddress:      req.SourceAddress,
			DestinationAddress: req.DestinationAddress,
			BroadcastID:        req.BroadcastID,
			FromAddress:        req.ToAddress,
			ToAddress:          matchRoutingGenFromRREQ.NextHop,
			Data:               string(byteData),
		}
		go server.ForwardRREP(matchRoutingGenFromRREQ.NextHop, &forwardRREPRequest)
	}

	return &pb.RoutingBaseResponse{
		ErrorCode: pb.RoutingErrorCode_OK,
	}, nil
}

func (server *RoutingServer) RERR(ctx context.Context, req *pb.RERRRequest) (*pb.RoutingBaseResponse, error) {
	log.Println("RERR...")
	rrep, _ := json.Marshal(req)
	log.Println(string(rrep))
	if req == nil {
		return &pb.RoutingBaseResponse{
			ErrorCode: pb.RoutingErrorCode_PARAM_INVALID,
			Response:  "Empty request",
		}, nil
	}

	if !server.CheckSelfIsTargetNode(ctx, req.ToAddress) {
		return &pb.RoutingBaseResponse{
			ErrorCode: pb.RoutingErrorCode_WRONG_NODE,
			Response:  "Wrong node",
		}, nil
	}
	err := server.Node.Repository.Routing.DeletedRoutingByNextHop(ctx, req.ErrorAddress, req.ToAddress)
	if err != nil {
		log.Println("DeletedRoutingByNextHop: ", err.Error())
	} else {
		neighborNodes, err := server.GetNeighborNodes(req.ToAddress)
		if err == nil {
			for _, neighborNode := range neighborNodes {
				neighborNodeAddress := neighborNode.PartnerAddress
				if neighborNodeAddress == req.FromAddress {
					continue
				}
				req.ToAddress = neighborNodeAddress
				go server.StartRERR(neighborNodeAddress, req)
			}
		}
	}

	return &pb.RoutingBaseResponse{
		ErrorCode: pb.RoutingErrorCode_OK,
	}, nil
}

func (server *RoutingServer) GetChannelBalance(ctx context.Context, address string, multisignAddress string) (*models.CreateCommitmentData, error) {
	latestCommitment, err := server.Node.Repository.Message.FindOneByChannelIDWithAction(
		context.Background(),
		address,
		fmt.Sprint(multisignAddress+":token:1"),
		models.ExchangeCommitment,
	)
	if err != nil {
		return nil, err
	}
	if latestCommitment.IsReplied {
		return nil, fmt.Errorf("Channel is broadcasted")
	}

	payload := models.CreateCommitmentData{}
	err = json.Unmarshal([]byte(latestCommitment.Data), &payload)
	if err != nil {
		return nil, err
	}
	return &payload, nil
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

func ExtractDataFromRREQ(dataString string) *models.RREQData {
	var data models.RREQData
	err := json.Unmarshal([]byte(dataString), &data)
	if err != nil {
		return nil
	}

	return &data
}

func ExtractDataFromRREP(dataString string) *models.RREPData {
	var data models.RREPData
	err := json.Unmarshal([]byte(dataString), &data)
	if err != nil {
		return nil
	}
	return &data
}

func (server *RoutingServer) GetSelfEndpoint() string {
	return server.Node.Config.LNode.External
}

func (server *RoutingServer) GetNeighborNodes(address string) ([]models.Whitelist, error) {
	// Query white list
	neighborAdress, err := server.Node.Repository.Whitelist.FindMany(context.Background(), getWalletAddress(address))
	if err != nil {
		return nil, err
	}

	// find endpoint of address
	return neighborAdress, nil
}

func (server *RoutingServer) StartRREP(toAddress string, req *pb.RREPRequest) error {
	conn := server.Client.CreateConn(getEndpoint(toAddress))
	defer conn.Close()
	rpcClient := pb.NewRoutingServiceClient(conn)
	// time.Sleep(1 * time.Second)
	ctxTimeout, cancelFunc := context.WithTimeout(context.Background(), time.Duration(server.Node.Config.LNode.TimeoutRequest)*time.Second)
	defer cancelFunc()
	response, err := rpcClient.RREP(ctxTimeout, req)
	if err != nil {
		log.Println("StartRREP: ", err.Error())
		return err
	} else {
		if response.ErrorCode != pb.RoutingErrorCode_OK {
			log.Println("StartRREP: ", response)
		}
	}
	return nil
}

func (server *RoutingServer) StartRERR(toAddress string, req *pb.RERRRequest) error {
	conn := server.Client.CreateConn(getEndpoint(toAddress))
	defer conn.Close()
	rpcClient := pb.NewRoutingServiceClient(conn)
	// time.Sleep(1 * time.Second)
	ctxTimeout, cancelFunc := context.WithTimeout(context.Background(), time.Duration(server.Node.Config.LNode.TimeoutRequest)*time.Second)
	defer cancelFunc()
	response, err := rpcClient.RERR(ctxTimeout, req)
	if err != nil {
		log.Println("StartRERR: ", err.Error())
		return err
	} else {
		if response.ErrorCode != pb.RoutingErrorCode_OK {
			log.Println("StartRERR: ", response)
		}
	}
	return nil
}

func (server *RoutingServer) ForwardRREQ(toAddress string, req *pb.RREQRequest) error {
	conn := server.Client.CreateConn(getEndpoint(toAddress))
	defer conn.Close()
	rpcClient := pb.NewRoutingServiceClient(conn)
	// time.Sleep(1 * time.Second)
	ctxTimeout, cancelFunc := context.WithTimeout(context.Background(), time.Duration(server.Node.Config.LNode.TimeoutRequest)*time.Second)
	defer cancelFunc()
	response, err := rpcClient.RREQ(ctxTimeout, req)
	if err != nil {
		log.Println("ForwardRREQ: ", err.Error())
		return err
	} else {
		if response.ErrorCode != pb.RoutingErrorCode_OK {
			log.Println("ForwardRREQ: ", response)
		}
		// go server.StartRERR(req.FromAddress, BuildRERRFromRREQ(req, response.Response))
	}
	return nil
}

func (server *RoutingServer) ForwardRREP(toAddress string, req *pb.RREPRequest) error {
	conn := server.Client.CreateConn(getEndpoint(toAddress))
	defer conn.Close()
	rpcClient := pb.NewRoutingServiceClient(conn)
	// time.Sleep(1 * time.Second)
	ctxTimeout, cancelFunc := context.WithTimeout(context.Background(), time.Duration(server.Node.Config.LNode.TimeoutRequest)*time.Second)
	defer cancelFunc()
	response, err := rpcClient.RREP(ctxTimeout, req)
	if err != nil {
		log.Println("ForwardRREP: ", err.Error())
		return err
	} else {
		if response.ErrorCode != pb.RoutingErrorCode_OK {
			log.Println("ForwardRREP: ", response)
			go server.StartRERR(req.FromAddress, BuildRERRFromRREP(req, response.Response))
		}
	}
	return nil
}

func BuildRREPFromRREQ(rreq *pb.RREQRequest) (rrep *pb.RREPRequest) {
	rrepData := models.RREPData{
		HopCounter:     0,
		SequenceNumber: time.Now().Unix(),
	}

	rrepDataByte, _ := json.Marshal(rrepData)

	rrep = &pb.RREPRequest{
		BroadcastID:        rreq.BroadcastID,
		ToAddress:          rreq.FromAddress,
		FromAddress:        rreq.ToAddress,
		DestinationAddress: rreq.SourceAddress,
		SourceAddress:      rreq.DestinationAddress,
		Data:               string(rrepDataByte),
	}
	return
}

func BuildRERRFromRREQ(rreq *pb.RREQRequest, msg string) (rerr *pb.RERRRequest) {
	log.Println("BuildRERRFromRREQ...")
	rerr = &pb.RERRRequest{
		DestinationAddress: rreq.DestinationAddress,
		FromAddress:        rreq.ToAddress,
		ToAddress:          rreq.FromAddress,
		Message:            msg,
	}
	return
}

func BuildRERRFromRREP(rrep *pb.RREPRequest, msg string) (rerr *pb.RERRRequest) {
	log.Println("BuildRERRFromRREP...")
	rerr = &pb.RERRRequest{
		DestinationAddress: rrep.DestinationAddress,
		FromAddress:        rrep.ToAddress,
		ToAddress:          rrep.FromAddress,
		Message:            msg,
	}
	return
}

func (server *RoutingServer) ProcessInvoiceSecret(ctx context.Context, req *pb.InvoiceSecretMessage) (*pb.RoutingBaseResponse, error) {
	log.Println("ProcessInvoiceSecret...run")
	existToAddress, err := server.Node.Repository.Address.FindByAddress(ctx, strings.Split(req.To, "@")[0])
	if err != nil {
		log.Println("FindByAddress... : existToAddress")
		return &pb.RoutingBaseResponse{
			ErrorCode: pb.RoutingErrorCode_VALIDATE_INVOICE_SECRET,
		}, err
	}
	//check hash
	receiverCommit, err := server.ValidateInvoiceSecret(ctx, req)
	if err != nil {
		log.Println("ValidateInvoiceSecret...")
		return &pb.RoutingBaseResponse{
			ErrorCode: pb.RoutingErrorCode_VALIDATE_INVOICE_SECRET,
		}, err
	}

	//luu DB
	data := models.FwdSecret{
		Owner:        req.To,
		HashcodeDest: req.Hashcode,
		Secret:       req.Secret,
	}
	if err := server.Node.Repository.FwdSecret.InsertSecret(ctx, &data); err != nil {
		log.Println("InsertSecret...")
		return &pb.RoutingBaseResponse{
			ErrorCode: pb.RoutingErrorCode_INSERT_SECRET,
		}, nil
	}

	split := strings.Split(req.Dest, "@")
	destAddr := split[0] // A

	fmt.Println("destAddr ", destAddr)
	activeAddress, err := server.Node.Repository.Address.FindByAddress(ctx, destAddr) //check coi A co trong db minh khong
	if err == mongo.ErrNoDocuments || (err == nil && req.To != req.Dest) {
		//minh khong co A hoac co A nhung minh khong phai A
		go func() {
			time.Sleep(1 * time.Second)
			msg := fmt.Sprintf("*Received Invoice Secret* \n"+
				"HashcodeDest: `%s`\n"+
				"Forwarding...", req.Hashcode)
			err := server.Client.SendTele(existToAddress.ClientId, msg)
			if err != nil {
				println("Noti Tele: ", err.Error())
			}
			nextEntry, err := server.Node.Repository.Routing.FindByDestAndBroadcastId(context.Background(), req.To, req.Dest, req.Hashcode)
			if err != nil {
				println("nextEntry_FindByDestAndBroadcastId", err.Error())
			}
			nextEndpoint := strings.Split(nextEntry.NextHop, "@")[1]
			rpcClient := pb.NewRoutingServiceClient(server.Client.CreateConn(nextEndpoint))
			response, err := rpcClient.ProcessInvoiceSecret(context.Background(), &pb.InvoiceSecretMessage{
				From:     req.To,
				To:       nextEntry.NextHop,
				Hashcode: req.Hashcode,
				Secret:   req.Secret,
				Dest:     nextEntry.DestinationAddress,
			})
			if err != nil {
				println("ProcessInvoiceSecret", err.Error())
			}
			if response.ErrorCode != pb.RoutingErrorCode_OK {
				//thangcq: bao Tele
				println("ProcessInvoiceSecret_Inside: ", response.ErrorCode)
			}
		}()
		return &pb.RoutingBaseResponse{
			ErrorCode: pb.RoutingErrorCode_OK,
		}, nil
	}
	if err != nil {
		log.Println("FindByAddress...")
		return &pb.RoutingBaseResponse{
			ErrorCode: pb.RoutingErrorCode_DESTINATION_ADDRESS_FIND_BY_ADDRESS,
		}, nil

	}
	// is dest -> phase commitment
	go func() {
		time.Sleep(1 * time.Second)
		msg := fmt.Sprintf("*Received Invoice Secret* \n"+
			"HashcodeDest: `%s`\n"+
			"*Trading Commitment...*", req.Hashcode)
		err := server.Client.SendTele(existToAddress.ClientId, msg)
		if err != nil {
			println("Noti Tele: ", err.Error())
		}
		invoice, err := server.Node.Repository.Invoice.FindByHashFrom(context.Background(), req.To, req.Hashcode)
		if err != nil {
			fmt.Println("req.To", req.To)
			fmt.Println("req.Hashcode", req.Hashcode)
			println("is dest - commitment :", err.Error())
			return
		}

		receiverCommitData := models.ReceiverCommitment{}
		err = json.Unmarshal([]byte(receiverCommit.Data), &receiverCommitData)
		if err != nil {
			println("Unmarshal receiverCommitData")
		}
		amount := receiverCommitData.CoinTransfer

		_, err = server.Client.LnTransfer(activeAddress.ClientId, receiverCommit.From, amount, &invoice.To, &receiverCommit.HashcodeDest)
		if err != nil {
			println("Trade commitment - LnTransfer:", err.Error())
		}
	}()

	return &pb.RoutingBaseResponse{
		ErrorCode: pb.RoutingErrorCode_OK,
	}, nil
}

func (server *RoutingServer) RequestInvoice(ctx context.Context, req *pb.IREQMessage) (*pb.IREPMessage, error) {
	secret, err := common.RandomSecret()
	if err != nil {
		println("RandomSecret:", err.Error())
		return &pb.IREPMessage{
			ErrorCode: err.Error(),
		}, nil
	}
	hashcode := common.ToHashCode(secret)
	err = server.Node.Repository.Invoice.InsertInvoice(ctx, &models.InvoiceData{
		Amount: req.Amount,
		From:   req.From,
		To:     req.To,
		Hash:   hashcode,
		Secret: secret,
	})
	if err != nil {
		return &pb.IREPMessage{
			ErrorCode: err.Error(),
		}, nil
	}
	return &pb.IREPMessage{
		From:      req.From,
		To:        req.To,
		Hash:      hashcode,
		Amount:    req.Amount,
		ErrorCode: "",
	}, nil
}

func (server *RoutingServer) ProcessFwdMessage(ctx context.Context, req *pb.FwdMessage) (*pb.FwdMessageResponse, error) {
	//Check "to" is active
	toAddress := strings.Split(req.To, "@")[0]
	existToAddress, err := server.Node.Repository.Address.FindByAddress(ctx, toAddress)
	if err != nil {
		return &pb.FwdMessageResponse{
			Response:  err.Error(),
			ErrorCode: "1005",
		}, err
	}
	toAccount, _ := account.NewAccount().ImportAccount(existToAddress.Mnemonic)

	//get "From" public key
	fromAddressFromDB, err := server.Client.Node.Repository.Whitelist.FindOneByPartnerAddress(context.Background(), toAddress, req.From)
	if err != nil {
		return &pb.FwdMessageResponse{
			Response:  err.Error(),
			ErrorCode: "1004",
		}, err
	}
	fromAccount := account.NewPKAccount(fromAddressFromDB.PartnerPubkey)

	//gen multiAddr
	multisigAddr, multiSigPubkey, _ := account.NewAccount().CreateMulSigAccountFromTwoAccount(fromAccount.PublicKey(), toAccount.PublicKey(), 2)

	var myCommitmentPayload models.SenderCommitment
	if err := json.Unmarshal([]byte(req.Data), &myCommitmentPayload); err != nil {
		return &pb.FwdMessageResponse{
			Response:  err.Error(),
			ErrorCode: "1006",
		}, err
	}

	//check hash code htlc
	exchangeHashcodeMessage, err := server.Node.Repository.Message.FindOneByChannelID(context.Background(), toAccount.AccAddress().String(), multisigAddr+":token:1")
	if err != nil {
		return &pb.FwdMessageResponse{
			Response:  err.Error(),
			ErrorCode: "1006",
		}, err
	}
	if exchangeHashcodeMessage.Action != models.ExchangeHashcode {
		return &pb.FwdMessageResponse{
			Response:  "partner has not sent hashcode yet",
			ErrorCode: "1006",
		}, err
	}

	var exchangeHashcodeData models.ExchangeHashcodeData
	err = json.Unmarshal([]byte(exchangeHashcodeMessage.Data), &exchangeHashcodeData)
	if err != nil {
		return nil, err
	}

	//TODO: validate more fields
	if exchangeHashcodeData.MyHashcode != myCommitmentPayload.HashcodeHTLC ||
		myCommitmentPayload.Creator != multisigAddr ||
		myCommitmentPayload.From != strings.Split(req.From, "@")[0] {
		//myCommitmentPayload.ToTimelockAddr != toAccount.AccAddress().String() ||
		//myCommitmentPayload.ToHashlockAddr != fromAccount.AccAddress().String() ||
		//myCommitmentPayload.Channelid != multisigAddr+":token:1" ||
		//myCommitmentPayload.Timelock != 100
		return &pb.FwdMessageResponse{
			Response:  "partner hashcode is not correct",
			ErrorCode: "1006",
		}, err
	}

	//build SenderCommit and sign
	channelClient := channel.NewChannel(*server.Client.ClientCtx)
	senderCMsg := channelClient.CreateSenderCommitmentMsg(
		multisigAddr,
		fromAccount.AccAddress().String(),
		myCommitmentPayload.CoinToSender,
		myCommitmentPayload.CoinToHTLC,
		myCommitmentPayload.CoinTransfer,
		myCommitmentPayload.HashcodeHTLC,
		myCommitmentPayload.HashcodeDest,
		myCommitmentPayload.Hops,
	)

	signSenderCommitmentMsg := channel.SignMsgRequest{
		Msg:      senderCMsg,
		GasLimit: 200000,
		GasPrice: "0token",
	}

	//sign sender commitment in advance
	strSigSender, err := channelClient.SignMultisigTxFromOneAccount(signSenderCommitmentMsg, toAccount, multiSigPubkey, false)
	if err != nil {
		return &pb.FwdMessageResponse{
			Response:  err.Error(),
			ErrorCode: "1006",
		}, err
	}
	//Store DB sender commit with 2 sig
	err = server.Node.Repository.FwdCommitment.InsertFwdMessage(ctx, &models.FwdMessage{
		Action:       req.Action, //SenderCommit
		PartnerSig:   req.Sig,
		OwnSig:       strSigSender,
		Data:         req.Data,
		From:         req.From,
		To:           req.To,
		HashcodeDest: req.HashcodeDest,
	})
	if err != nil {
		return &pb.FwdMessageResponse{
			Response:  err.Error(),
			ErrorCode: "1006",
		}, err
	}

	//Build and sign receiver commit
	receiverCMsg := channelClient.CreateReceiverCommitmentMsg(
		multisigAddr,
		toAddress,
		myCommitmentPayload.CoinToHTLC,
		myCommitmentPayload.CoinToSender,
		myCommitmentPayload.CoinTransfer,
		myCommitmentPayload.HashcodeHTLC,
		myCommitmentPayload.HashcodeDest,
		myCommitmentPayload.Hops,
	)

	signReceiverCommitmentMsg := channel.SignMsgRequest{
		Msg:      receiverCMsg,
		GasLimit: 200000,
		GasPrice: "0token",
	}

	strSigReceiver, err := channelClient.SignMultisigTxFromOneAccount(signReceiverCommitmentMsg, toAccount, multiSigPubkey, false)
	if err != nil {
		return &pb.FwdMessageResponse{
			Response:  err.Error(),
			ErrorCode: "1006",
		}, err
	}

	partnerCommitmentPayload, err := json.Marshal(models.ReceiverCommitment{
		Creator:        receiverCMsg.Creator,
		From:           receiverCMsg.From,
		ChannelID:      receiverCMsg.ChannelID,
		CoinToReceiver: receiverCMsg.CoinToReceiver.Amount.Int64(),
		CoinToHTLC:     receiverCMsg.CoinToHtlc.Amount.Int64(),
		HashcodeHTLC:   receiverCMsg.HashcodeHtlc,
		TimelockHTLC:   receiverCMsg.TimelockHtlc,
		CoinTransfer:   receiverCMsg.CoinTransfer.Amount.Int64(),
		HashcodeDest:   receiverCMsg.HashcodeDest,
		TimelockSender: receiverCMsg.TimelockSender,
		Multisig:       receiverCMsg.Multisig,
	})

	if err != nil {
		return &pb.FwdMessageResponse{
			Response:  err.Error(),
			ErrorCode: "1006",
		}, err
	}
	//find invoice in db, exist => is Dest

	needNext := false
	invoice, err := server.Node.Repository.Invoice.FindByHash(ctx, req.To, req.HashcodeDest)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			needNext = true
		} else {
			return &pb.FwdMessageResponse{
				Response:  err.Error(),
				ErrorCode: "checkIsDest",
			}, err
		}
	}
	if needNext {
		msg := fmt.Sprintf("*Received and Forwarding FWD Commit worth `%d` token.* \n"+
			"HashcodeDest: `%s`", myCommitmentPayload.CoinTransfer, myCommitmentPayload.HashcodeDest)
		err := server.Client.SendTele(existToAddress.ClientId, msg)
		if err != nil {
			println("Noti Tele Fwd: ", err.Error())
		}
		go func() {
			log.Println("a")
			time.Sleep(time.Second * 2)
			log.Println("b")
			err = server.Client.LnTransferMulti(existToAddress.ClientId, req.Dest, myCommitmentPayload.CoinTransfer, &req.HashcodeDest, true, myCommitmentPayload.Hops-1)
			if err != nil {
				println("Trade fwd commitment - LnTransferMulti:", err.Error())
			}
		}()
	} else {
		//is Dest
		//to phase reveal C's secret, call processInvoiceSecret to B
		go func() {
			time.Sleep(time.Second * 2)
			msg := fmt.Sprintf("*Received FWD Commit worth `%d` token.* \n"+
				"HashcodeDest: `%s` \n"+
				"*Revealing Secret...*", myCommitmentPayload.CoinTransfer, myCommitmentPayload.HashcodeDest)
			err := server.Client.SendTele(existToAddress.ClientId, msg)
			if err != nil {
				println("Noti Tele Reveal: ", err.Error())
			}
			fromSplit := strings.Split(req.From, "@")
			rpcClient := pb.NewRoutingServiceClient(server.Client.CreateConn(fromSplit[1]))
			response, err := rpcClient.ProcessInvoiceSecret(context.Background(), &pb.InvoiceSecretMessage{
				Hashcode: invoice.Hash,
				Secret:   invoice.Secret,
				From:     req.To,
				To:       req.From,
				Dest:     invoice.From,
			})
			if err != nil {
				log.Println("ProcessInvoiceSecret2...")
				log.Println("", err.Error())
				// return nil, err
			} else {
				if response.ErrorCode != pb.RoutingErrorCode_OK {
					println("start ProcessInvoiceSecret: ", response.ErrorCode)
				}
			}
		}()
	}

	return &pb.FwdMessageResponse{
		Response:   string(partnerCommitmentPayload),
		PartnerSig: strSigReceiver,
		ErrorCode:  "",
	}, nil
}

func GenerateRewardFromRREQ(rreq *pb.RREQRequest) int64 {
	return 0
}
