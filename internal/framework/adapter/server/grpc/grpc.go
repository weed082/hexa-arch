package grpc

import (
	"context"
	"log"
	"net"
	"sync"

	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/server/grpc/chat"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/port"
	"google.golang.org/grpc"
)

type Grpc struct {
	wg      *sync.WaitGroup
	ctx     context.Context
	Server  *grpc.Server
	chatApp port.ChatApp
}

func NewServer(wg *sync.WaitGroup, ctx context.Context, chatApp port.ChatApp) *Grpc {
	return &Grpc{
		wg:      wg,
		ctx:     ctx,
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
	chatServer := chat.NewServer(g.wg, g.ctx, g.chatApp)
	chat.RegisterChatServiceServer(g.Server, chatServer) // register chat server
	log.Println(g.Server.Serve(listener))                // serve grpc server
}
