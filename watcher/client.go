package watcher

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/tendermint/tendermint/rpc/client/http"
)

const (
	LocalHost = "http://0.0.0.0:26657"
)

func getClient(endpoint string) (*http.HTTP, error) {
	client, err := client.NewClientFromNode(endpoint)
	if err != nil {
		return nil, err
	}

	return client, nil
}
