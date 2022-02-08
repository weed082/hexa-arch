package grpc

import (
	"log"
	"net"

	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/server/grpc/chat"
	bidichatpb "github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/server/grpc/chat/pb"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/port"
	"google.golang.org/grpc"
)

type Grpc struct {
	logger  *log.Logger
	server  *grpc.Server
	chatWp  port.Consumer
	chatApp port.Chat
}

func NewServer(logger *log.Logger, chatWp port.Consumer, chatApp port.Chat) *Grpc {
	return &Grpc{
		logger:  logger,
		chatWp:  chatWp,
		chatApp: chatApp,
	}
}

func (g *Grpc) Run(port string) {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		g.logger.Fatalf("failed to listen on port %s", port)
	}

	g.server = grpc.NewServer()
	// bi-directional chat
	bidiChatServer := chat.NewServer(g.logger, g.chatWp, g.chatApp)
	bidichatpb.RegisterChatServiceServer(g.server, bidiChatServer)

	err = g.server.Serve(listener)
	if err != nil {
		g.logger.Fatalf("grpc serve error : %s", err)
	}
	g.logger.Println("grpc: Server closed")
}

func (g *Grpc) Stop() {
	g.server.GracefulStop()
}
