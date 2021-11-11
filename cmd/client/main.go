package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/server/grpc/chat"
	"google.golang.org/grpc"
)

var count = 0
var mutex = sync.Mutex{}

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 3000; i++ {
		wg.Add(1)
		fmt.Printf("client count: %d", i)
		go runClient(wg)
	}
	wg.Wait()
}

func runClient(wg sync.WaitGroup) {
	defer wg.Done()
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
		mutex.Lock()
		count++
		mutex.Unlock()
		fmt.Println(message.Name, message.Body)
		fmt.Printf("message count: %d\n", count)
	}
}
