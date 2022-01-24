package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"strconv"

	"github.com/shivanshu1333/gRPC-examples/ClientStreaming/computeAverageApi/proto"

	"google.golang.org/grpc"
)

type server struct{
	proto.UnimplementedComputeServiceServer
}

func (*server) Compute(stream proto.ComputeService_ComputeServer) error {
	fmt.Printf("Compute function was invoked with a streaming request\n")
	totalVal, totalReq := int64(0), int64(0)
	result := ""
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// we have finished reading the client stream
			avgVal := totalVal/totalReq
			result = "Compute avg is " + strconv.Itoa(int(avgVal))
			return stream.SendAndClose(&proto.Response{
				Result: result,
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}

		val := req.GetValue().GetVal()
		totalVal += val
		totalReq++
	}
}

func main() {
	fmt.Println("Hello world")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterComputeServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}