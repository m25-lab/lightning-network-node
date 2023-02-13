package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/m25-lab/lightning-network-node/core_chain_sdk/account"
	"github.com/m25-lab/lightning-network-node/database/models"
	"github.com/m25-lab/lightning-network-node/rpc/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

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
	data, err := json.Marshal(models.AddWhitelistData{
		Publickey: acc.PublicKey().String(),
	})
	if err != nil {
		return nil, err
	}
	message := models.Message{
		ID:        primitive.NewObjectID(),
		ChannelID: "",
		Action:    models.AddWhitelist,
		Data:      string(data),
		Users:     []string{acc.AccAddress().String() + "@" + client.Node.Config.LNode.External, toAddress},
	}
	err = client.Node.Repository.Message.InsertOne(context.Background(), &message)
	if err != nil {
		return nil, err
	}

	//@todo: send message
	rpcClient := pb.NewMessageServiceClient(client.CreateConn(toEndpoint))
	if err != nil {
		return nil, err
	}
	response, err := rpcClient.SendMessage(context.Background(), &pb.SendMessageRequest{
		MessageId:       message.ID.Hex(),
		ChannelId:       message.ChannelID,
		Action:          message.Action,
		Data:            string(data),
		From:            acc.AccAddress().String() + "@" + client.Node.Config.LNode.External,
		To:              toAddress,
		AcceptMessageId: "",
	})
	if err != nil {
		return nil, err
	}
	if response.ErrorCode != "" {
		return nil, errors.New(response.ErrorCode)
	}

	return &message, nil
}

func (client *Client) AcceptAddWhitelist(clientId string, messageId string) (*models.Message, error) {
	//@todo: check message is valid
	message, err := client.Node.Repository.Message.FindOneById(context.Background(), messageId)
	if err != nil {
		return nil, err
	}
	var addWhitelist models.AddWhitelistData
	if err := json.Unmarshal([]byte(message.Data), &addWhitelist); err != nil {
		return nil, errors.New("invalid data")
	}

	//@todo: create from account
	fromAccount, err := client.CurrentAccount(clientId)
	if err != nil {
		return nil, err
	}

	//@todo: create to account
	toEndpoint := strings.Split(message.Users[0], "@")[1]
	toAccount := account.NewPKAccount(addWhitelist.Publickey)

	//@todo: create message
	savedMessage := models.Message{
		ID:        primitive.NewObjectID(),
		ChannelID: "",
		Action:    models.AcceptAddWhitelist,
		Data:      fromAccount.PublicKey().String(),
		Users:     []string{fromAccount.AccAddress().String() + "@" + client.Node.Config.LNode.External, message.Users[0]},
	}
	err = client.Node.Repository.Message.InsertOne(context.Background(), &savedMessage)
	if err != nil {
		return nil, err
	}

	//@todo create multi account
	acc := account.NewAccount()
	multiAddr, _, _ := acc.CreateMulSigAccountFromTwoAccount(fromAccount.PublicKey(), toAccount.PublicKey(), 2)

	//@todo: save whitelist
	savedWhitelist := models.Whitelist{
		ID: primitive.NewObjectID(),
		Users: []string{
			fromAccount.AccAddress().String() + "@" + client.Node.Config.LNode.External, message.Users[0],
		},
		Pubkeys: []string{
			fromAccount.PublicKey().String(), addWhitelist.Publickey,
		},
		MultiAddress: multiAddr,
		MultiPubkey:  "",
	}
	err = client.Node.Repository.Whitelist.InsertOne(context.Background(), &savedWhitelist)
	if err != nil {
		return nil, err
	}

	//@todo: send message
	data, err := json.Marshal(models.AddWhitelistData{
		Publickey: fromAccount.PublicKey().String(),
	})
	if err != nil {
		return nil, err
	}
	rpcClient := pb.NewMessageServiceClient(client.CreateConn(toEndpoint))
	response, err := rpcClient.SendMessage(context.Background(), &pb.SendMessageRequest{
		MessageId:       savedMessage.ID.Hex(),
		ChannelId:       message.ChannelID,
		Action:          models.AcceptAddWhitelist,
		Data:            string(data),
		From:            fromAccount.AccAddress().String() + "@" + client.Node.Config.LNode.External,
		To:              message.Users[0],
		AcceptMessageId: message.ID.Hex(),
	})
	if err != nil {
		return nil, err
	}
	if response.ErrorCode != "" {
		return nil, errors.New(response.ErrorCode)
	}

	return &savedMessage, nil
}

func (client *Client) ParseCallbackData(data string) (string, string, error) {
	splitedData := strings.Split(data, ":")
	if len(splitedData) != 2 {
		return "", "", errors.New("invalid data")
	}

	return splitedData[0], splitedData[1], nil
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

	telMsg.Text = fmt.Sprintf("✅ *Whitelist request accepted*\n`%s` has accepted your request to add them to your whitelist.", msg.Users[0])

	if _, err := client.Bot.Send(telMsg); err != nil {
		return err
	}

	return nil
}
