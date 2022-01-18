package port

//! 1. Chat
type Chat interface {
	JoinRoom(rooms map[int][]Client, roomIdx int, client Client)
	ExitRoom(rooms map[int][]Client, roomIdx, userIdx int) error
	ExitAllRooms(rooms map[int][]Client, roomIdxs *[]int, client Client)
	BroadcastMsg(rooms map[int][]Client, msg Message)
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
