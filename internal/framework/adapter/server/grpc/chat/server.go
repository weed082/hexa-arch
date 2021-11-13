package chat

import (
	"io"
	"log"

	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/port"
)

const (
	CREATE_ROOM_REQ = 1
	REMOVE_ROOM_REQ = 2
	JOIN_ROOM_REQ   = 3
	EXIT_ROOM_REQ   = 4
	TEXT_MSG_REQ    = 5
	IMAGE_MSG_REQ   = 6
)

type server struct {
	rooms    map[int][]interface{}
	msgChan  chan *Message
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
		msgChan:  make(chan *Message),
		roomChan: make(chan roomReq),
		chatApp:  chatApp,
	}
	go server.work() // TODO: need to find a way to close it gracefully + make this as a worker pool
	return server
}

//! --------------------- (1) chat work ---------------------
func (s *server) work() {
	for {
		select {
		case roomReq := <-s.roomChan:
			roomIdx := roomReq.roomIdx
			c := roomReq.client
			switch roomReq.request {
			case CREATE_ROOM_REQ:
				s.chatApp.CreateRoom(s.rooms, c) // create room and add stream to rooms
			case REMOVE_ROOM_REQ:
				s.chatApp.RemoveRoom(roomIdx, s.rooms) // create room and add stream to rooms
			case JOIN_ROOM_REQ:
				s.chatApp.JoinRoom(s.rooms[roomIdx], c) // create room and add stream to rooms
			case EXIT_ROOM_REQ:
				for index, participant := range s.rooms[roomIdx] {
					participant := participant.(client)
					if c.userIdx == participant.userIdx {
						s.chatApp.ExitRoom(roomIdx, c.userIdx, s.rooms, index)
						break
					}
				}
			}
		case msg := <-s.msgChan:
			s.sendMessage(msg)
		}
	}
}

// send message to
func (s *server) sendMessage(msg *Message) {
	for _, c := range s.rooms[int(msg.RoomIdx)] {
		err := c.(ChatService_ChatServiceServer).Send(msg)
		if err != nil {
			log.Printf("sending message error: %s", err)
			continue
		}
	}
}

//! --------------------- (2) grpc request ---------------------
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
		case REMOVE_ROOM_REQ, JOIN_ROOM_REQ, EXIT_ROOM_REQ:
			s.roomChan <- roomReq{int(msg.Request), int(msg.RoomIdx), client{int(msg.UserIdx), stream}}
		case TEXT_MSG_REQ:
			s.msgChan <- msg
		}
	}
}
