package main

import (
	"log"
	"net"

	"grpc_demo/chat"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp",":9000")
	if err!=nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()

	s := chat.Server{}

	chat.RegisterChatServiceServer(grpcServer, &s) 

	if err := grpcServer.Serve(lis); err !=nil {
		log.Fatal(err)
	}

}