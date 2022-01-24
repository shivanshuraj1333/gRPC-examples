package main

import (
	"context"
	"fmt"
	"github.com/shivanshu1333/gRPC-examples/BiDirectionalStreaming/findMaximumApi/proto"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
)

func doBiDiStreaming(c proto.MaximumApiClient) {
	fmt.Println("Starting to do a BiDi Streaming RPC...")

	// we create a stream by invoking the client
	stream, err := c.FindMaximum(context.Background())
	if err != nil {
		log.Fatalf("Error while creating stream: %v", err)
		return
	}

	requests := []*proto.Request{
		{
			Value: &proto.Number{
				Val: 1,
			},
		},
		{
			Value: &proto.Number{
				Val: 5,
			},
		},
		{
			Value: &proto.Number{
				Val: 3,
			},
		},
		{
			Value: &proto.Number{
				Val: 6,
			},
		},
		{
			Value: &proto.Number{
				Val: 2,
			},
		},
		{
			Value: &proto.Number{
				Val: 20,
			},
		},
	}

	waitc := make(chan struct{})
	// we send a bunch of messages to the client (go routine)
	go func() {
		// function to send a bunch of messages
		for _, req := range requests {
			fmt.Printf("Sending message: %v\n", req)
			stream.Send(req)
			time.Sleep(1000 * time.Millisecond)
		}
		stream.CloseSend()
	}()
	// we receive a bunch of messages from the client (go routine)
	go func() {
		// function to receive a bunch of messages
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while receiving: %v", err)
			}
			fmt.Printf("Received: %v\n", res.GetResult())
		}
		close(waitc)
	}()
	// block until everything is done
	<-waitc
}

func main() {

	fmt.Println("Hello I'm a client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := proto.NewMaximumApiClient(cc)

	doBiDiStreaming(c)
}