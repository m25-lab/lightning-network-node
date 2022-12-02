package main

import (
	"log"
	"net"

	"github.com/m25-lab/lightning-network-node/internal/channel"
	client "github.com/m25-lab/lightning-network-node/internal/client"
	"github.com/m25-lab/lightning-network-node/internal/pb"
	"github.com/m25-lab/lightning-network-node/internal/tx"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	//runGrpcGateway()
	client.OpenChannel()
}

func runGrpcGateway() {
	channelServer, err := channel.NewServer()
	if err != nil {
		log.Fatalf("failed to create channel server: %v", err)
	}
	txServer, err := tx.NewServer()
	if err != nil {
		log.Fatalf("failed to create commitment server: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterChannelServiceServer(grpcServer, channelServer)
	pb.RegisterTxServiceServer(grpcServer, txServer)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", "0.0.0.0:9090")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Listening on %s", listener.Addr())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}
