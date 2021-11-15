package grpc

import (
	"log"
	"net"

	"github.com/ByungHakNoh/hexagonal-microservice/internal/core"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/server/grpc/chat"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/port"
	"google.golang.org/grpc"
)

type Grpc struct {
	Server  *grpc.Server
	wp      *core.WorkerPool
	chatApp port.ChatApp
}

func NewServer(wp *core.WorkerPool, chatApp port.ChatApp) *Grpc {
	return &Grpc{
		wp:      wp,
		chatApp: chatApp,
	}
}

func (g *Grpc) Run(port string) {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen on port %s", port)
	}
	defer listener.Close()
	g.Server = grpc.NewServer()
	chatServer := chat.NewServer(g.wp, g.chatApp)
	chat.RegisterChatServiceServer(g.Server, chatServer) // register chat server
	log.Println(g.Server.Serve(listener))                // serve grpc server
}
