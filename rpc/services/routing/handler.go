package routing

import (
	"context"
	"encoding/json"
	"github.com/m25-lab/lightning-network-node/core_chain_sdk/account"
	"github.com/m25-lab/lightning-network-node/core_chain_sdk/channel"
	"github.com/m25-lab/lightning-network-node/core_chain_sdk/common"
	"github.com/m25-lab/lightning-network-node/database/models"
	"github.com/m25-lab/lightning-network-node/rpc/pb"
	"go.mongodb.org/mongo-driver/mongo"
	"strings"
)

func (server *RoutingServer) ProcessInvoiceSecret(ctx context.Context, req *pb.InvoiceSecretMessage) (*pb.RoutingResponse, error) {
	//check hash
	receiverCommit, err := server.ValidateInvoiceSecret(ctx, req)
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

	activeAddress, err := server.Node.Repository.Address.FindByAddress(ctx, destAddr)
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
	// is dest -> phase commitment
	go func() {
		amount := receiverCommit.CoinTransfer
		dest, err := server.Node.Repository.RoutingEntry.FindDestByHash(ctx, req.Hashcode)
		if err != nil {
			println("FindDestByHash", err.Error())
			return
		}
		err = server.Client.LnTransfer(activeAddress.ClientId, receiverCommit.From, amount, dest, &receiverCommit.HashcodeDest)
		if err != nil {
			println("Trade commitment - LnTransfer:", err.Error())
		}
	}()

	return &pb.RoutingResponse{
		Response:  "success",
		ErrorCode: "",
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
	//TODO: Check Hashcode Dest in DB to know if self is DEST -> to Reveal secret
	//TODO: Recheck mentioned phase because DB use FWDMessage
	return &pb.FwdMessageResponse{
		Response:   string(partnerCommitmentPayload),
		PartnerSig: strSigReceiver,
		ErrorCode:  "",
	}, nil
}
