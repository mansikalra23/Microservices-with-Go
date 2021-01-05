package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/mansikalra23/Microservices/gRPC/message/chat"

	"google.golang.org/grpc"
)

// Server ....
type Server struct {
}

// SayHello ...
func (s *Server) SayHello(ctx context.Context, in *chat.Message) (*chat.Message, error) {
	log.Printf("Receive message body from client: %s", in.Body)
	var msg string
	fmt.Print("Enter response for client : ")
	fmt.Scan(&msg)
	return &chat.Message{Body: msg}, nil
}
func main() {
	fmt.Println("Messge service using grpc")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 3030))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := Server{}
	grpcServer := grpc.NewServer()
	chat.RegisterChatServiceServer(grpcServer, &s)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
