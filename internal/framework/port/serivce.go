package port

//! 1. Chat
type Chat interface {
	CreateRoom(roomIdxs *[]int) (int, error)
	JoinRoom(roomIdx int, client Client)
	ExitRoom(roomIdx, userIdx int)
	ExitAllRooms(roomIdxs *[]int, client Client)
	BroadcastMsg(msg Message)
}

//** (1) chat client
type Client interface {
	GetUserIdx() int
	SendMsg(msg interface{}) error
}

//* (2) message
type Message interface {
	GetRoomIdx() int
}

//!  2. User
type User interface {
	Register()
	Signin()
}
