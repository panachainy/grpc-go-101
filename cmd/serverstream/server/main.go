package main

import (
	"fmt"
	pb "grpc-go-101/cmd/serverstream/pingpong"
	"net"

	"google.golang.org/grpc"
)

type PingPongServer struct {
	pb.UnimplementedPingPongServer
}

func (s *PingPongServer) ServerStream(p *pb.Ping, stream pb.PingPong_ServerStreamServer) error {
	pongs := []*pb.Pong{
		{Id: 1, Message: "ping1"},
		{Id: 2, Message: "ping2"},
		{Id: 3, Message: "ping3"},
		{Id: 4, Message: "ping4"},
		{Id: 5, Message: "ping5"},
		{Id: 0, Message: "ping"},
		{Id: 0, Message: "ping"},
		{Id: 0, Message: "ping"},
		{Id: 0, Message: "ping"},
		{Id: 0, Message: "ping"},
		{Id: 0, Message: "ping"},
		{Id: 0, Message: "ping"},
		{Id: 0, Message: "ping"},
		{Id: 0, Message: "ping"},
		{Id: 0, Message: "ping"},
		{Id: 0, Message: "ping"},
		{Id: 0, Message: "ping"},
		{Id: 99, Message: "ping last"},
	}

	for _, pong := range pongs {
		fmt.Println("================= inner loop")

		err := stream.Send(pong)
		if err != nil {
			panic(err)
		}
	}

	// when return nil is end of req, Client can check with io.EOF
	return nil
}

func StartPingPongServer() {
	server := PingPongServer{}

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
