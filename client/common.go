package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/m25-lab/lightning-network-node/tools/crypto"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/m25-lab/lightning-network-node/core_chain_sdk/account"
	"github.com/m25-lab/lightning-network-node/database/models"
	"github.com/m25-lab/lightning-network-node/rpc/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	bankTypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

type ChannelBalanceStruct struct {
	ChannelId      string `json:"channel_id"`
	MyBalance      int64  `json:"my_balance"`
	PartnerBalance int64  `json:"partner_balance"`
}

func (client *Client) CreateConn(endpoint string) *grpc.ClientConn {
	conn, err := grpc.Dial(endpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		panic(err)
	}
	return conn
}

func (client *Client) CreateAccount(clientId string) (*account.PrivateKeySerialized, error) {
	acc := account.NewAccount()

	account, err := acc.CreateAccount()
	if err != nil {
		return nil, err
	}

	err = client.Node.Repository.Address.DeleteByClientId(context.Background(), clientId)
	if err != nil {
		return nil, err
	}

	err = client.Node.Repository.Address.InsertOne(context.Background(), &models.Address{
		ID:       primitive.NewObjectID(),
		Address:  account.AccAddress().String(),
		Pubkey:   account.PublicKey().String(),
		Mnemonic: account.Mnemonic(),
		ClientId: clientId,
	})

	if err != nil {
		return nil, err
	}

	return account, nil
}

func (client *Client) CurrentAccount(clientId string) (*account.PrivateKeySerialized, error) {
	existedAddress, err := client.Node.Repository.Address.FindByClientId(context.Background(), clientId)
	if err != nil {
		return nil, err
	}

	acc := account.NewAccount()
	account, err := acc.ImportAccount(existedAddress.Mnemonic)
	if err != nil {
		return nil, err
	}

	return account, nil
}

func (client *Client) ImportAccount(clientId string, mnemonic string) (*account.PrivateKeySerialized, error) {
	acc := account.NewAccount()
	account, err := acc.ImportAccount(mnemonic)
	if err != nil {
		return nil, err
	}

	err = client.Node.Repository.Address.DeleteByClientId(context.Background(), clientId)
	if err != nil {
		return nil, err
	}

	err = client.Node.Repository.Address.InsertOne(context.Background(), &models.Address{
		ID:       primitive.NewObjectID(),
		Address:  account.AccAddress().String(),
		Pubkey:   account.PublicKey().String(),
		Mnemonic: account.Mnemonic(),
		ClientId: clientId,
	})

	if err != nil {
		return nil, err
	}

	return account, nil
}

func (client *Client) AddWhitelist(clientId string, toAddress string) (*models.Message, error) {
	//@todo: check address is valid
	splitedAddress := strings.Split(toAddress, "@")
	if len(splitedAddress) != 2 {
		return nil, errors.New("address is invalid")
	}
	toEndpoint := splitedAddress[1]

	//@todo: create from account
	acc, err := client.CurrentAccount(clientId)
	if err != nil {
		return nil, err
	}

	//@todo: create message
	payload, err := json.Marshal(models.AddWhitelistData{
		Pubkey: acc.PublicKey().String(),
	})
	if err != nil {
		return nil, err
	}
	messageId := primitive.NewObjectID()
	message := models.Message{
		ID:         messageId,
		OriginalID: messageId,
		ChannelID:  "",
		Action:     models.AddWhitelist,
		Data:       string(payload),
		Owner:      acc.AccAddress().String(),
		Users:      []string{acc.AccAddress().String() + "@" + client.Node.Config.LNode.External, toAddress},
		IsReplied:  false,
	}
	err = client.Node.Repository.Message.InsertOne(context.Background(), &message)
	if err != nil {
		return nil, err
	}

	//@todo: send message
	rpcClient := pb.NewMessageServiceClient(client.CreateConn(toEndpoint))
	response, err := rpcClient.SendMessage(context.Background(), &pb.SendMessageRequest{
		MessageId:       message.ID.Hex(),
		ChannelID:       message.ChannelID,
		Action:          message.Action,
		Data:            string(payload),
		From:            acc.AccAddress().String() + "@" + client.Node.Config.LNode.External,
		To:              toAddress,
		ReliedMessageId: "",
	})
	if err != nil {
		return nil, err
	}
	if response.ErrorCode != "" {
		return nil, errors.New(response.ErrorCode + ": " + response.Response)
	}

	return &message, nil
}

func (client *Client) AcceptAddWhitelist(clientId string, messageId string) (*models.Message, error) {
	//@todo: create from account
	fromAccount, err := client.CurrentAccount(clientId)
	if err != nil {
		return nil, err
	}

	//@todo: check message is valid
	reliedMessage, err := client.Node.Repository.Message.FindOneById(
		context.Background(),
		fromAccount.AccAddress().String(),
		messageId,
	)
	if err != nil {
		return nil, err
	}

	var addWhitelist models.AddWhitelistData
	if err := json.Unmarshal([]byte(reliedMessage.Data), &addWhitelist); err != nil {
		return nil, errors.New("invalid data")
	}

	//@todo: create to account
	toEndpoint := strings.Split(reliedMessage.Users[0], "@")[1]
	toAccount := account.NewPKAccount(addWhitelist.Pubkey)

	//@todo create multi account
	acc := account.NewAccount()
	multiAddr, _, _ := acc.CreateMulSigAccountFromTwoAccount(fromAccount.PublicKey(), toAccount.PublicKey(), 2)

	//@todo: save whitelist
	savedWhitelist := models.Whitelist{
		ID:             primitive.NewObjectID(),
		Owner:          fromAccount.AccAddress().String(),
		PartnerAddress: reliedMessage.Users[0],
		PartnerPubkey:  addWhitelist.Pubkey,
		MultiAddress:   multiAddr,
		MultiPubkey:    "",
	}
	err = client.Node.Repository.Whitelist.InsertOne(context.Background(), &savedWhitelist)
	if err != nil {
		return nil, err
	}

	//@todo: create message
	payload, err := json.Marshal(models.AddWhitelistData{
		Pubkey: fromAccount.PublicKey().String(),
	})
	if err != nil {
		return nil, err
	}
	ID := primitive.NewObjectID()
	savedMessage := models.Message{
		ID:              ID,
		OriginalID:      ID,
		ChannelID:       "",
		Action:          models.AcceptAddWhitelist,
		Data:            string(payload),
		Owner:           fromAccount.AccAddress().String(),
		Users:           []string{fromAccount.AccAddress().String() + "@" + client.Node.Config.LNode.External, reliedMessage.Users[0]},
		ReliedMessageId: reliedMessage.ID.Hex(),
		IsReplied:       false,
	}
	err = client.Node.Repository.Message.InsertOne(context.Background(), &savedMessage)
	if err != nil {
		return nil, err
	}

	reliedMessage.IsReplied = true
	err = client.Node.Repository.Message.Update(context.Background(), reliedMessage.ID, reliedMessage)
	if err != nil {
		return nil, err
	}

	//@todo: send message
	rpcClient := pb.NewMessageServiceClient(client.CreateConn(toEndpoint))
	response, err := rpcClient.SendMessage(context.Background(), &pb.SendMessageRequest{
		MessageId:       savedMessage.ID.Hex(),
		ChannelID:       reliedMessage.ChannelID,
		Action:          models.AcceptAddWhitelist,
		Data:            string(payload),
		From:            fromAccount.AccAddress().String() + "@" + client.Node.Config.LNode.External,
		To:              reliedMessage.Users[0],
		ReliedMessageId: reliedMessage.OriginalID.Hex(),
	})
	if err != nil {
		return nil, err
	}
	if response.ErrorCode != "" {
		return nil, errors.New(response.ErrorCode + ":" + response.Response)
	}

	return &savedMessage, nil
}

func (client *Client) ResolveAddWhitelist(clientId int64, msg *models.Message) error {
	telMsg := tgbotapi.NewMessage(clientId, "")
	telMsg.ParseMode = "Markdown"

	telMsg.Text = fmt.Sprintf("👋 *New whitelist request*\n`%s` has sent you a request to add them to your whitelist. Do you want to accept?", msg.Users[0])

	telMsg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Accept", fmt.Sprintf("%s:%s", models.AcceptAddWhitelist, msg.ID.Hex())),
		),
	)
	if _, err := client.Bot.Send(telMsg); err != nil {
		return err
	}

	return nil
}

func (client *Client) ResolveAcceptAddWhitelist(clientId int64, msg *models.Message) error {
	telMsg := tgbotapi.NewMessage(clientId, "")
	telMsg.ParseMode = "Markdown"

	telMsg.Text = fmt.Sprintf("✅ *Whitelist request accepted*\n `%s` has accepted your request to add them to your whitelist.", msg.Users[0])

	if _, err := client.Bot.Send(telMsg); err != nil {
		return err
	}

	return nil
}

func (client *Client) ListWhitelist(clientId string) ([]models.Whitelist, error) {
	currentAccount, err := client.CurrentAccount(clientId)
	if err != nil {
		return nil, err
	}
	whitelists, err := client.Node.Repository.Whitelist.FindMany(context.Background(), currentAccount.AccAddress().String())
	if err != nil {
		return nil, err
	}

	return whitelists, nil
}

func (client *Client) Balance(address string) (int64, error) {
	bankRes, err := client.l1Client.bank.Balance(
		context.Background(),
		&bankTypes.QueryBalanceRequest{Address: address, Denom: "token"},
	)
	if err != nil {
		return 0, err
	}

	return bankRes.Balance.Amount.Int64(), nil
}

func (client *Client) ChannelBalance(clientId string, channelID string) (*ChannelBalanceStruct, error) {
	currentAccount, err := client.CurrentAccount(clientId)
	if err != nil {
		return nil, nil
	}

	lastestCommitment, err := client.Node.Repository.Message.FindOneByChannelIDWithAction(
		context.Background(),
		currentAccount.AccAddress().String(),
		channelID,
		models.ExchangeCommitment,
	)
	if err != nil {
		return nil, err
	}

	payload := models.CreateCommitmentData{}
	err = json.Unmarshal([]byte(lastestCommitment.Data), &payload)
	if err != nil {
		return nil, err
	}

	return &ChannelBalanceStruct{
		ChannelId:      channelID,
		MyBalance:      payload.CoinToHtlc,
		PartnerBalance: payload.CoinToCreator,
	}, nil
}

func (client *Client) RequestInvoice(clientId int64, msg *models.Message) error {
	telMsg := tgbotapi.NewMessage(clientId, "")
	telMsg.ParseMode = "Markdown"

	var invoice models.InvoiceData
	err := json.Unmarshal([]byte(msg.Data), &invoice)
	if err != nil {
		return err
	}

	telMsg.Text = fmt.Sprintf("👋 *Request invoice*\n `%s` request invoice to send you %d token. Do you want to accept ?", msg.Users[0], invoice.Amount)

	telMsg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Accept", fmt.Sprintf("%s:%s", models.AcceptRequestInvoice, msg.ID.Hex())),
		),
	)

	if _, err := client.Bot.Send(telMsg); err != nil {
		return err
	}

	return nil
}

func (client *Client) CallbackAcceptInvoice(clientId int64, msg *models.Message) error {
	telMsg := tgbotapi.NewMessage(clientId, "")
	telMsg.ParseMode = "Markdown"

	var invoice models.InvoiceData
	err := json.Unmarshal([]byte(msg.Data), &invoice)
	if err != nil {
		return err
	}

	telMsg.Text = fmt.Sprintf("✅ *Accepted request invoice*\n `%s` accept request invoice to receive you %d token with hash %s", msg.Users[0], invoice.Amount, invoice.HashSecret)

	if _, err := client.Bot.Send(telMsg); err != nil {
		return err
	}

	return nil
}

func (client *Client) AcceptRequestInvoice(clientId, messageId string) (*models.InvoiceData, error) {
	fromAccount, err := client.CurrentAccount(clientId)
	if err != nil {
		return nil, err
	}

	reliedMessage, err := client.Node.Repository.Message.FindOneByOriginalID(
		context.Background(),
		fromAccount.AccAddress().String(),
		messageId,
	)
	if err != nil {
		return nil, err
	}

	var invoice models.InvoiceData
	err = json.Unmarshal([]byte(reliedMessage.Data), &invoice)
	if err != nil {
		return nil, err
	}

	senderAddress, err := client.Node.Repository.Address.FindByAddress(
		context.Background(),
		invoice.From,
	)
	if err != nil {
		return nil, err
	}

	hashSecret, err := crypto.CryptoEngine.Hash(reliedMessage.Data)
	if err != nil {
		return nil, err
	}

	invoice.HashSecret = hashSecret
	hashedInvoiceData, err := json.Marshal(invoice)
	if err != nil {
		return nil, err
	}

	ID := primitive.NewObjectID()
	savedMessage := models.Message{
		ID:              ID,
		OriginalID:      ID,
		Action:          models.AcceptRequestInvoice,
		Data:            string(hashedInvoiceData),
		Owner:           fromAccount.AccAddress().String(),
		Users:           []string{fromAccount.AccAddress().String() + "@" + client.Node.Config.LNode.External, reliedMessage.Users[0] + "@" + client.Node.Config.LNode.External},
		ReliedMessageId: reliedMessage.ID.Hex(),
		IsReplied:       false,
	}
	err = client.Node.Repository.Message.InsertOne(context.Background(), &savedMessage)
	if err != nil {
		return nil, err
	}

	reliedMessage.IsReplied = true
	err = client.Node.Repository.Message.Update(context.Background(), reliedMessage.ID, reliedMessage)
	if err != nil {
		return nil, err
	}

	senderClientID, err := strconv.Atoi(senderAddress.ClientId)
	if err != nil {
		return nil, err
	}

	err = client.CallbackAcceptInvoice(int64(senderClientID), &savedMessage)
	if err != nil {
		return nil, err
	}

	return &invoice, nil
}
