package main

import (
	"context"
	"fmt"
	pb "grpc-go-101/cmd/bidistream/pingpong"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
)

func SendPing(c pb.PingPongClient) {
	fmt.Println("Start ping..")

	stream, err := c.BidirectionalStream(context.Background())
	if err != nil {
		log.Fatalf("Error create stream: %v\n", err)
		return
	}

	pings := []*pb.Ping{
		{Id: 1, Message: "ping1"},
		{Id: 2, Message: "ping2"},
		{Id: 3, Message: "ping3"},
		{Id: 4, Message: "ping4"},
		{Id: 5, Message: "ping5"},
		{Id: 6, Message: "ping6"},
	}

	waitc := make(chan struct{})

	go func() {
		for _, ping := range pings {
			fmt.Printf("Ping id: %v message: %v\n", ping.Id, ping.Message)
			err := stream.Send(ping)
			if err != nil {
				log.Fatalf("Error while sending request: %v\n", err)
				break
			}
			time.Sleep(1000 * time.Millisecond)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while receiveing: %v\n", err)
				break
			}
			fmt.Printf("Received Id: %v message: %v\n", res.Id, res.Message)
		}
		close(waitc)
	}()
	// block until everthing is done
	<-waitc
}

func StartPingPongClient() {
	// insecure for test only
	opts := []grpc.DialOption{grpc.WithInsecure()}

	conn, err := grpc.Dial("127.0.0.1:9000", opts...)
	if err != nil {
		panic(err)
	}

	client := pb.NewPingPongClient(conn)

	SendPing(client)
	fmt.Println("done send ping")
}

func main() {
	StartPingPongClient()
}
