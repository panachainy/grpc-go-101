package main

import (
	"context"
	"fmt"
	pb "grpc-go-101/cmd/serverstream/pingpong"
	"io"

	"google.golang.org/grpc"
)

func ServerStream(client pb.PingPongClient) (string, error) {
	resMessage := ""
	stream, err := client.ServerStream(context.Background(), &pb.Ping{
		Id:      1,
		Message: "ping",
	})
	if err != nil {
		return "", err
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			return "", err
		}

		resMessage += fmt.Sprintf(" [id: %v, message: %v]", res.Id, res.Message)

	}

	return resMessage, nil
}

func SendClient() {
	// insecure for test only
	opts := []grpc.DialOption{grpc.WithInsecure()}

	conn, err := grpc.Dial("127.0.0.1:9000", opts...)
	if err != nil {
		panic(fmt.Sprintf("SendClient 1: %v", err.Error()))
	}

	client := pb.NewPingPongClient(conn)

	res, err := ServerStream(client)
	if err != nil {
		panic(fmt.Sprintf("SendClient 2: %v", err.Error()))
	}

	fmt.Printf("done ping %v\n", res)
}

func main() {
	SendClient()
}
