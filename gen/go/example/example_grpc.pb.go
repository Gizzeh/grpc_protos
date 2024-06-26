// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.6.1
// source: example/example.proto

package example

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ExampleService_Greeting_FullMethodName    = "/example.ExampleService/Greeting"
	ExampleService_GetNumbers_FullMethodName  = "/example.ExampleService/GetNumbers"
	ExampleService_SendNumbers_FullMethodName = "/example.ExampleService/SendNumbers"
	ExampleService_Square_FullMethodName      = "/example.ExampleService/Square"
)

// ExampleServiceClient is the client API for ExampleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExampleServiceClient interface {
	Greeting(ctx context.Context, in *GreetingRequest, opts ...grpc.CallOption) (*GreetingResponse, error)
	GetNumbers(ctx context.Context, in *GetNumbersRequest, opts ...grpc.CallOption) (ExampleService_GetNumbersClient, error)
	SendNumbers(ctx context.Context, opts ...grpc.CallOption) (ExampleService_SendNumbersClient, error)
	Square(ctx context.Context, opts ...grpc.CallOption) (ExampleService_SquareClient, error)
}

type exampleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewExampleServiceClient(cc grpc.ClientConnInterface) ExampleServiceClient {
	return &exampleServiceClient{cc}
}

func (c *exampleServiceClient) Greeting(ctx context.Context, in *GreetingRequest, opts ...grpc.CallOption) (*GreetingResponse, error) {
	out := new(GreetingResponse)
	err := c.cc.Invoke(ctx, ExampleService_Greeting_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleServiceClient) GetNumbers(ctx context.Context, in *GetNumbersRequest, opts ...grpc.CallOption) (ExampleService_GetNumbersClient, error) {
	stream, err := c.cc.NewStream(ctx, &ExampleService_ServiceDesc.Streams[0], ExampleService_GetNumbers_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &exampleServiceGetNumbersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ExampleService_GetNumbersClient interface {
	Recv() (*Number, error)
	grpc.ClientStream
}

type exampleServiceGetNumbersClient struct {
	grpc.ClientStream
}

func (x *exampleServiceGetNumbersClient) Recv() (*Number, error) {
	m := new(Number)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *exampleServiceClient) SendNumbers(ctx context.Context, opts ...grpc.CallOption) (ExampleService_SendNumbersClient, error) {
	stream, err := c.cc.NewStream(ctx, &ExampleService_ServiceDesc.Streams[1], ExampleService_SendNumbers_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &exampleServiceSendNumbersClient{stream}
	return x, nil
}

type ExampleService_SendNumbersClient interface {
	Send(*Number) error
	CloseAndRecv() (*SendNumbersResponse, error)
	grpc.ClientStream
}

type exampleServiceSendNumbersClient struct {
	grpc.ClientStream
}

func (x *exampleServiceSendNumbersClient) Send(m *Number) error {
	return x.ClientStream.SendMsg(m)
}

func (x *exampleServiceSendNumbersClient) CloseAndRecv() (*SendNumbersResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(SendNumbersResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *exampleServiceClient) Square(ctx context.Context, opts ...grpc.CallOption) (ExampleService_SquareClient, error) {
	stream, err := c.cc.NewStream(ctx, &ExampleService_ServiceDesc.Streams[2], ExampleService_Square_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &exampleServiceSquareClient{stream}
	return x, nil
}

type ExampleService_SquareClient interface {
	Send(*Number) error
	Recv() (*Number, error)
	grpc.ClientStream
}

type exampleServiceSquareClient struct {
	grpc.ClientStream
}

func (x *exampleServiceSquareClient) Send(m *Number) error {
	return x.ClientStream.SendMsg(m)
}

func (x *exampleServiceSquareClient) Recv() (*Number, error) {
	m := new(Number)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ExampleServiceServer is the server API for ExampleService service.
// All implementations must embed UnimplementedExampleServiceServer
// for forward compatibility
type ExampleServiceServer interface {
	Greeting(context.Context, *GreetingRequest) (*GreetingResponse, error)
	GetNumbers(*GetNumbersRequest, ExampleService_GetNumbersServer) error
	SendNumbers(ExampleService_SendNumbersServer) error
	Square(ExampleService_SquareServer) error
	mustEmbedUnimplementedExampleServiceServer()
}

// UnimplementedExampleServiceServer must be embedded to have forward compatible implementations.
type UnimplementedExampleServiceServer struct {
}

func (UnimplementedExampleServiceServer) Greeting(context.Context, *GreetingRequest) (*GreetingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Greeting not implemented")
}
func (UnimplementedExampleServiceServer) GetNumbers(*GetNumbersRequest, ExampleService_GetNumbersServer) error {
	return status.Errorf(codes.Unimplemented, "method GetNumbers not implemented")
}
func (UnimplementedExampleServiceServer) SendNumbers(ExampleService_SendNumbersServer) error {
	return status.Errorf(codes.Unimplemented, "method SendNumbers not implemented")
}
func (UnimplementedExampleServiceServer) Square(ExampleService_SquareServer) error {
	return status.Errorf(codes.Unimplemented, "method Square not implemented")
}
func (UnimplementedExampleServiceServer) mustEmbedUnimplementedExampleServiceServer() {}

// UnsafeExampleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExampleServiceServer will
// result in compilation errors.
type UnsafeExampleServiceServer interface {
	mustEmbedUnimplementedExampleServiceServer()
}

func RegisterExampleServiceServer(s grpc.ServiceRegistrar, srv ExampleServiceServer) {
	s.RegisterService(&ExampleService_ServiceDesc, srv)
}

func _ExampleService_Greeting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GreetingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleServiceServer).Greeting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExampleService_Greeting_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleServiceServer).Greeting(ctx, req.(*GreetingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExampleService_GetNumbers_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetNumbersRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ExampleServiceServer).GetNumbers(m, &exampleServiceGetNumbersServer{stream})
}

type ExampleService_GetNumbersServer interface {
	Send(*Number) error
	grpc.ServerStream
}

type exampleServiceGetNumbersServer struct {
	grpc.ServerStream
}

func (x *exampleServiceGetNumbersServer) Send(m *Number) error {
	return x.ServerStream.SendMsg(m)
}

func _ExampleService_SendNumbers_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ExampleServiceServer).SendNumbers(&exampleServiceSendNumbersServer{stream})
}

type ExampleService_SendNumbersServer interface {
	SendAndClose(*SendNumbersResponse) error
	Recv() (*Number, error)
	grpc.ServerStream
}

type exampleServiceSendNumbersServer struct {
	grpc.ServerStream
}

func (x *exampleServiceSendNumbersServer) SendAndClose(m *SendNumbersResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *exampleServiceSendNumbersServer) Recv() (*Number, error) {
	m := new(Number)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ExampleService_Square_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ExampleServiceServer).Square(&exampleServiceSquareServer{stream})
}

type ExampleService_SquareServer interface {
	Send(*Number) error
	Recv() (*Number, error)
	grpc.ServerStream
}

type exampleServiceSquareServer struct {
	grpc.ServerStream
}

func (x *exampleServiceSquareServer) Send(m *Number) error {
	return x.ServerStream.SendMsg(m)
}

func (x *exampleServiceSquareServer) Recv() (*Number, error) {
	m := new(Number)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ExampleService_ServiceDesc is the grpc.ServiceDesc for ExampleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExampleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "example.ExampleService",
	HandlerType: (*ExampleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Greeting",
			Handler:    _ExampleService_Greeting_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetNumbers",
			Handler:       _ExampleService_GetNumbers_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SendNumbers",
			Handler:       _ExampleService_SendNumbers_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Square",
			Handler:       _ExampleService_Square_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "example/example.proto",
}
