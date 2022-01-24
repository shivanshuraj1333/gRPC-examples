package main

import (
	"context"
	"fmt"
	"github.com/shivanshu1333/gRPC-examples/ClientStreaming/longGreet/proto"
	"google.golang.org/grpc"
	"log"
	"time"
)

func doClientStreaming(c proto.GreetServiceClient) {
	fmt.Println("Starting to do a Client Streaming RPC...")

	requests := []*proto.LongGreetRequest{
		&proto.LongGreetRequest{
			Greeting: &proto.Greeting{
				FirstName: "Stephane",
			},
		},
		&proto.LongGreetRequest{
			Greeting: &proto.Greeting{
				FirstName: "John",
			},
		},
		&proto.LongGreetRequest{
			Greeting: &proto.Greeting{
				FirstName: "Lucy",
			},
		},
		&proto.LongGreetRequest{
			Greeting: &proto.Greeting{
				FirstName: "Mark",
			},
		},
		&proto.LongGreetRequest{
			Greeting: &proto.Greeting{
				FirstName: "Piper",
			},
		},
	}

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("error while calling LongGreet: %v", err)
	}

	// we iterate over our slice and send each message individually
	for _, req := range requests {
		fmt.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1000 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receiving response from LongGreet: %v", err)
	}
	fmt.Printf("LongGreet Response: %v\n", res)

}

const (
	address = "0.0.0.0:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Error occured %v", err)
	}
	defer conn.Close()

	c := proto.NewGreetServiceClient(conn)

	doClientStreaming(c)
}