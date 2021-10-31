package chat

import (
	"context"
	"log"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("received message from server : %s", message)
	return &Message{Body: "Hello Form the Server!"}, nil
}
