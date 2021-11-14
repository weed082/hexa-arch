package port

import (
	"html/template"
)

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

type TemplateApp interface {
	RenderPage() (*template.Template, interface{}, error)
}

type FileApp interface {
}
