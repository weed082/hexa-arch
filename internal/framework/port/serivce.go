package port

//! ChatApp
type ChatApp interface {
	CreateRoom(client Client, rooms map[int][]Client) (int, error)
	RemoveRoom(roomIdx int, rooms map[int][]Client) error
	JoinRoom(roomIdx int, client Client, rooms map[int][]Client) error
	ExitRoom(roomIdx, userIdx int, rooms map[int][]Client) error
}

type Client interface {
	GetUserIdx() int
}

//!  User
type UserApp interface {
	Register()
	Signin()
}

type FileApp interface {
}
