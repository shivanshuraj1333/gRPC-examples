package main

import (
	"context"
	"github.com/shivanshu1333/gRPC-examples/unary/example3/proto"
	"google.golang.org/grpc"
	"log"
)

const (
	address = "0.0.0.0:50051"
)

func makeUnaryCall(c proto.CalServiceClient) {
	req := &proto.Request{
		Add: &proto.Add{
			Val1: 22,
			Val2: 27,
		},
	}
	res, err := c.AddNumbers(context.Background(), req)
	if err != nil {
		log.Fatalf("Error occured, %v", err)
	}
	log.Printf("Response from server: %v", res.Res)
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Error occured %v", err)
	}
	defer conn.Close()

	c := proto.NewCalServiceClient(conn)

	makeUnaryCall(c)
}