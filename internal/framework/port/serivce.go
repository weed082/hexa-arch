package port

//! ChatApp
type Client interface {
	GetUserIdx() int
}

type ChatApp interface {
	CreateRoom(client Client, rooms map[int][]Client) (int, error)
	RemoveRoom(roomIdx int, rooms map[int][]Client) error
	JoinRoom(client Client, clients []Client) error
	ExitRoom(roomIdx, userIdx int, rooms map[int][]Client) error
}

//!  User
type UserApp interface {
	Register()
	Signin()
}

type FileApp interface {
}
