package tx

import "github.com/m25-lab/lightning-network-node/node/rpc/pb"

type TxServer struct {
	pb.UnimplementedTxServiceServer
}

func NewServer() (*TxServer, error) {
	return &TxServer{}, nil
}
