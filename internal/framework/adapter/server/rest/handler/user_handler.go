package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/server/rest/router"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/port"
)

type UserHandler struct {
	logger *log.Logger
	app    port.User
}

func NewUserHandler(logger *log.Logger, app port.User) *UserHandler {
	return &UserHandler{
		logger: logger,
		app:    app,
	}
}

func (handler *UserHandler) Register(r *router.Router) {
	r.Get("/register/:test", handler.test)
}

func (handler *UserHandler) test(w http.ResponseWriter, r *http.Request) {
	fmt.Println("yes")
}
