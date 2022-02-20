package port

//! 1. Chat
type Chat interface {
	CreateRoom() (int, error)
	Join(roomIdx int, client ChatClient)
	Exit(roomIdx, userIdx int)
	ConnectAll(client ChatClient)
	DisconnectAll(client ChatClient)
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
