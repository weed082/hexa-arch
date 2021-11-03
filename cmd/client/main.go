package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/server/grpc/chat"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect to: %s", err)
	}
	defer conn.Close()
	client := chat.NewChatServiceClient(conn)
	stream, err := client.ChatService(context.Background())
	if err != nil {
		log.Fatalf("response error: %s", err)
	}
	for {
		message, err := stream.Recv()
		if err != nil {
			log.Fatalf("client stream error: %s", err)
		}
		fmt.Println(message.Name, message.Body)
	}
}
