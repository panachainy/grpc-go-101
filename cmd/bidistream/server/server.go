package main

import (
	"fmt"
	pb "grpc-go-101/cmd/bidistream/pingpong"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
)

type PingPongServer struct {
	pb.UnimplementedPingPongServer
}

func (s *PingPongServer) ServerStream(stream pb.PingPong_BidirectionalStreamServer) error {
	fmt.Println("in coming...")

	allReqMessage := ""
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			fmt.Println(allReqMessage)
			fmt.Println("====================== done receive ======================")
			return nil
		}
		if err != nil {
			log.Fatalf("Error: %v\n", err)
		}
		allReqMessage += fmt.Sprintf(" [id: %v, message: %v]", req.Id, req.Message)

		sendErr := stream.Send(&pb.Pong{Id: req.Id, Message: req.Message + " [pong]"})
		if sendErr != nil {
			log.Fatalf("Error while sending data to client: %v", sendErr)
			return sendErr
		}
	}
}

func StartPingPongServer() {
	server := PingPongServer{}

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterPingPongServer(grpcServer, &server)

	fmt.Println("starting server")

	if err = grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}

func main() {
	StartPingPongServer()
}
