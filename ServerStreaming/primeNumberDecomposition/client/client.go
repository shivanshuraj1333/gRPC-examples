package main

import (
	"context"
	"fmt"
	"github.com/shivanshu1333/gRPC-examples/ServerStreaming/primeNumberDecomposition/proto"
	"google.golang.org/grpc"
	"io"
	"log"
)

func doServerStreaming(c proto.StreamServiceClient) {
	fmt.Println("Starting to do a streaming RPC...")

	req := &proto.Request{
		Request: &proto.Number{
			Value: 2727,
		},
	}
	resStream, err := c.Stream(context.Background(), req)
	if err != nil {
		log.Fatalf("error occured %v", err)
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			// we've reached the end of stream
			break
		}
		if err != nil {
			log.Fatalf("error while reading stream: %v", err)
		}
		log.Printf("Response: %v", msg.GetResult())
	}
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

	c := proto.NewStreamServiceClient(conn)

	doServerStreaming(c)
}