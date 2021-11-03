package chat

import (
	"io"
	"log"
)

type Server struct {
	streams []ChatService_ChatServiceServer
	msgCh   chan *Message
}

func NewServer() *Server {
	return &Server{
		streams: []ChatService_ChatServiceServer{},
		msgCh:   make(chan *Message),
	}
}

func (s *Server) ChatService(stream ChatService_ChatServiceServer) error {
	s.streams = append(s.streams, stream)
	log.Printf("seding : %d", len(s.streams))

	go s.sendMessage()              // send message to client
	return s.receiveMessage(stream) // receive message from client
}

func (s *Server) sendMessage() {
	for {
		msg := <-s.msgCh
		for _, stream := range s.streams {
			err := stream.Send(msg)
			//TODO: error should be send to the channel
			if err != nil {
				log.Fatalf("sending message error: %s", err)
				break
			}
		}
	}
}

func (s *Server) receiveMessage(stream ChatService_ChatServiceServer) error {
	for {
		message, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("receiving message err: %s", err)
			return err
		}
		s.msgCh <- message
	}
}
