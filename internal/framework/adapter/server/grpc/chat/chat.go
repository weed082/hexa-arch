package chat

import (
	"fmt"
	"log"
	"time"
)

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) ChatService(csi ChatService_ChatServiceServer) error {
	fmt.Println("client connected")
	for {
		err := csi.Send(&Message{Name: "name1", Body: "body1"})
		if err != nil {
			log.Fatalf("error sending message: %s", err)
			break
		}
		time.Sleep(1 * time.Second)
	}
	return nil
}
