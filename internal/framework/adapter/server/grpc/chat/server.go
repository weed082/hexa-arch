package chat

import (
	"io"
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/ByungHakNoh/hexagonal-microservice/internal/core/concurrency"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/server/grpc/chat/pb"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/port"
)

var m = sync.Mutex{}

const (
	CREATE_ROOM_REQ = 1
	JOIN_ROOM_REQ   = 2
	EXIT_ROOM_REQ   = 3
	TEXT_MSG_REQ    = 4
	IMAGE_MSG_REQ   = 5
)

type Server struct {
	wp          *concurrency.WorkerPool
	rooms       map[int][]port.Client
	chatApp     port.ChatApp
	roomPoolIdx int
	msgPoolIdx  int
}

func NewServer(wp *concurrency.WorkerPool, chatApp port.ChatApp) *Server {
	s := &Server{
		wp:          wp,
		chatApp:     chatApp,
		rooms:       make(map[int][]port.Client),
		roomPoolIdx: wp.AddPool(1),
		msgPoolIdx:  wp.AddPool(1),
	}
	log.Printf("room pool : %d", s.roomPoolIdx)
	log.Printf("msg pool : %d", s.msgPoolIdx)
	s.rooms[1] = []port.Client{}
	return s
}

//! --------------------- (1) grpc request ---------------------
func (s *Server) ChatService(stream pb.ChatService_ChatServiceServer) error {
	s.wp.RegisterJob(s.roomPoolIdx, concurrency.Job{Callback: s.joinRoomJob(1, Client{1, stream})})
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
			s.wp.RegisterJob(s.roomPoolIdx, concurrency.Job{Callback: s.createRoomJob(Client{int(msg.UserIdx), stream})})
		case JOIN_ROOM_REQ:
			s.wp.RegisterJob(s.roomPoolIdx, concurrency.Job{Callback: s.joinRoomJob(int(msg.RoomIdx), Client{int(msg.UserIdx), stream})})
		case EXIT_ROOM_REQ:
			s.wp.RegisterJob(s.roomPoolIdx, concurrency.Job{Callback: s.exitRoomJob(int(msg.RoomIdx), int(msg.UserIdx))})
			// s.wp.RegisterJobCallback(concurrency.Job{Callback: s.exitRoomJob(int(msg.RoomIdx), int(msg.UserIdx))})
		case TEXT_MSG_REQ:
			s.wp.RegisterJob(s.msgPoolIdx, concurrency.Job{Callback: s.broadcastMsgJob(&pb.MsgRes{RoomIdx: msg.RoomIdx, UserIdx: msg.UserIdx, Body: msg.Body})})
			// s.wp.RegisterJobCallback(concurrency.Job{Callback: s.broadcastMsgJob(&pb.MsgRes{RoomIdx: msg.RoomIdx, UserIdx: msg.UserIdx, Body: msg.Body})})
		}
	}
}

//! ----------- 1) chat room -----------
func (s *Server) createRoomJob(c Client) func() {
	return func() {
		roomIdx, err := s.chatApp.CreateRoom(c, s.rooms)
		if err != nil {
			log.Printf("create room error : %s", err) // TODO: need to send an error to client
			return
		}
		s.broadcastMsg(&pb.MsgRes{RoomIdx: int32(roomIdx)})
	}
}

func (s *Server) joinRoomJob(roomIdx int, c Client) func() {
	return func() {
		log.Printf("count: %v", len(s.rooms[1]))
		time.Sleep(time.Duration(rand.Int31n(5)) * time.Microsecond)
		m.Lock()
		err := s.chatApp.JoinRoom(roomIdx, c, s.rooms)
		m.Unlock()

		if err != nil {
			log.Printf("join room err : %s", err) // TODO: need to send an error to client
			return
		}
		// s.broadcastMsg(&pb.MsgRes{RoomIdx: 1, UserIdx: 1})
	}
}

func (s *Server) exitRoomJob(roomIdx, userIdx int) func() {
	return func() {
		err := s.chatApp.ExitRoom(roomIdx, userIdx, s.rooms)
		if err != nil {
			log.Printf("exit room err : %s", err) // TODO: need to send an error to client
			return
		}
		s.broadcastMsg(&pb.MsgRes{RoomIdx: int32(roomIdx), UserIdx: int32(userIdx)})
	}
}

//! ----------- 2) broadcast -----------
func (s *Server) broadcastMsg(msg *pb.MsgRes) {
	log.Println(len(s.rooms[1]))
	for _, c := range s.rooms[int(msg.RoomIdx)] {
		err := c.(Client).stream.Send(msg)
		if err != nil {
			log.Printf("sending message error: %s", err)
			continue
		}
	}
}

func (s *Server) broadcastMsgJob(msg *pb.MsgRes) func() {
	return func() {
		s.broadcastMsg(msg)
	}
}
