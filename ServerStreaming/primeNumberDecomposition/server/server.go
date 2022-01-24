package main

import (
	"fmt"
	"github.com/shivanshu1333/gRPC-examples/ServerStreaming/primeNumberDecomposition/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
	"time"
)

type server struct{
	proto.UnimplementedStreamServiceServer
}

func (s *server) Stream(req *proto.Request, stream proto.StreamService_StreamServer) error {
	fmt.Printf("Received a streaming request %v\n", req)
	num := req.GetRequest().GetValue()
	var k int64 = 2
	for {
		if num<=1 {
			break
		}
		if num % k == 0 {
			num = num/k

			result := "Started streaming service " + " prime number: " + strconv.Itoa(int(k))
			res := &proto.Response{
				Result: result,
			}
			err := stream.Send(res)
			if err != nil {
				return err
			}
			time.Sleep(100 * time.Millisecond)

		} else {
			k = k + 1
		}
	}
	return nil
}


const (
	address = "0.0.0.0:50051"
)

func main() {
	fmt.Println("Basic ServerStreaming Service initiating....")

	lis, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterStreamServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}