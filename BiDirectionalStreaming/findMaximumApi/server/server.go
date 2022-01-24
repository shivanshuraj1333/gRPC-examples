package main

import (
	"fmt"
	"github.com/shivanshu1333/gRPC-examples/BiDirectionalStreaming/findMaximumApi/proto"
	"google.golang.org/grpc"
	"io"
	"log"
	"math"
	"net"
)

type server struct{
	proto.UnimplementedMaximumApiServer
}

func (*server) FindMaximum(stream proto.MaximumApi_FindMaximumServer) error {
	fmt.Printf("FindMaximum function invoked with a streaming request\n")
	var curMaxVal int64 = math.MinInt64
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
			return err
		}
		curVal := req.GetValue().GetVal()
		if curVal > curMaxVal {
			curMaxVal = curVal
			sendErr := stream.Send(&proto.Response{
				Result: curMaxVal,
			})
			if sendErr != nil {
				log.Fatalf("Error while sending data to client: %v", sendErr)
				return sendErr
			}
		}
	}
}

func main() {
	fmt.Println("Hello world")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterMaximumApiServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}