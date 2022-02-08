package chat

import (
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
	// c := &Client{stream: stream}
	// roomIdxs := &[]int{}
	// defer s.app.ExitAllRooms(s.rooms, roomIdxs, c)

	// for {
	// 	msg, err := stream.Recv()
	// 	if err == io.EOF {
	// 		return nil
	// 	}
	// 	if err != nil {
	// 		s.logger.Printf("receiving message err: %s", err)
	// 		return err
	// 	}
	// 	if c.userIdx == 0 {
	// 		c.userIdx = int(msg.UserIdx)
	// 	}
	// 	switch msg.Request {
	// 	case CREATE_ROOM_REQ:
	// 		s.app.createRoomJob(roomIdxs, c)) // TODO: get roomIdx after inserting room config in db
	// 	case JOIN_ROOM_REQ:
	// 		s.chatPool.RegisterJob(s.joinRoomJob(roomIdxs, int(msg.RoomIdx), c))
	// 	case EXIT_ROOM_REQ:
	// 		s.chatPool.RegisterJob(s.exitRoomJob(int(msg.RoomIdx), int(msg.UserIdx)))
	// 	case TEXT_MSG_REQ:
	// 		s.chatPool.RegisterJob(s.broadcastMsgJob(&pb.MsgRes{RoomIdx: msg.RoomIdx, UserIdx: msg.UserIdx, Body: msg.Body}))
	// 	}
	// }
	return nil
}

//! ----------- 1) chat room -----------
// func (s *Server) createRoomJob(roomIdxs *[]int, c *Client) func() {
// 	roomIdx := 1 // TODO: get roomIdx from db
// 	*roomIdxs = append(*roomIdxs, roomIdx)
// 	return func() {
// 		s.app.JoinRoom(s.rooms, roomIdx, c)
// 		msg := &Message{&pb.MsgRes{RoomIdx: int32(roomIdx)}}
// 		s.app.BroadcastMsg(s.rooms, msg)
// 	}
// }

// func (s *Server) joinRoomJob(roomIdxs *[]int, roomIdx int, c *Client) func() {
// 	*roomIdxs = append(*roomIdxs, roomIdx) // TODO: put this logic right after db insert
// 	return func() {
// 		s.app.JoinRoom(s.rooms, roomIdx, c)
// 		msg := &Message{&pb.MsgRes{RoomIdx: int32(roomIdx), UserIdx: int32(c.userIdx)}}
// 		s.app.BroadcastMsg(s.rooms, msg)
// 	}
// }

// func (s *Server) exitRoomJob(roomIdx, userIdx int) func() {
// 	return func() {
// 		err := s.app.ExitRoom(s.rooms, roomIdx, userIdx)
// 		if err != nil {
// 			s.logger.Printf("exit room err : %s", err) // TODO: send an error to client
// 			return
// 		}
// 		msg := &Message{&pb.MsgRes{RoomIdx: int32(roomIdx), UserIdx: int32(userIdx)}}
// 		s.app.BroadcastMsg(s.rooms, msg)
// 	}
// }

// func (s *Server) exitAllRooms(roomIdx *[]int, c *Client) func() {
// 	return func() {
// 		s.app.ExitAllRooms(s.rooms, roomIdx, c)
// 	}
// }

// //! ----------- 2) broadcast -----------
// func (s *Server) broadcastMsgJob(msg *pb.MsgRes) func() {
// 	return func() {
// 		s.app.BroadcastMsg(s.rooms, &Message{msg})
// 	}
// }
