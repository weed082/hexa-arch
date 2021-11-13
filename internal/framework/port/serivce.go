package port

import (
	"html/template"
)

//! ChatApp
type ChatApp interface {
	CreateRoom(rooms map[int][]interface{}, client interface{}) (int, error)
	RemoveRoom(roomIdx int, rooms map[int][]interface{}) error
	JoinRoom(clients []interface{}, client interface{}) error
	ExitRoom(roomIdx, userIdx int, rooms map[int][]interface{}, removeIdx int) error
}

//!  User
type UserApp interface {
	Register()
	Signin()
}

type TemplateApp interface {
	RenderPage() (*template.Template, interface{}, error)
}

type FileApp interface {
}
