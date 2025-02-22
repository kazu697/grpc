// proto bufferのバージョンを指定

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: src/api/hello.proto

// proto bufferをパッケージ化する
// 他のproto bufferを参照できるようになる

package grpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	GreetingService_Hello_FullMethodName             = "/hello.GreetingService/Hello"
	GreetingService_HelloServerStream_FullMethodName = "/hello.GreetingService/HelloServerStream"
)

// GreetingServiceClient is the client API for GreetingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// サービス　メソッドをまとめたもの
type GreetingServiceClient interface {
	// メソッド（grpcで呼び出され関数（Procedure）
	// HelloRequest・HelloResponseは型
	Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
	// サーバーストリーミングRPC
	HelloServerStream(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (GreetingService_HelloServerStreamClient, error)
}

type greetingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGreetingServiceClient(cc grpc.ClientConnInterface) GreetingServiceClient {
	return &greetingServiceClient{cc}
}

func (c *greetingServiceClient) Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, GreetingService_Hello_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greetingServiceClient) HelloServerStream(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (GreetingService_HelloServerStreamClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &GreetingService_ServiceDesc.Streams[0], GreetingService_HelloServerStream_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &greetingServiceHelloServerStreamClient{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GreetingService_HelloServerStreamClient interface {
	Recv() (*HelloResponse, error)
	grpc.ClientStream
}

type greetingServiceHelloServerStreamClient struct {
	grpc.ClientStream
}

func (x *greetingServiceHelloServerStreamClient) Recv() (*HelloResponse, error) {
	m := new(HelloResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreetingServiceServer is the server API for GreetingService service.
// All implementations must embed UnimplementedGreetingServiceServer
// for forward compatibility
//
// サービス　メソッドをまとめたもの
type GreetingServiceServer interface {
	// メソッド（grpcで呼び出され関数（Procedure）
	// HelloRequest・HelloResponseは型
	Hello(context.Context, *HelloRequest) (*HelloResponse, error)
	// サーバーストリーミングRPC
	HelloServerStream(*HelloRequest, GreetingService_HelloServerStreamServer) error
	mustEmbedUnimplementedGreetingServiceServer()
}

// UnimplementedGreetingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGreetingServiceServer struct {
}

func (UnimplementedGreetingServiceServer) Hello(context.Context, *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hello not implemented")
}
func (UnimplementedGreetingServiceServer) HelloServerStream(*HelloRequest, GreetingService_HelloServerStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method HelloServerStream not implemented")
}
func (UnimplementedGreetingServiceServer) mustEmbedUnimplementedGreetingServiceServer() {}

// UnsafeGreetingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreetingServiceServer will
// result in compilation errors.
type UnsafeGreetingServiceServer interface {
	mustEmbedUnimplementedGreetingServiceServer()
}

func RegisterGreetingServiceServer(s grpc.ServiceRegistrar, srv GreetingServiceServer) {
	s.RegisterService(&GreetingService_ServiceDesc, srv)
}

func _GreetingService_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreetingServiceServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GreetingService_Hello_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreetingServiceServer).Hello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GreetingService_HelloServerStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(HelloRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GreetingServiceServer).HelloServerStream(m, &greetingServiceHelloServerStreamServer{ServerStream: stream})
}

type GreetingService_HelloServerStreamServer interface {
	Send(*HelloResponse) error
	grpc.ServerStream
}

type greetingServiceHelloServerStreamServer struct {
	grpc.ServerStream
}

func (x *greetingServiceHelloServerStreamServer) Send(m *HelloResponse) error {
	return x.ServerStream.SendMsg(m)
}

// GreetingService_ServiceDesc is the grpc.ServiceDesc for GreetingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GreetingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hello.GreetingService",
	HandlerType: (*GreetingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _GreetingService_Hello_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "HelloServerStream",
			Handler:       _GreetingService_HelloServerStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "src/api/hello.proto",
}
