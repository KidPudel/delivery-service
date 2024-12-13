package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	deliveryServer "github.com/KidPudel/delivery-service/internal/adapters/grpc"
	pb "github.com/KidPudel/delivery-service/proto/delivery"
)

func main() {
	listenConfig, err := net.Listen("tcp", "localhost:50052")
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	pb.RegisterDeliveryServer(server, deliveryServer.NewDeliveryServer())
	server.Serve(listenConfig)
}
