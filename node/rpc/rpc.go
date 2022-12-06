package rpc

import (
	"fmt"
	"net"

	"github.com/m25-lab/lightning-network-node/node"
	"github.com/m25-lab/lightning-network-node/node/rpc/pb"
	"github.com/m25-lab/lightning-network-node/node/rpc/service-servers/channel"
	channelServer "github.com/m25-lab/lightning-network-node/node/rpc/service-servers/channel"
	p2pserver "github.com/m25-lab/lightning-network-node/node/rpc/service-servers/p2p"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type RPCServer struct {
	grpcServer     *grpc.Server
	serviceServers ServiceServers
}

type ServiceServers struct {
	p2pServer     *p2pserver.PeerToPeerServer
	channelServer *channel.ChannelServer
}

func New(node *node.LightningNode) (*RPCServer, error) {
	var err error
	var rpcServer RPCServer

	//Init service servers
	rpcServer.serviceServers.p2pServer, err = p2pserver.New(node)
	if err != nil {
		return nil, err
	}

	rpcServer.serviceServers.channelServer, err = channelServer.New()
	if err != nil {
		return nil, err
	}

	rpcServer.grpcServer = grpc.NewServer()

	//Register service servers
	pb.RegisterPeerToPeerServer(rpcServer.grpcServer, rpcServer.serviceServers.p2pServer)
	pb.RegisterChannelServiceServer(rpcServer.grpcServer, rpcServer.serviceServers.channelServer)

	reflection.Register(rpcServer.grpcServer)

	return &rpcServer, nil
}

func (gateway *RPCServer) RunGateway() error {
	listener, err := net.Listen("tcp", "0.0.0.0:2525")
	if err != nil {
		return err
	}

	fmt.Println("Server is running in port 2525")

	err = gateway.grpcServer.Serve(listener)
	if err != nil {
		return err
	}

	return nil
}
