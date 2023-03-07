package watcher

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"

	"github.com/m25-lab/lightning-network-node/config"
	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
	"github.com/tendermint/tendermint/rpc/client/http"
	"github.com/tendermint/tendermint/types"
)

const pathLogFile = "log/broadcast.log"

type Publisher struct {
	curHeight    int64
	lastestBlock int64
	writer       *kafka.Writer
	blockQue     []*types.Block
	client       *http.HTTP
	logger       *log.Logger
	runErr       error
}

func NewPublisher(config *config.Config, client *http.HTTP) (*Publisher, error) {
	height, err := loadHistory(pathLogFile)
	if err != nil {
		return nil, err
	}

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: config.Kafka.Brokers,
		Topic:   "blocks",
	})

	logger, err := createFileLogger(pathLogFile)
	if err != nil {
		return nil, err
	}

	status, err := client.Status(context.Background())
	if err != nil {
		return nil, err
	}

	publisher := &Publisher{
		height,
		status.SyncInfo.LatestBlockHeight,
		writer,
		[]*types.Block{},
		client,
		logger,
		errors.New("Publisher is not running. Use 'go .Start()' method to start"),
	}

	return publisher, nil
}

func (p *Publisher) Start() {
	p.runErr = nil

	for {
		height := p.pickBlock()
		if height == -1 {
			log.Debug("Reach latest")
			continue
		}

		block, err := p.getBlock(height)
		if err != nil {
			p.runErr = err
			return
		}

		if err := p.writeBlock(block); err != nil {
			p.runErr = err
			return
		}

		p.curHeight = height
	}
}

func (p *Publisher) PushBlock(block *types.Block) error {
	if p.runErr != nil {
		return p.runErr
	}

	p.blockQue = append(p.blockQue, block)
	p.lastestBlock = block.Height
	return nil
}

func (p *Publisher) popBlock() *types.Block {
	if len(p.blockQue) == 0 {
		return nil
	}

	block := p.blockQue[0]
	p.blockQue = p.blockQue[1:]

	return block
}

func (p *Publisher) getBlock(height int64) (*types.Block, error) {
	// TODO: Change to use method support to get multiple block
	if len(p.blockQue) > 0 && p.blockQue[0].Height == height {
		return p.popBlock(), nil
	}

	res, err := p.client.Block(context.Background(), &height)
	if err != nil {
		return nil, err
	}

	return res.Block, nil
}

func (p *Publisher) pickBlock() int64 {
	if p.curHeight == -1 {
		return p.lastestBlock
	} else {
		if p.curHeight < p.lastestBlock {
			return p.curHeight + 1
		} else {
			return -1
		}
	}
}

func (p *Publisher) writeBlock(block *types.Block) error {
	bytes, err := json.Marshal(block)
	if err != nil {
		return err
	}

	err = p.writer.WriteMessages(context.Background(), kafka.Message{
		Key:   []byte(strconv.FormatInt(block.Height, 10)),
		Value: bytes,
	})

	if err != nil {
		return err
	}

	p.logger.Info("Broadcast block: ", block.Height)
	log.Info("Broadcast block: ", block.Height)

	return nil
}
