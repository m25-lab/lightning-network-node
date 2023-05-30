package routing

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/m25-lab/lightning-network-node/core_chain_sdk/account"
	"github.com/m25-lab/lightning-network-node/core_chain_sdk/channel"
	"github.com/m25-lab/lightning-network-node/core_chain_sdk/common"
	"github.com/m25-lab/lightning-network-node/database/models"
	"github.com/m25-lab/lightning-network-node/rpc/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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
		fmt.Println("RREQ")
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
		fmt.Println("RREP")
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

func (server *RoutingServer) ProcessInvoiceSecret(ctx context.Context, req *pb.InvoiceSecretMessage) (*pb.RoutingBaseResponse, error) {
	//check hash
	receiverCommit, err := server.ValidateInvoiceSecret(ctx, req)
	if err != nil {
		return &pb.RoutingBaseResponse{
			ErrorCode: pb.RoutingErrorCode_VALIDATE_INVOICE_SECRET,
		}, fmt.Errorf("Validate invoice secret error")
	}
	//luu DB
	data := models.FwdSecret{
		HashcodeDest: req.Hashcode,
		Secret:       req.Secret,
	}
	if err := server.Node.Repository.FwdSecret.InsertSecret(ctx, &data); err != nil {
		return &pb.RoutingBaseResponse{
			ErrorCode: pb.RoutingErrorCode_INSERT_SECRET,
		}, fmt.Errorf("Insert secret error")
	}

	split := strings.Split(req.Dest, "@")
	destAddr := split[0] // A
	toEndpoint := split[1]

	activeAddress, err := server.Node.Repository.Address.FindByAddress(ctx, destAddr) //check coi A co trong db minh khong
	if err != nil {
		if err == mongo.ErrNoDocuments {
			go func() {
				nextEntry, err := server.Node.Repository.Routing.FindByDestAndBroadcastId(ctx, destAddr, req.Dest, req.Hashcode)
				if err != nil {
					println("nextEntry_FindByDestAndBroadcastId", err.Error())
				}
				rpcClient := pb.NewRoutingServiceClient(server.Client.CreateConn(toEndpoint))
				_, err = rpcClient.ProcessInvoiceSecret(ctx, &pb.InvoiceSecretMessage{
					Hashcode: req.Hashcode,
					Secret:   req.Secret,
					Dest:     nextEntry.DestinationAddress,
				})
				if err != nil {
					println("ProcessInvoiceSecret", err.Error())
				}
			}()
			return &pb.RoutingBaseResponse{
				ErrorCode: pb.RoutingErrorCode_OK,
			}, nil
		} else {
			return &pb.RoutingBaseResponse{
				ErrorCode: pb.RoutingErrorCode_DESTINATION_ADDRESS_FIND_BY_ADDRESS,
			}, fmt.Errorf("Find destination address by address error")
		}
	}
	// is dest -> phase commitment
	go func() {
		amount := receiverCommit.CoinTransfer
		nexthops, err := server.Node.Repository.Routing.FindRouting(ctx, models.Routing{
			BroadcastID: req.Hashcode,
		})
		if err != nil {
			println("FindRouting", err.Error())
			return
		}
		dest := nexthops[0].NextHop

		err = server.Client.LnTransfer(activeAddress.ClientId, receiverCommit.From, amount, &dest, &receiverCommit.HashcodeDest)
		if err != nil {
			println("Trade commitment - LnTransfer:", err.Error())
		}
	}()

	return &pb.RoutingBaseResponse{
		ErrorCode: pb.RoutingErrorCode_OK,
	}, nil
}

func (server *RoutingServer) RequestInvoice(ctx context.Context, req *pb.IREQMessage) (*pb.IREPMessage, error) {
	//TODO: check to address is active

	secret, err := common.RandomSecret()
	if err != nil {
		println("RandomSecret:", err.Error())
		return &pb.IREPMessage{
			ErrorCode: err.Error(),
		}, nil
	}
	hashcode := common.ToHashCode(secret)
	server.Node.Repository.Invoice.InsertInvoice(ctx, &models.InvoiceData{
		Amount: req.Amount,
		From:   req.From,
		To:     req.To,
		Hash:   hashcode,
		Secret: secret,
	})
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
		}, nil
	}
	toAccount, _ := account.NewAccount().ImportAccount(existToAddress.Mnemonic)

	//get "From" public key
	fromAddressFromDB, err := server.Client.Node.Repository.Whitelist.FindOneByPartnerAddress(context.Background(), toAddress, req.From)
	if err != nil {
		return &pb.FwdMessageResponse{
			Response:  err.Error(),
			ErrorCode: "1004",
		}, nil
	}
	fromAccount := account.NewPKAccount(fromAddressFromDB.PartnerPubkey)

	//gen multiAddr
	multisigAddr, multiSigPubkey, _ := account.NewAccount().CreateMulSigAccountFromTwoAccount(fromAccount.PublicKey(), toAccount.PublicKey(), 2)

	var myCommitmentPayload models.SenderCommitment
	if err := json.Unmarshal([]byte(req.Data), &myCommitmentPayload); err != nil {
		return &pb.FwdMessageResponse{
			Response:  err.Error(),
			ErrorCode: "1006",
		}, nil
	}

	//check hash code htlc
	exchangeHashcodeMessage, err := server.Node.Repository.Message.FindOneByChannelID(context.Background(), toAccount.AccAddress().String(), multisigAddr+":token:1")
	if err != nil {
		return &pb.FwdMessageResponse{
			Response:  err.Error(),
			ErrorCode: "1006",
		}, nil
	}
	if exchangeHashcodeMessage.Action != models.ExchangeHashcode {
		return &pb.FwdMessageResponse{
			Response:  "partner has not sent hashcode yet",
			ErrorCode: "1006",
		}, nil
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
		}, nil
	}

	channelClient := channel.NewChannel(*server.Client.ClientCtx)

	//build SenderCommit and sign
	senderCMsg := channelClient.CreateSenderCommitmentMsg(
		multisigAddr,
		fromAccount.AccAddress().String(),
		myCommitmentPayload.CoinToSender,
		myCommitmentPayload.CoinToHTLC,
		myCommitmentPayload.CoinTransfer,
		myCommitmentPayload.HashcodeHTLC,
		myCommitmentPayload.HashcodeDest,
	)

	signSenderCommitmentMsg := channel.SignMsgRequest{
		Msg:      senderCMsg,
		GasLimit: 200000,
		GasPrice: "0token", //TODO: 0token or 0stake
	}

	//sign sender commitment in advance
	strSigSender, err := channelClient.SignMultisigTxFromOneAccount(signSenderCommitmentMsg, toAccount, multiSigPubkey)
	if err != nil {
		return &pb.FwdMessageResponse{
			Response:  err.Error(),
			ErrorCode: "1006",
		}, nil
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

	//Build and sign receiver commit
	receiverCMsg := channelClient.CreateReceiverCommitmentMsg(
		multisigAddr,
		toAddress,
		myCommitmentPayload.CoinToHTLC,
		myCommitmentPayload.CoinToSender,
		myCommitmentPayload.CoinTransfer,
		myCommitmentPayload.HashcodeHTLC,
		myCommitmentPayload.HashcodeDest,
	)

	signReceiverCommitmentMsg := channel.SignMsgRequest{
		Msg:      receiverCMsg,
		GasLimit: 200000,
		GasPrice: "0token", //TODO: 0token or 0stake
	}

	strSigReceiver, err := channelClient.SignMultisigTxFromOneAccount(signReceiverCommitmentMsg, toAccount, multiSigPubkey)
	if err != nil {
		return &pb.FwdMessageResponse{
			Response:  err.Error(),
			ErrorCode: "1006",
		}, nil
	}

	partnerCommitmentPayload, err := json.Marshal(models.ReceiverCommitment{
		Creator:        receiverCMsg.Creator,
		From:           receiverCMsg.From,
		ChannelID:      receiverCMsg.Channelid,
		CoinToReceiver: receiverCMsg.Cointoreceiver.Amount.Int64(),
		CoinToHTLC:     receiverCMsg.Cointohtlc.Amount.Int64(),
		HashcodeHTLC:   receiverCMsg.Hashcodehtlc,
		TimelockHTLC:   receiverCMsg.Timelockhtlc,
		CoinTransfer:   receiverCMsg.Cointransfer.Amount.Int64(),
		HashcodeDest:   receiverCMsg.Hashcodedest,
		TimelockSender: receiverCMsg.Timelocksender,
		Multisig:       receiverCMsg.Multisig,
	})

	if err != nil {
		return &pb.FwdMessageResponse{
			Response:  err.Error(),
			ErrorCode: "1006",
		}, nil
	}
	//find invoice in db, exist => is Dest

	needNext := false
	_, err = server.Node.Repository.Invoice.FindByHash(ctx, req.HashcodeDest)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			needNext = true
		} else {
			return &pb.FwdMessageResponse{
				Response:  err.Error(),
				ErrorCode: "checkIsDest",
			}, nil
		}
	}
	if needNext {
		//find next hop and reuse
		go func() {
			nextHop, err := server.Node.Repository.Routing.FindByDestAndBroadcastId(ctx, toAddress, req.Dest, req.HashcodeDest)
			if err != nil {
				println("Missing routing entry for:", req.Dest)
				return
			}
			err = server.Client.LnTransferMulti(existToAddress.ClientId, nextHop.NextHop, myCommitmentPayload.CoinTransfer, &req.Dest, &req.HashcodeDest)
			if err != nil {
				println("Trade fwd commitment - LnTransferMulti:", err.Error())
			}

		}()
	} else {
		//TODO: to phase reveal C's secret, call processInvoiceSecret to B
	}

	return &pb.FwdMessageResponse{
		Response:   string(partnerCommitmentPayload),
		PartnerSig: strSigReceiver,
		ErrorCode:  "",
	}, nil
}
