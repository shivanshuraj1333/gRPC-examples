package main

import (
	"context"
	"github.com/shivanshu1333/gRPC-examples/unary/example1/greetpb"
	"google.golang.org/grpc"
	"log"
)

const (
	address = "0.0.0.0:50051"
)

func makeUnaryCall(c greetpb.GreetServiceClient) {	req := &greetpb.GreetRequest{
	Greeting: &greetpb.Greeting{
		FirstName: "shivanshu",
		LastName: "shrivastava",
	},
}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("Error occured, %v", err)
	}
	log.Printf("Response from server: %v", res.Result)
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Error occured %v", err)
	}
	defer conn.Close()

	c := greetpb.NewGreetServiceClient(conn)

	makeUnaryCall(c)
}

