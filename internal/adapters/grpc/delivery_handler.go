package grpc

import (
	"context"
	"fmt"
	"io"
	"math/rand"

	"github.com/segmentio/kafka-go"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"

	kc "github.com/KidPudel/delivery-service/internal/infrastructure/kafka"
	pb "github.com/KidPudel/delivery-service/proto/delivery"
)

type DeliveryServer struct {
	pb.DeliveryServer
	kafkaClient *kc.KafkaClient
}

func NewDeliveryServer(kafkaClient *kc.KafkaClient) *DeliveryServer {
	return &DeliveryServer{
		kafkaClient: kafkaClient,
	}
}

// to which are calling
func (server *DeliveryServer) SendToDelivery(ctx context.Context, order *pb.OrderInfo) (*pb.DeliveryAcknowledgment, error) {
	server.kafkaClient.Writer.WriteMessages(
		ctx,
		kafka.Message{Value: []byte("order is accepted!")},
		kafka.Message{Value: []byte("we've already prepared the order for you")},
		kafka.Message{Value: []byte("you can track the order using you own rpc function 'StartTrackingOrder', that will trigger calling our delivery service rpc function 'FindEachOther'")},
	)
	return &pb.DeliveryAcknowledgment{
		Response: proto.String(fmt.Sprintf("order is accepted: %s", *order.Comment)),
	}, nil
}

func (server *DeliveryServer) FindEachOther(stream grpc.BidiStreamingServer[pb.Position, pb.Position]) error {
	for {
		// not in parallel, because we depend on the comming value, to then respond
		position, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		fmt.Println("client position: ", position)

		if err := stream.Send(&pb.Position{
			Lat:  proto.Int32(rand.Int31()),
			Long: proto.Int32(rand.Int31()),
		}); err != nil {
			return err
		}

	}
	return nil
}
