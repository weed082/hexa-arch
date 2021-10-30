package handler

import (
	"net/http"

	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/server/rest/core/router"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/port"
)

type UserHandler struct {
	app port.UserApp
}

func NewUserHandler(app port.UserApp) *UserHandler {
	return &UserHandler{app}
}

func (handler *UserHandler) Register(r *router.Router) {
	r.Get("/register", handler.test)
}

func (handler *UserHandler) test(w http.ResponseWriter, r *http.Request) {
}

func (handler *UserHandler) register(w http.ResponseWriter, r *http.Request) {

}

func (handler *UserHandler) signin(w http.ResponseWriter, r *http.Request) {

}
