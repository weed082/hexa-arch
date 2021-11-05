package grpc

import (
	"log"
	"net"

	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/server/grpc/chat"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/port"
	"google.golang.org/grpc"
)

type Server struct {
	fileApp port.FileApp
	userApp port.UserApp
}

// TODO: param need user app
func NewServer(fileApp port.FileApp) Server {
	return Server{
		fileApp: fileApp,
	}
}

func (s *Server) Run(port string) {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen on port %s", port)
	}
	defer listener.Close()

	grpcServer := grpc.NewServer()
	chatServer := chat.NewServer(s.fileApp)
	chat.RegisterChatServiceServer(grpcServer, chatServer) // register chat server
	log.Fatal(grpcServer.Serve(listener))                  // serve grpc server
}
