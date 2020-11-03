package main

import (
	"fmt"
	"log"
	"net"

	"github.com/mansikalra23/Microservices/gRPC/chat"
	"google.golang.org/grpc" // official grpc package
)

// main function will listen on the port for incomming tcp connections
func main() {

	fmt.Println("gRPC Basic Program.")

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("unable to listen: %v", err)
	}

	s := chat.Server{}

	// registering the endpoints
	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("unable to serve: %s", err)
	}
}
