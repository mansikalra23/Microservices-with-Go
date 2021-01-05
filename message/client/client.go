package main

import (
	"fmt"
	"log"

	"github.com/mansikalra23/Microservices/gRPC/message/chat"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":3030", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	c := chat.NewChatServiceClient(conn)
	var msg string
	fmt.Print("Enter message for server : ")
	fmt.Scan(&msg)
	response, err := c.SayHello(context.Background(), &chat.Message{Body: msg})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.Body)
}
