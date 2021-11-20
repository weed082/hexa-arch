package grpc

import (
	"log"
	"net"

	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/server/grpc/chat"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/server/grpc/chat/pb"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/port"
	"google.golang.org/grpc"
)

type Grpc struct {
	server  *grpc.Server
	chatWp  port.WorkerPool
	chatApp port.Chat
}

func NewServer(chatWp port.WorkerPool, chatApp port.Chat) *Grpc {
	return &Grpc{
		chatWp:  chatWp,
		chatApp: chatApp,
	}
}

func (g *Grpc) Run(port string) {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen on port %s", port)
	}

	g.server = grpc.NewServer()
	chatServer := chat.NewServer(g.chatWp, g.chatApp)
	pb.RegisterChatServiceServer(g.server, chatServer) // register chat server

	err = g.server.Serve(listener)
	if err != nil {
		log.Fatalf("grpc serve error : %s", err)
	}
	log.Println("grpc: Server closed")
}

func (g *Grpc) Stop() {
	g.server.GracefulStop()
}
