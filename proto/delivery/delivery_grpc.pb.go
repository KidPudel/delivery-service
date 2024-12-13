// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.1
// source: proto/delivery/delivery.proto

package delivery

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Delivery_SendToDelivery_FullMethodName = "/delivery.Delivery/SendToDelivery"
	Delivery_FindEachOther_FullMethodName  = "/delivery.Delivery/FindEachOther"
)

// DeliveryClient is the client API for Delivery service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DeliveryClient interface {
	SendToDelivery(ctx context.Context, in *OrderInfo, opts ...grpc.CallOption) (*DeliveryAcknowledgment, error)
	FindEachOther(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[Position, Position], error)
}

type deliveryClient struct {
	cc grpc.ClientConnInterface
}

func NewDeliveryClient(cc grpc.ClientConnInterface) DeliveryClient {
	return &deliveryClient{cc}
}

func (c *deliveryClient) SendToDelivery(ctx context.Context, in *OrderInfo, opts ...grpc.CallOption) (*DeliveryAcknowledgment, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeliveryAcknowledgment)
	err := c.cc.Invoke(ctx, Delivery_SendToDelivery_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deliveryClient) FindEachOther(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[Position, Position], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Delivery_ServiceDesc.Streams[0], Delivery_FindEachOther_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[Position, Position]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Delivery_FindEachOtherClient = grpc.BidiStreamingClient[Position, Position]

// DeliveryServer is the server API for Delivery service.
// All implementations must embed UnimplementedDeliveryServer
// for forward compatibility.
type DeliveryServer interface {
	SendToDelivery(context.Context, *OrderInfo) (*DeliveryAcknowledgment, error)
	FindEachOther(grpc.BidiStreamingServer[Position, Position]) error
	mustEmbedUnimplementedDeliveryServer()
}

// UnimplementedDeliveryServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedDeliveryServer struct{}

func (UnimplementedDeliveryServer) SendToDelivery(context.Context, *OrderInfo) (*DeliveryAcknowledgment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendToDelivery not implemented")
}
func (UnimplementedDeliveryServer) FindEachOther(grpc.BidiStreamingServer[Position, Position]) error {
	return status.Errorf(codes.Unimplemented, "method FindEachOther not implemented")
}
func (UnimplementedDeliveryServer) mustEmbedUnimplementedDeliveryServer() {}
func (UnimplementedDeliveryServer) testEmbeddedByValue()                  {}

// UnsafeDeliveryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DeliveryServer will
// result in compilation errors.
type UnsafeDeliveryServer interface {
	mustEmbedUnimplementedDeliveryServer()
}

func RegisterDeliveryServer(s grpc.ServiceRegistrar, srv DeliveryServer) {
	// If the following call pancis, it indicates UnimplementedDeliveryServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Delivery_ServiceDesc, srv)
}

func _Delivery_SendToDelivery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeliveryServer).SendToDelivery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Delivery_SendToDelivery_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeliveryServer).SendToDelivery(ctx, req.(*OrderInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Delivery_FindEachOther_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DeliveryServer).FindEachOther(&grpc.GenericServerStream[Position, Position]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Delivery_FindEachOtherServer = grpc.BidiStreamingServer[Position, Position]

// Delivery_ServiceDesc is the grpc.ServiceDesc for Delivery service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Delivery_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "delivery.Delivery",
	HandlerType: (*DeliveryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendToDelivery",
			Handler:    _Delivery_SendToDelivery_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "FindEachOther",
			Handler:       _Delivery_FindEachOther_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/delivery/delivery.proto",
}