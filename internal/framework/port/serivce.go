package port

import "github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/server/grpc/chat/pb"

//! 1. Chat
type Chat interface {
	CreateRoom(client Client) (int, error)
	JoinRoom(roomIdx int, client Client) error
	ExitRoom(roomIdx, userIdx int) error
	BroadcastMsg(*pb.MsgRes)
}

//** (1) chat client
type Client interface {
	GetUserIdx() int
	SendMsg(msg *pb.MsgRes) error
}

//!  2. User
type User interface {
	Register()
	Signin()
}
