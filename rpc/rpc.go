package rpc

import (
	"fmt"
	"net"

	"github.com/m25-lab/lightning-network-node/rpc/pb"
	nodeInfoServer "github.com/m25-lab/lightning-network-node/rpc/services/node_info"

	"github.com/m25-lab/lightning-network-node/node"
	channelServer "github.com/m25-lab/lightning-network-node/rpc/services/channel"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type RPCServer struct {
	grpcServer     *grpc.Server
	serviceServers ServiceServers
}

type ServiceServers struct {
	channelServer  *channelServer.ChannelServer
	nodeInfoServer *nodeInfoServer.NodeInfoServer
}

func New(node *node.LightningNode) (*RPCServer, error) {
	var err error
	var rpcServer RPCServer

	//Init service servers
	rpcServer.serviceServers.channelServer, err = channelServer.New(node)
	if err != nil {
		return nil, err
	}

	rpcServer.serviceServers.nodeInfoServer, err = nodeInfoServer.New(node)
	if err != nil {
		return nil, err
	}

	rpcServer.grpcServer = grpc.NewServer()

	//Register service servers
	pb.RegisterChannelServiceServer(rpcServer.grpcServer, rpcServer.serviceServers.channelServer)
	pb.RegisterNodeServiceServer(rpcServer.grpcServer, rpcServer.serviceServers.nodeInfoServer)

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
