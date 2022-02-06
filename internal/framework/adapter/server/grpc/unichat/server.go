package unichat

import (
	"context"
	"log"

	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/server/grpc/unichat/pb"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/port"
)

type Server struct {
	logger   *log.Logger
	chatPool port.Consumer
	chatApp  port.Chat
	rooms    map[int][]port.Client
}

func NewServer(logger *log.Logger, chatPool port.Consumer, app port.Chat) *Server {
	return &Server{
		logger:   logger,
		chatPool: chatPool,
		chatApp:  app,
		rooms:    make(map[int][]port.Client),
	}
}

func (s *Server) CreateRoom(ctx context.Context, user *pb.User) (*pb.Empty, error) {
	return &pb.Empty{}, nil
}

func (s *Server) JoinRoom(ctx context.Context, user *pb.Participant) (*pb.Empty, error) {
	return &pb.Empty{}, nil
}

func (s *Server) ExitRoom(ctx context.Context, user *pb.Participant) (*pb.Empty, error) {
	return &pb.Empty{}, nil
}

func (s *Server) SendMsg(ctx context.Context, user *pb.Msg) (*pb.Empty, error) {
	return &pb.Empty{}, nil
}

func (s *Server) ReceiveMsg(user *pb.User, stream pb.ChatService_ReceiveMsgServer) error {
	return nil
}
