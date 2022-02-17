package main

import (
	"context"
	"fmt"
	pb "grpc-go-101/cmd/clientstream/pingpong"
	"time"

	"google.golang.org/grpc"
)

func SentPing(client pb.PingPongClient) {
	pings := []*pb.Ping{
		&pb.Ping{Message: "ping1"},
		&pb.Ping{Message: "ping2"},
		&pb.Ping{Message: "ping3"},
		&pb.Ping{Message: "ping4"},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := client.ClientStream(ctx)
	if err != nil {
		panic(err)
	}

	for _, ping := range pings {
		stream.Send(ping)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		panic(err)
	}

	fmt.Printf("res %v\n", res)
}

func main() {
	opts := []grpc.DialOption{grpc.WithInsecure()}
	conn, err := grpc.Dial("127.0.0.1:9000", opts...)
	if err != nil {
		panic(err)
	}

	client := pb.NewPingPongClient(conn)
	SentPing(client)
}
