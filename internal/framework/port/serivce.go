package port

//! ChatApp
type ChatApp interface {
	CreateRoom(client interface{}, rooms map[int][]interface{}) (int, error)
	RemoveRoom(roomIdx int, rooms map[int][]interface{}) error
	JoinRoom(client interface{}, clients []interface{}) error
	ExitRoom(roomIdx, userIdx, removeIdx int, rooms map[int][]interface{}) error
}

//!  User
type UserApp interface {
	Register()
	Signin()
}

type FileApp interface {
}
