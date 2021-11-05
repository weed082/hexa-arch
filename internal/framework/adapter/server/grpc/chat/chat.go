package chat

import (
	"io"
	"log"

	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/port"
)

// test result : can manage 7935 clients
type Server struct {
	streams []ChatService_ChatServiceServer
	msgCh   chan *Message
	fileApp port.FileApp
}

func NewServer(fileApp port.FileApp) *Server {
	server := &Server{
		streams: []ChatService_ChatServiceServer{},
		msgCh:   make(chan *Message),
		fileApp: fileApp,
	}
	go server.sendMessage() // TODO: need to make worker pool to handle sending message
	return server
}

func (server *Server) ChatService(stream ChatService_ChatServiceServer) error {
	server.streams = append(server.streams, stream)
	log.Printf("current client count : %d", len(server.streams))

	for {
		message, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("receiving message err: %s", err)
			return err
		}

		// TODO: if msg type is file

		server.msgCh <- message // send message to msgCh
	}
}

func (server *Server) sendMessage() {
	for {
		msg := <-server.msgCh
		for _, stream := range server.streams {
			err := stream.Send(msg)
			//TODO: error should be send to the channel
			if err != nil {
				log.Fatalf("sending message error: %s", err)
				break
			}
		}
	}
}
