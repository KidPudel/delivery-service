package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	deliveryServer "github.com/KidPudel/delivery-service/internal/adapters/grpc"
	"github.com/KidPudel/delivery-service/internal/infrastructure/kafka"
	pb "github.com/KidPudel/delivery-service/proto/delivery"
)

func main() {
	listenConfig, err := net.Listen("tcp", "localhost:50052")
	if err != nil {
		log.Fatal(err)
	}

	// kafka
	kafkaClient := kafka.NewKafkaClient()

	server := grpc.NewServer()
	pb.RegisterDeliveryServer(server, deliveryServer.NewDeliveryServer(kafkaClient))
	server.Serve(listenConfig)
}
