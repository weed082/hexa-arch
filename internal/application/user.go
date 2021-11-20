package application

import "github.com/ByungHakNoh/hexagonal-microservice/internal/framework/port"

type User struct {
	repo port.UserRepo
}

func NewUser(repo port.UserRepo) *User {
	return &User{repo}
}

func (app *User) Register() {

}

func (app *User) Signin() {

}
