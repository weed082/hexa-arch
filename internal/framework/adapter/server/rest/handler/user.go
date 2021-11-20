package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/port"
)

type User struct {
	logger *log.Logger
	app    port.User
}

func NewUser(logger *log.Logger, app port.User) *User {
	return &User{
		logger: logger,
		app:    app,
	}
}

func (handler *User) Register(r port.Router) {
	r.Get("/register/:test", handler.test)
}

func (handler *User) test(w http.ResponseWriter, r *http.Request) {
	fmt.Println("yes")
}
