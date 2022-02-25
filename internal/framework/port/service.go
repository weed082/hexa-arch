package port

//! 1. Chat
type Chat interface {
	CreateRoom() (int, error)
	JoinRoom(roomIdx int, client ChatClient) error
	ExitRoom(roomIdx, userIdx int) error
	ConnectRoom(roomIdx int, client ChatClient)
	ConnectRooms(rooms []int, client ChatClient)
  DisconnectRoom(roomIdx, userIdx int)
	DisconnectRooms(client ChatClient)
	SendMsg(msg ChatMsg)
}

//* (1) chat client
type ChatClient interface {
	GetUserIdx() int
	SendMsg(msg interface{}) error
}

//* (2) message
type ChatMsg interface {
	GetRoomIdx() int
}

//!  2. User
type User interface {
	Register()
	Signin()
}
