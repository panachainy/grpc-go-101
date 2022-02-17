package main

import (
	"fmt"
	pb "grpc-go-101/cmd/clientstream/pingpong"
	"io"
	"net"

	"google.golang.org/grpc"
)

type PingPongServerImpl struct {
	pb.UnimplementedPingPongServer
}

func (s *PingPongServerImpl) ClientStream(stream pb.PingPong_ClientStreamServer) error {
	pingCount := 0
	message := ""

	for {
		ping, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.Pong{
				Id:      1,
				Message: fmt.Sprintf("combine msg: %s, count: %d", message, pingCount),
			})
		}

		if err != nil {
			return err
		}

		pingCount++

		message += ping.Message
	}
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
