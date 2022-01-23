package main

import (
	"context"
	"fmt"
	"github.com/shivanshu1333/gRPC-examples/unary/example3/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct{
	proto.UnimplementedCalServiceServer
}

func (s *server) AddNumbers(ctx context.Context, req *proto.Request) (*proto.Response, error) {
	value := req.Add.GetVal1() + req.Add.GetVal2()
	res := &proto.Response{
		Res: value,
	}
	return res, nil
}

func main() {
	fmt.Println("Basic Unary Service initiating....")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterCalServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatal("Failed to serve: %v", err)
	}
}