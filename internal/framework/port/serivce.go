package port

import "sync"

//! ChatApp
type ChatApp interface {
	CreateRoom(mtx *sync.Mutex, client Client, rooms map[int][]Client) (int, error)
	RemoveRoom(mtx *sync.Mutex, roomIdx int, rooms map[int][]Client) error
	JoinRoom(mtx *sync.Mutex, roomIdx int, client Client, rooms map[int][]Client) error
	ExitRoom(mtx *sync.Mutex, roomIdx, userIdx int, rooms map[int][]Client) error
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
