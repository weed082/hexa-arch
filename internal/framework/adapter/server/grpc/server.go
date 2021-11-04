package grpc

import (
	"log"
	"net"

	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/server/grpc/chat"
	"google.golang.org/grpc"
)

type Server struct {
}

func NewServer() Server {
	return Server{}
}

func (s *Server) Run(port string) {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen on port %s", port)
	}
	defer listener.Close()

	grpcServer := grpc.NewServer()
	chatServer := chat.NewServer()
	chat.RegisterChatServiceServer(grpcServer, chatServer) // register chat server
	log.Fatal(grpcServer.Serve(listener))                  // serve grpc server
}
