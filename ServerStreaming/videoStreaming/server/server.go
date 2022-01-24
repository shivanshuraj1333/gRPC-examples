package main

import (
	"fmt"
	"github.com/shivanshu1333/gRPC-examples/ServerStreaming/videoStreaming/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
	"time"
)

type server struct{
	proto.UnimplementedStreamServiceServer
}

func (s *server) Stream(req *proto.StreamRequest, stream proto.StreamService_StreamServer) error {
	fmt.Printf("Received a streaming request %v\n", req)
	videoName := req.GetStream().GetVideoName()
	for i := 0; i < 10; i++ {
		result := "Started streaming service " + videoName + " res id: " + strconv.Itoa(i)
		res := &proto.StreamResponse{
			Result: result,
		}
		err := stream.Send(res)
		if err != nil {
			return err
		}
		time.Sleep(100 * time.Millisecond)
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