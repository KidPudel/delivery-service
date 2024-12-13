package grpc

import (
	"context"
	"fmt"
	"io"
	"math/rand"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"

	pb "github.com/KidPudel/delivery-service/proto/delivery"
)

type DeliveryServer struct {
	pb.DeliveryServer
}

func NewDeliveryServer() *DeliveryServer {
	return &DeliveryServer{}
}

// to which are calling
func (deliveryserver *DeliveryServer) SendToDelivery(_ context.Context, order *pb.OrderInfo) (*pb.DeliveryAcknowledgment, error) {
	return &pb.DeliveryAcknowledgment{
		Response: proto.String(fmt.Sprintf("order is accepted: %s", *order.Comment)),
	}, nil
}

func (deliveryserver *DeliveryServer) FindEachOther(stream grpc.BidiStreamingServer[pb.Position, pb.Position]) error {
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
