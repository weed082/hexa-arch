package chat

import (
	"io"
	"log"

	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/port"
)

type Server struct {
	streams    []ChatService_ChatServiceServer
	msgChan    chan *Message
	streamChan chan ChatService_ChatServiceServer
	fileApp    port.FileApp
}

func NewServer(fileApp port.FileApp) *Server {
	server := &Server{
		streams:    []ChatService_ChatServiceServer{},
		msgChan:    make(chan *Message),
		streamChan: make(chan ChatService_ChatServiceServer),
		fileApp:    fileApp,
	}
	// TODO: need to find a way to close it gracefully + make this as a worker pool
	go server.work()
	return server
}

func (s *Server) ChatService(stream ChatService_ChatServiceServer) error {
	s.streamChan <- stream

	for {
		message, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("receiving message err: %s", err)
			return err
		}
		s.msgChan <- message
	}
}

func (s *Server) work() {
	for {
		select {
		case msg := <-s.msgChan:
			s.sendMessage(msg)
		case stream := <-s.streamChan:
			s.streams = append(s.streams, stream)
			log.Printf("current client count : %d", len(s.streams))
		}
	}
}

func (s Server) sendMessage(msg *Message) {
	for _, stream := range s.streams {
		err := stream.Send(msg)
		if err != nil {
			log.Fatalf("sending message error: %s", err)
			continue
		}
	}
}
