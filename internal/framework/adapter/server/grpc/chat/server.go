package chat

import (
	"io"
	"log"
	"math/rand"
	"time"

	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/server/grpc/chat/pb"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/port"
)

const (
	CREATE_ROOM_REQ = 1
	JOIN_ROOM_REQ   = 2
	EXIT_ROOM_REQ   = 3
	TEXT_MSG_REQ    = 4
	IMAGE_MSG_REQ   = 5
)

type Server struct {
	wp      port.WorkerPool
	chatApp port.Chat
}

func NewServer(wp port.WorkerPool, chatApp port.Chat) *Server {
	return &Server{
		wp:      wp,
		chatApp: chatApp,
	}
}

//! --------------------- (1) grpc request ---------------------
func (s *Server) ChatService(stream pb.ChatService_ChatServiceServer) error {
	s.wp.RegisterJob(s.joinRoomJob(1, &Client{1, stream}))
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Printf("receiving message err: %s", err)
			return err
		}
		switch msg.Request {
		case CREATE_ROOM_REQ:
			s.wp.RegisterJob(s.createRoomJob(&Client{int(msg.UserIdx), stream}))
		case JOIN_ROOM_REQ:
			s.wp.RegisterJob(s.joinRoomJob(int(msg.RoomIdx), &Client{int(msg.UserIdx), stream}))
		case EXIT_ROOM_REQ:
			s.wp.RegisterJob(s.exitRoomJob(int(msg.RoomIdx), int(msg.UserIdx)))
		case TEXT_MSG_REQ:
			s.wp.RegisterJob(s.broadcastMsgJob(&pb.MsgRes{RoomIdx: msg.RoomIdx, UserIdx: msg.UserIdx, Body: msg.Body}))
		}
	}
}

//! ----------- 1) chat room -----------
func (s *Server) createRoomJob(c *Client) func() {
	return func() {
		roomIdx, err := s.chatApp.CreateRoom(c)
		if err != nil {
			log.Printf("create room error : %s", err) // TODO: need to send an error to client
			return
		}
		s.chatApp.BroadcastMsg(&pb.MsgRes{RoomIdx: int32(roomIdx)})
	}
}

func (s *Server) joinRoomJob(roomIdx int, c *Client) func() {
	return func() {
		time.Sleep(time.Duration(rand.Int31n(5)) * time.Microsecond)
		err := s.chatApp.JoinRoom(roomIdx, c)
		if err != nil {
			log.Printf("join room err : %s", err) // TODO: need to send an error to client
			return
		}
		s.chatApp.BroadcastMsg(&pb.MsgRes{RoomIdx: 1, UserIdx: 1})
	}
}

func (s *Server) exitRoomJob(roomIdx, userIdx int) func() {
	return func() {
		err := s.chatApp.ExitRoom(roomIdx, userIdx)
		if err != nil {
			log.Printf("exit room err : %s", err) // TODO: need to send an error to client
			return
		}
		s.chatApp.BroadcastMsg(&pb.MsgRes{RoomIdx: int32(roomIdx), UserIdx: int32(userIdx)})
	}
}

//! ----------- 2) broadcast -----------
func (s *Server) broadcastMsgJob(msg *pb.MsgRes) func() {
	return func() {
		s.chatApp.BroadcastMsg(msg)
	}
}
