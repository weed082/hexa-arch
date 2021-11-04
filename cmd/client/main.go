package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/server/grpc/chat"
	"google.golang.org/grpc"
)

func main() {
	for i := 0; i < 3000; i++ {
		go runClient()
	}
	runClient()
}

func runClient() {
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

	go receiveMessage(stream)
	sendMessage(stream)
}

func sendMessage(stream chat.ChatService_ChatServiceClient) {

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		err := stream.Send(&chat.Message{Name: "client", Body: scanner.Text()})
		if err != nil {
			log.Fatalf("send to server err: %s", err)
			break
		}
	}
}

func receiveMessage(stream chat.ChatService_ChatServiceClient) {
	for {
		message, err := stream.Recv()
		if err != nil {
			log.Fatalf("client stream error: %s", err)
			break
		}
		fmt.Println(message.Name, message.Body)
	}
}
