package main

import (
	"context"
	"fmt"
	"github.com/shivanshu1333/gRPC-examples/ClientStreaming/computeAverageApi/proto"
	"google.golang.org/grpc"
	"log"
	"time"
)

func doClientStreaming(c proto.ComputeServiceClient) {
	fmt.Println("Starting to do a Client Streaming RPC...")

	requests := []*proto.Request{
		{
			Value : &proto.Compute{
				Val: 1,
			},
		},
		{
			Value : &proto.Compute{
				Val: 2,
			},
		},
		{
			Value : &proto.Compute{
				Val: 3,
			},
		},
		{
			Value : &proto.Compute{
				Val: 4,
			},
		},
		{
			Value : &proto.Compute{
				Val: 5,
			},
		},
		{
			Value : &proto.Compute{
				Val: 6,
			},
		},
		{
			Value : &proto.Compute{
				Val: 7,
			},
		},
	}

	stream, err := c.Compute(context.Background())
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

	c := proto.NewComputeServiceClient(conn)

	doClientStreaming(c)
}