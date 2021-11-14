package chat

import (
	"io"
	"log"

	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/port"
)

const (
	CREATE_ROOM_REQ = 1
	JOIN_ROOM_REQ   = 2
	EXIT_ROOM_REQ   = 3
	TEXT_MSG_REQ    = 4
	IMAGE_MSG_REQ   = 5
)

type server struct {
	rooms    map[int][]interface{}
	msgChan  chan *MsgRes
	roomChan chan roomReq
	chatApp  port.ChatApp
}

type client struct {
	userIdx int
	stream  ChatService_ChatServiceServer
}

type roomReq struct {
	request int
	roomIdx int
	client  client
}

func NewServer(chatApp port.ChatApp) *server {
	server := &server{
		rooms:    make(map[int][]interface{}),
		msgChan:  make(chan *MsgRes),
		roomChan: make(chan roomReq),
		chatApp:  chatApp,
	}
	go server.handleMsg() // TODO: need to find a way to close it gracefully + make this as a worker pool
	return server
}

//! --------------------- (1) grpc request ---------------------
func (s *server) ChatService(stream ChatService_ChatServiceServer) error {
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
			s.roomChan <- roomReq{int(msg.Request), 0, client{int(msg.UserIdx), stream}}
		case JOIN_ROOM_REQ, EXIT_ROOM_REQ:
			s.roomChan <- roomReq{int(msg.Request), int(msg.RoomIdx), client{int(msg.UserIdx), stream}}
		case TEXT_MSG_REQ:
			s.msgChan <- &MsgRes{RoomIdx: msg.RoomIdx, UserIdx: msg.UserIdx, Body: msg.Body}
		}
	}
}

//! --------------------- (2) handleMsg ---------------------
func (s *server) handleMsg() {
	for {
		select {
		case roomReq := <-s.roomChan:
			roomIdx := roomReq.roomIdx
			c := roomReq.client
			switch roomReq.request {
			case CREATE_ROOM_REQ:
				s.createRoom(c)
			case JOIN_ROOM_REQ:
				s.joinRoom(roomIdx, c)
			case EXIT_ROOM_REQ:
				s.exitRoom(roomIdx, c.userIdx)
			}
		case msg := <-s.msgChan:
			s.broadcastMsg(msg)
		}
	}
}

//! ----------- 1) chat room -----------
func (s *server) createRoom(c client) {
	roomIdx, err := s.chatApp.CreateRoom(c, s.rooms)
	if err != nil {
		log.Printf("create room error : %s", err)
		return
	}
	s.msgChan <- &MsgRes{RoomIdx: int32(roomIdx)}
}

func (s *server) joinRoom(roomIdx int, c client) {
	err := s.chatApp.JoinRoom(c, s.rooms[roomIdx])
	if err != nil {
		log.Printf("join room err : %s", err)
		return
	}
	s.msgChan <- &MsgRes{RoomIdx: int32(roomIdx), UserIdx: int32(c.userIdx)}
}

func (s *server) exitRoom(roomIdx, userIdx int) {
	for index, participant := range s.rooms[roomIdx] {
		participant := participant.(client)
		if userIdx == participant.userIdx {
			s.chatApp.ExitRoom(roomIdx, userIdx, index, s.rooms)
			break
		}
	}
}

func (s *server) broadcastMsg(msg *MsgRes) {
	for _, c := range s.rooms[int(msg.RoomIdx)] {
		err := c.(client).stream.Send(msg)
		if err != nil {
			log.Printf("sending message error: %s", err)
			continue
		}
	}
}
