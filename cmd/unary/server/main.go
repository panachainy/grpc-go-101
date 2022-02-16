package main

import (
	"context"
	"fmt"
	pb "grpc-go-101/pingpong"
	"net"

	"google.golang.org/grpc"
)

type PingPongServerImpl struct{ pb.UnimplementedPingPongServer }

func (s *PingPongServerImpl) Unary(ctx context.Context, ping *pb.Ping) (*pb.Pong, error) {
	fmt.Println("Ping received: ", ping.GetMessage())

	resp := pb.Pong{
		Id:      ping.Id,
		Message: ping.Message,
	}

	return &resp, nil
}

func StartPingPongServer() {
	server := PingPongServerImpl{}

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterPingPongServer(grpcServer, &server)

	if err = grpcServer.Serve(lis); err != nil {
		panic(err)
	}

	fmt.Println("pingpong started")
}

func main() {
	StartPingPongServer()
}
