package main

import (
	"context"
	"fmt"
	"github.com/shivanshu1333/gRPC-examples/BiDirectionalStreaming/GreetEveryone/proto"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
)

func doBiDiStreaming(c proto.GreetServiceClient) {
	fmt.Println("Starting to do a BiDi Streaming RPC...")

	// we create a stream by invoking the client
	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error while creating stream: %v", err)
		return
	}

	requests := []*proto.GreetEveryoneRequest{
		{
			Greeting: &proto.Greeting{
				FirstName: "Stephane",
			},
		},
		{
			Greeting: &proto.Greeting{
				FirstName: "John",
			},
		},
		{
			Greeting: &proto.Greeting{
				FirstName: "Lucy",
			},
		},
		{
			Greeting: &proto.Greeting{
				FirstName: "Mark",
			},
		},
		{
			Greeting: &proto.Greeting{
				FirstName: "Piper",
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

	c := proto.NewGreetServiceClient(cc)

	doBiDiStreaming(c)
}