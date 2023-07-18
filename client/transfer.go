package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/m25-lab/lightning-network-node/rpc/pb"

	channeltypes "github.com/m25-lab/channel/x/channel/types"
	"github.com/m25-lab/lightning-network-node/core_chain_sdk/account"
	"github.com/m25-lab/lightning-network-node/core_chain_sdk/bank"
	"github.com/m25-lab/lightning-network-node/core_chain_sdk/common"
	"github.com/m25-lab/lightning-network-node/database/models"
)

type AccountPacked struct {
	fromAccount *account.PrivateKeySerialized
	toAccount   *account.PKAccount
	toEndpoint  string
}

type LnTransferRes struct {
	ChannelID      string
	CommitmentID   string
	OwnBalance     int64
	PartnerBalance int64
}

func (client *Client) LnTransfer(
	clientId string,
	to string,
	amount int64,
	fwdDest *string,
	hashcodeDest *string,
) (*LnTransferRes, error) {
	//create account packed
	fromAccount, err := client.CurrentAccount(clientId)
	if err != nil {
		return nil, err
	}
	existedWhitelist, err := client.Node.Repository.Whitelist.FindOneByPartnerAddress(context.Background(), fromAccount.AccAddress().String(), to)
	if err != nil {
		fmt.Println("fromAccount.AccAddress().String() ", fromAccount.AccAddress().String())
		fmt.Println("to ", to)
		fmt.Println("FindOneByPartnerAddress...")
		fmt.Println("err ", err.Error())
		return nil, err
	}
	toAccount := account.NewPKAccount(existedWhitelist.PartnerPubkey)
	toEndpoint := strings.Split(to, "@")[1]
	accountPacked := &AccountPacked{
		fromAccount: fromAccount,
		toAccount:   toAccount,
		toEndpoint:  toEndpoint,
	}

	//check multisigAddr active
	multisigAddr, _, _ := account.NewAccount().CreateMulSigAccountFromTwoAccount(accountPacked.fromAccount.PublicKey(), accountPacked.toAccount.PublicKey(), 2)
	multisigAddrBalance, err := client.Balance(multisigAddr)
	if err != nil {
		return nil, err
	}
	if multisigAddrBalance == 0 {
		err = client.Transfer(clientId, multisigAddr, 1)
		if err != nil {
			return nil, err
		}
	}

	fromAmount := int64(0)
	toAmount := amount
	hashcodePayload := models.ExchangeHashcodeData{}

	//check channel open
	isOpenChannel := true
	_, err = client.l1Client.channel.Channel(
		context.Background(),
		&channeltypes.QueryGetChannelRequest{
			Index: multisigAddr + ":token:1",
		},
	)
	if err != nil && err.Error() == "rpc error: code = NotFound desc = not found" {
		isOpenChannel = false
	}
	println("isOpenChannel :", isOpenChannel)
	if !isOpenChannel {
		fromBalance, err := client.Balance(fromAccount.AccAddress().String())
		if err != nil {
			return nil, err
		}
		if fromBalance < amount {
			return nil, fmt.Errorf("not enough balance")
		}
	} else {
		lastestCommitment, err := client.Node.Repository.Message.FindOneByChannelIDWithAction(
			context.Background(),
			fromAccount.AccAddress().String(),
			multisigAddr+":token:1",
			models.ExchangeCommitment,
		)
		if err != nil {
			fmt.Println("FindOneByChannelIDWithAction... ", err.Error())
			return nil, err
		}

		payload := models.CreateCommitmentData{}
		err = json.Unmarshal([]byte(lastestCommitment.Data), &payload)
		if err != nil {
			return nil, err
		}

		fromAmount = payload.CoinToHtlc - amount
		if fromAmount < 0 {
			return nil, fmt.Errorf("not enough balance in channel")
		}
		toAmount = payload.CoinToCreator + amount

		//get last exchangehashcode to reveal
		latestHashCode, err := client.Node.Repository.Message.FindOneByChannelIDWithAction(
			context.Background(),
			fromAccount.AccAddress().String(),
			multisigAddr+":token:1",
			models.ExchangeHashcode,
		)
		if err != nil {
			fmt.Println("FindOneByChannelIDWithAction...")
			return nil, err
		}
		err = json.Unmarshal([]byte(latestHashCode.Data), &hashcodePayload)
		if err != nil {
			return nil, err
		}
	}

	//exchange hashcode
	_, err = client.ExchangeHashcode(clientId, accountPacked)
	if err != nil {
		return nil, err
	}

	savedMesssage, err := client.ExchangeCommitment(clientId, accountPacked, fromAmount, toAmount, fwdDest, hashcodeDest, !isOpenChannel)
	if err != nil {
		return nil, err
	}

	//open channel
	if !isOpenChannel {
		err = client.OpenChannel(clientId, accountPacked)
		if err != nil {
			return nil, err
		}
	} else {
		_, err = client.ExchangeSecret(clientId, accountPacked, hashcodePayload)
		if err != nil {
			return nil, err
		}
	}
	if hashcodeDest != nil {
		msg := fmt.Sprintf("⚡ *Transfer successfully.* \n Transfer `%d` to `%s` \n "+
			"Your balance: `%d` \n Partner balance: `%d` \n"+
			"CommitmentID: `%s`", amount, to, fromAmount, toAmount, savedMesssage.ID.Hex())
		client.SendTele(clientId, msg)
	}

	return &LnTransferRes{
		ChannelID:      multisigAddr + ":token:1",
		CommitmentID:   savedMesssage.ID.Hex(),
		OwnBalance:     fromAmount,
		PartnerBalance: toAmount,
	}, nil
}

func (client *Client) Transfer(clientId string, toAddress string, value int64) error {
	if strings.Contains(toAddress, "@") {
		parsedAddress := strings.Split(toAddress, "@")
		toAddress = parsedAddress[0]
	}
	fmt.Print("toAddress: ", toAddress, "\n")

	fromAccount, err := client.CurrentAccount(clientId)
	if err != nil {
		return err
	}

	bankClient := bank.NewBank(*client.ClientCtx, "token", 60)
	request := &bank.TransferRequest{
		PrivateKey: fromAccount.PrivateKeyToString(),
		Receiver:   toAddress,
		Amount:     big.NewInt(value),
		GasLimit:   100000,
		GasPrice:   "0token",
	}

	txBuilder, err := bankClient.TransferRawDataWithPrivateKey(request)
	if err != nil {
		return err
	}

	txJson, err := common.TxBuilderJsonEncoder(client.ClientCtx.TxConfig, txBuilder)
	if err != nil {
		return err
	}

	txByte, err := common.TxBuilderJsonDecoder(client.ClientCtx.TxConfig, txJson)
	if err != nil {
		return err
	}

	broadcastResponse, err := client.ClientCtx.BroadcastTxCommit(txByte)
	if err != nil {
		return err
	}
	fmt.Printf("Transfer: %s\n", broadcastResponse.String())

	return nil
}

func (client *Client) LnTransferMulti(
	clientId string,
	to string,
	amount int64,
	hashcodeDest *string,
	isSkipGetInvoice bool,
	hops int64,
) error {
	//request invoice
	fromAccount, err := client.CurrentAccount(clientId)
	if err != nil {
		return err
	}
	selfAddress := fromAccount.AccAddress().String() + "@" + client.Node.Config.LNode.External

	var invoiceResponse *pb.IREPMessage
	if !isSkipGetInvoice {
		invoiceResponse, err = client.GetInvoice(fromAccount, amount, to)
		if err != nil {
			return err
		}
		hashcodeDest = &invoiceResponse.Hash
	}

	// try get next hop

	log.Println("selfAddress", selfAddress)
	log.Println("to", to)
	if hashcodeDest == nil {
		return fmt.Errorf("Nil")
	} else {
		log.Println("hash ", *hashcodeDest)
	}
	nextHop, err := client.Node.Repository.Routing.FindByDestAndBroadcastId(context.Background(), selfAddress, to, *hashcodeDest)
	if err != nil {
		go client.StartRouting(*hashcodeDest, amount, selfAddress, to)
		msg := "Hệ thống đang định tuyến"
		err := client.SendTele(clientId, msg)
		if err != nil {
			log.Println("In nextHop - SendTele :", err.Error())
		}
		return fmt.Errorf("routing...")
	}

	nextHopSplit := strings.Split(nextHop.NextHop, "@")
	existedWhitelist, err := client.Node.Repository.Whitelist.FindOneByPartnerAddress(context.Background(), fromAccount.AccAddress().String(), nextHop.NextHop)
	if err != nil {
		log.Println("FindOneByPartnerAddress...")
		return err
	}

	toAccount := account.NewPKAccount(existedWhitelist.PartnerPubkey)
	toEndpoint := nextHopSplit[1]

	//??
	accountPacked := &AccountPacked{
		fromAccount: fromAccount,
		toAccount:   toAccount,
		toEndpoint:  toEndpoint,
	}

	//check multisigAddr active
	multisigAddr, _, _ := account.NewAccount().CreateMulSigAccountFromTwoAccount(accountPacked.fromAccount.PublicKey(), accountPacked.toAccount.PublicKey(), 2)
	multisigAddrBalance, err := client.Balance(multisigAddr)
	if err != nil {
		log.Println("Balance...")
		return err
	}
	if multisigAddrBalance < amount {
		//err = client.Transfer(clientId, multisigAddr, 1)
		if err != nil {
			return errors.New("Multisig not enough balance:" + strconv.FormatInt(multisigAddrBalance, 10))
		}
	}

	fromAmount := int64(0)
	toAmount := amount

	//check channel open
	_, err = client.l1Client.channel.Channel(
		context.Background(),
		&channeltypes.QueryGetChannelRequest{
			Index: multisigAddr + ":token:1",
		},
	)
	if err != nil && err.Error() == "rpc error: code = NotFound desc = not found" {
		return errors.New("missing chanel with: " + nextHop.NextHop)
	}

	lastestCommitment, err := client.Node.Repository.Message.FindOneByChannelIDWithAction(
		context.Background(),
		fromAccount.AccAddress().String(),
		multisigAddr+":token:1",
		models.ExchangeCommitment,
	)
	if err != nil {
		log.Println("FindOneByChannelIDWithAction...")
		return err
	}
	if lastestCommitment.IsReplied {
		return errors.New("channel with " + nextHop.NextHop + " broadcasted.")
	}
	payload := models.CreateCommitmentData{}
	err = json.Unmarshal([]byte(lastestCommitment.Data), &payload)
	if err != nil {
		log.Println("CreateCommitmentData...")
		return err
	}

	fromAmount = payload.CoinToHtlc - amount
	if fromAmount < 0 {
		return fmt.Errorf("not enough balance in channel")
	}
	toAmount = payload.CoinToCreator

	//exchange hashcode
	_, err = client.ExchangeHashcode(clientId, accountPacked)
	if err != nil {
		log.Println("ExchangeHashcode...")
		return err
	}

	if hops == 0 {
		hops = nextHop.HopCounter
	}
	_, err = client.ExchangeFwdCommitment(clientId, accountPacked, fromAmount, toAmount, amount, to, hashcodeDest, hops)
	if err != nil {
		log.Println("ExchangeFwdCommitment...", err.Error())
		return err
	}

	return nil
}

func (client *Client) GetInvoice(fromAccount *account.PrivateKeySerialized, amount int64, to string) (*pb.IREPMessage, error) {
	rpcClient := pb.NewRoutingServiceClient(client.CreateConn(strings.Split(to, "@")[1]))
	selfAddress := fromAccount.AccAddress().String() + "@" + client.Node.Config.LNode.External
	invoiceResponse, err := rpcClient.RequestInvoice(context.Background(), &pb.IREQMessage{
		Amount: amount,
		From:   selfAddress,
		To:     to,
	})
	if err != nil {
		return nil, err
	}
	if invoiceResponse.ErrorCode != "" {
		return nil, errors.New(invoiceResponse.ErrorCode)
	}

	// save invoice
	err = client.Node.Repository.Invoice.InsertInvoice(context.Background(), &models.InvoiceData{
		Amount: amount,
		From:   selfAddress,
		To:     to,
		Hash:   invoiceResponse.Hash,
	})
	if err != nil {
		return &pb.IREPMessage{
			ErrorCode: err.Error(),
		}, nil
	}

	return invoiceResponse, nil
}

func (client *Client) StartRouting(invoiceHash string, amount int64, selfAddress, destAddress string) error {
	log.Println("StartRouting... run")
	rreqData := models.RREQData{
		Amount:         amount,
		HopCounter:     -1,
		RemainReward:   0,
		SequenceNumber: time.Now().Unix(),
	}
	rreqDataByte, _ := json.Marshal(rreqData)
	rpcClient := pb.NewRoutingServiceClient(client.CreateConn(client.Node.Config.LNode.External))
	res, err := rpcClient.RREQ(context.Background(), &pb.RREQRequest{
		BroadcastID:        invoiceHash,
		DestinationAddress: destAddress,
		SourceAddress:      selfAddress,
		// FromAddress:
		ToAddress: selfAddress,
		Data:      string(rreqDataByte),
	})
	if err != nil {
		log.Println("StartRouting err: ", err.Error())
	} else {
		log.Println("StartRouting res: ", res.ErrorCode)
	}
	return nil
}
