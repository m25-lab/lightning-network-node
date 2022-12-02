package rpc

import (
	"fmt"
	"net"

	"github.com/m25-lab/lightning-network-node/node/rpc/pb"
	"github.com/m25-lab/lightning-network-node/node/rpc/services/channel"
	"github.com/m25-lab/lightning-network-node/node/rpc/services/routing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type NodeServer struct {
	grpcServer     *grpc.Server
	serviceServers ServiceServers
}

type ServiceServers struct {
	routingServer *routing.RoutingServer
	channelServer *channel.ChannelServer
}

func New() (*NodeServer, error) {
	var err error
	var nodeServer NodeServer

	//Init service servers
	nodeServer.serviceServers.routingServer, err = routing.NewServer()
	if err != nil {
		return nil, err
	}

	nodeServer.serviceServers.channelServer, err = channel.NewServer()
	if err != nil {
		return nil, err
	}

	nodeServer.grpcServer = grpc.NewServer()
	reflection.Register(nodeServer.grpcServer)

	//Register service servers
	pb.RegisterRoutingServiceServer(nodeServer.grpcServer, nodeServer.serviceServers.routingServer)
	//pb.RegisterChannelServiceServer(nodeServer.grpcServer, nodeServer.serviceServers.channelServer)

	return &nodeServer, nil
}

func (gateway *NodeServer) Run() error {
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
