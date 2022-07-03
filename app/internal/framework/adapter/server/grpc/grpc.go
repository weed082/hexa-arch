package grpc

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

type Grpc struct {
	logger *log.Logger
	server *grpc.Server
}

func NewServer(logger *log.Logger) *Grpc {
	return &Grpc{logger: logger}
}

func (g *Grpc) Run(port string) {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		g.logger.Fatalf("failed to listen on port %s", port)
	}
	g.server = grpc.NewServer()
	err = g.server.Serve(listener)
	if err != nil {
		g.logger.Fatalf("grpc serve error : %s", err)
	}
	g.logger.Println("grpc: Server closed")
}

func (g *Grpc) Stop() {
	g.server.GracefulStop()
}
