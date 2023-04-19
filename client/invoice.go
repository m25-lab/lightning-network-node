package client

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/m25-lab/lightning-network-node/database/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (client *Client) CreateInvoice(clientId string, lastReceiverAddressStr string, amount int64) (*models.InvoiceData, error) {
	firstSenderAddress, err := client.Node.Repository.Address.FindByClientId(context.Background(), clientId)
	if err != nil {
		return nil, err
	}

	lastReceiverAddress, err := client.Node.Repository.Address.FindByAddress(context.Background(), lastReceiverAddressStr)
	if err != nil {
		return nil, err
	}

	lastReceiverClientID, err := strconv.ParseInt(lastReceiverAddress.ClientId, 10, 64)
	if err != nil {
		return nil, err
	}

	invoice := models.InvoiceData{
		Amount: amount,
		From:   firstSenderAddress.Address,
		To:     lastReceiverAddressStr,
	}

	invoiceData, _ := json.Marshal(invoice)

	ID := primitive.NewObjectID()
	receiverMsg := &models.Message{
		ID:         ID,
		OriginalID: ID,
		Action:     models.RequestInvoice,
		Owner:      lastReceiverAddress.Address,
		Data:       string(invoiceData),
		Users:      []string{firstSenderAddress.Address + "@" + client.Node.Config.LNode.External, lastReceiverAddressStr + "@" + client.Node.Config.LNode.External},
		IsReplied:  false,
	}
	err = client.RequestInvoice(lastReceiverClientID, receiverMsg)
	if err != nil {
		return nil, err
	}

	err = client.Node.Repository.Message.InsertOne(context.Background(), receiverMsg)
	if err != nil {
		return nil, err
	}

	return &invoice, nil
}
