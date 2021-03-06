// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: cmd/serverstream/pingpong/pingpong-serverstream.proto

package pingpong

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

// PingPongClient is the client API for PingPong service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PingPongClient interface {
	ServerStream(ctx context.Context, in *Ping, opts ...grpc.CallOption) (PingPong_ServerStreamClient, error)
}

type pingPongClient struct {
	cc grpc.ClientConnInterface
}

func NewPingPongClient(cc grpc.ClientConnInterface) PingPongClient {
	return &pingPongClient{cc}
}

func (c *pingPongClient) ServerStream(ctx context.Context, in *Ping, opts ...grpc.CallOption) (PingPong_ServerStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &PingPong_ServiceDesc.Streams[0], "/pingpong.PingPong/ServerStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &pingPongServerStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PingPong_ServerStreamClient interface {
	Recv() (*Pong, error)
	grpc.ClientStream
}

type pingPongServerStreamClient struct {
	grpc.ClientStream
}

func (x *pingPongServerStreamClient) Recv() (*Pong, error) {
	m := new(Pong)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PingPongServer is the server API for PingPong service.
// All implementations must embed UnimplementedPingPongServer
// for forward compatibility
type PingPongServer interface {
	ServerStream(*Ping, PingPong_ServerStreamServer) error
	mustEmbedUnimplementedPingPongServer()
}

// UnimplementedPingPongServer must be embedded to have forward compatible implementations.
type UnimplementedPingPongServer struct {
}

func (UnimplementedPingPongServer) ServerStream(*Ping, PingPong_ServerStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ServerStream not implemented")
}
func (UnimplementedPingPongServer) mustEmbedUnimplementedPingPongServer() {}

// UnsafePingPongServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PingPongServer will
// result in compilation errors.
type UnsafePingPongServer interface {
	mustEmbedUnimplementedPingPongServer()
}

func RegisterPingPongServer(s grpc.ServiceRegistrar, srv PingPongServer) {
	s.RegisterService(&PingPong_ServiceDesc, srv)
}

func _PingPong_ServerStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Ping)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PingPongServer).ServerStream(m, &pingPongServerStreamServer{stream})
}

type PingPong_ServerStreamServer interface {
	Send(*Pong) error
	grpc.ServerStream
}

type pingPongServerStreamServer struct {
	grpc.ServerStream
}

func (x *pingPongServerStreamServer) Send(m *Pong) error {
	return x.ServerStream.SendMsg(m)
}

// PingPong_ServiceDesc is the grpc.ServiceDesc for PingPong service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PingPong_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pingpong.PingPong",
	HandlerType: (*PingPongServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ServerStream",
			Handler:       _PingPong_ServerStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "cmd/serverstream/pingpong/pingpong-serverstream.proto",
}
