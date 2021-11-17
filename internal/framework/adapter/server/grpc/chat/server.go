package chat

import (
	"io"
	"log"
	"math/rand"
	"time"

	"github.com/ByungHakNoh/hexagonal-microservice/internal/core/concurrency"
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
	wp      *concurrency.WorkerPool
	rooms   map[int][]port.Client
	chatApp port.ChatApp
}

func NewServer(wp *concurrency.WorkerPool, chatApp port.ChatApp) *Server {
	server := &Server{
		wp:      wp,
		chatApp: chatApp,
		rooms:   make(map[int][]port.Client),
	}
	server.rooms[1] = []port.Client{}
	return server
}

//! --------------------- (1) grpc request ---------------------
func (s *Server) ChatService(stream pb.ChatService_ChatServiceServer) error {
	s.wp.RegisterJobCallback(concurrency.Job{Callback: s.joinRoom(1, Client{1, stream})})
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
			s.wp.RegisterJobCallback(concurrency.Job{Callback: s.createRoom(Client{int(msg.UserIdx), stream})})
		case JOIN_ROOM_REQ:
			s.wp.RegisterJobCallback(concurrency.Job{Callback: s.joinRoom(int(msg.RoomIdx), Client{int(msg.UserIdx), stream})})
		case EXIT_ROOM_REQ:
			s.wp.RegisterJobCallback(concurrency.Job{Callback: s.exitRoom(int(msg.RoomIdx), int(msg.UserIdx))})
		case TEXT_MSG_REQ:
			s.wp.RegisterJobCallback(concurrency.Job{Callback: s.broadcastMsg(&pb.MsgRes{RoomIdx: msg.RoomIdx, UserIdx: msg.UserIdx, Body: msg.Body})})
		}
	}
}

//! ----------- 1) chat room -----------
func (s *Server) createRoom(c Client) func() {
	return func() {
		roomIdx, err := s.chatApp.CreateRoom(c, s.rooms)
		if err != nil {
			log.Printf("create room error : %s", err) // TODO: need to send an error to client
			return
		}
		s.wp.RegisterJobCallback(concurrency.Job{Callback: s.broadcastMsg(&pb.MsgRes{RoomIdx: int32(roomIdx)})})
	}
}

func (s *Server) joinRoom(roomIdx int, c Client) func() {
	return func() {
		m := rand.Intn(4)
		time.Sleep(time.Duration(m) * time.Second)
		err := s.chatApp.JoinRoom(roomIdx, c, s.rooms)
		if err != nil {
			log.Printf("join room err : %s", err) // TODO: need to send an error to client
			return
		}
		// s.wp.RegisterJobCallback(concurrency.Job{Callback: s.broadcastMsg(&pb.MsgRes{RoomIdx: int32(roomIdx), UserIdx: int32(c.userIdx)})})
	}
}

func (s *Server) exitRoom(roomIdx, userIdx int) func() {
	return func() {
		err := s.chatApp.ExitRoom(roomIdx, userIdx, s.rooms)
		if err != nil {
			log.Printf("exit room err : %s", err) // TODO: need to send an error to client
			return
		}
		s.wp.RegisterJobCallback(concurrency.Job{Callback: s.broadcastMsg(&pb.MsgRes{RoomIdx: int32(roomIdx), UserIdx: int32(userIdx)})})
	}
}

//! ----------- 2) broadcast -----------
func (s *Server) broadcastMsg(msg *pb.MsgRes) func() {
	return func() {
		for _, c := range s.rooms[int(msg.RoomIdx)] {
			err := c.(Client).stream.Send(msg)
			if err != nil {
				log.Printf("sending message error: %s", err)
				continue
			}
		}
	}
}
