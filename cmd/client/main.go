package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/server/grpc/chat"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/server/grpc/chat/pb"
	"google.golang.org/grpc"
)

const userIdx = 1

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 1; i++ {
		wg.Add(1)
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
	client := pb.NewChatServiceClient(conn)
	stream, err := client.ChatService(context.Background())
	if err != nil {
		log.Fatalf("response error: %s", err)
	}

	go receiveMessage(stream)
	sendMessage(stream)
}

func sendMessage(stream pb.ChatService_ChatServiceClient) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var msg *pb.MsgReq
		switch scanner.Text() {
		case "create":
			msg = &pb.MsgReq{Request: chat.CREATE_ROOM_REQ, UserIdx: userIdx}
		case "join":
			msg = &pb.MsgReq{Request: chat.JOIN_ROOM_REQ, RoomIdx: 1, UserIdx: userIdx}
		case "exit":
			msg = &pb.MsgReq{Request: chat.EXIT_ROOM_REQ, RoomIdx: 1, UserIdx: userIdx}
		default:
			msg = &pb.MsgReq{Request: chat.TEXT_MSG_REQ, RoomIdx: 1, UserIdx: userIdx, Body: scanner.Text()}
		}
		err := stream.Send(msg)
		if err != nil {
			log.Fatalf("send to server err: %s", err)
			break
		}
	}
}

func receiveMessage(stream pb.ChatService_ChatServiceClient) {
	for {
		message, err := stream.Recv()
		if err != nil {
			log.Fatalf("client stream error: %s", err)
			break
		}
		fmt.Println(message.Body)
	}
}
