package main

import (
	"google.golang.org/grpc"
	"log"
)

const (
	address = "0.0.0.0:50051"
)

func main(){
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Printf("Error occured")
	}
	defer conn.Close()

}
