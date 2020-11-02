package main

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.Clientconn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect: %s", err)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	message := chat.Message{
		Body: "Hello Kloudone!",
	}

	response, err := c.SayHello(context.Background(), &message)

	if err != nil {
		log.Fatalf("Error in SayHello, %s", err)
	}

	log.Printf("Respomse  : %s", response.Body)
}
