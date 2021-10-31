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
	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &chat.Server{})
	if grpcServer.Serve(listener) != nil {
		log.Fatalf("failed to serve grpc server over port %s", port)
	}
}
