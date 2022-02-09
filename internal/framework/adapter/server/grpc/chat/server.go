package chat

import (
	"io"
	"log"

	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/server/grpc/chat/pb"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/port"
)

const (
	CREATE_ROOM_REQ = 1
	JOIN_ROOM_REQ   = 2
	EXIT_ROOM_REQ   = 3
	TEXT_MSG_REQ    = 4
	IMAGE_MSG_REQ   = 5

	ERROR_MSG_RES = 1
)

type Server struct {
	logger *log.Logger
	app    port.Chat
}

func NewServer(logger *log.Logger, app port.Chat) *Server {
	return &Server{
		logger: logger,
		app:    app,
	}
}

//! --------------------- (1) grpc request ---------------------
func (s *Server) ChatService(stream pb.ChatService_ChatServiceServer) error {
	c := &Client{stream: stream}
	defer s.app.DisconnectAll(c)

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			s.logger.Printf("receiving message err: %s", err)
			return err
		}
		if c.userIdx == 0 {
			c.userIdx = int(msg.UserIdx)
		}
		switch msg.Request {
		case CREATE_ROOM_REQ:
			s.createRoom(c)
		case JOIN_ROOM_REQ:
			s.app.Join(int(msg.RoomIdx), c)
		case EXIT_ROOM_REQ:
			s.app.Exit(int(msg.RoomIdx), int(msg.UserIdx))
		case TEXT_MSG_REQ:
			msg := &Message{&pb.MsgRes{RoomIdx: msg.RoomIdx, UserIdx: msg.UserIdx, Body: msg.Body}}
			s.app.SendMsg(msg)
		}
	}
}

func (s *Server) createRoom(c port.ChatClient) {
	roomIdx, err := s.app.CreateRoom()
	if err != nil {
		s.logger.Printf("create room failed : %s", err)
	}
	s.app.Join(roomIdx, c)
	msg := &Message{&pb.MsgRes{RoomIdx: int32(roomIdx)}}
	s.app.SendMsg(msg)
}
