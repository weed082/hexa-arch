package application

import (
	"log"

	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/port"
)

type User struct {
	logger *log.Logger
	repo   port.UserRepo
}

func NewUser(logger *log.Logger, repo port.UserRepo) *User {
	return &User{
		logger: logger,
		repo:   repo,
	}
}

func (app *User) Register() {

}

func (app *User) Signin() {

}
