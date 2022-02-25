package port

//! 1. Chat
type Chat interface {
	CreateRoom() (int, error)
	Connect(client ChatClient)
  Disconnect(client ChatClient)
	JoinRoom(roomIdx int, client ChatClient) error
	ExitRoom(roomIdx, userIdx int) error
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
