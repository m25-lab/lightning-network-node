package core

import (
	"github.com/m25-lab/lightning-network-node/node/p2p/peer"
	"github.com/m25-lab/lightning-network-node/node/rpc/pb"
)

// A channel will be connect two node-info with one of them is our node-info.
type Channel struct {
	Address    Address
	Peers      [2]peer.Peer
	Amount     [2]uint64
	Sequence   uint32
	Signatures [2]string
	IsOpen     bool
}

type ListChannel []Channel

func (listChannel *ListChannel) Proto() ([]*pb.Channel, error) {
	listChannelProto := make([]*pb.Channel, len(*listChannel))

	for index, channel := range *listChannel {
		peers := peer.ListPeer(channel.Peers[:])
		peerProtos, err := peers.Proto()
		if err != nil {
			return nil, err
		}

		listChannelProto[index] = &pb.Channel{
			Address:    string(channel.Address[:]),
			Peers:      peerProtos,
			Amount:     channel.Amount[:],
			Sequence:   channel.Sequence,
			Signatures: channel.Signatures[:],
			IsOpen:     channel.IsOpen,
		}
	}

	return listChannelProto, nil
}
