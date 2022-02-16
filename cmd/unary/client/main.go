package main

import (
	"context"
	"fmt"
	pb "grpc-go-101/pingpong"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func SendPing(client pb.PingPongClient) (*pb.Pong, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ping := &pb.Ping{
		Id:      1,
		Message: "ping",
	}

	pong, err := client.Unary(ctx, ping)
	statusCode := status.Code(err)
	if statusCode != codes.OK {
		return nil, err
	}

	fmt.Printf("Pong received: %s, statusCode: %s\n", pong.GetMessage(), statusCode.String())
	return pong, err
}

func StartPingPongClient() {
	// insecure for test only
	opts := []grpc.DialOption{grpc.WithInsecure()}

	conn, err := grpc.Dial("127.0.0.1:9000", opts...)
	if err != nil {
		panic(err)
	}

	client := pb.NewPingPongClient(conn)

	res, err := SendPing(client)
	if err != nil {
		panic(err)
	}

	fmt.Printf("done ping %v\n", res)
}

func main() {
	StartPingPongClient()
}
