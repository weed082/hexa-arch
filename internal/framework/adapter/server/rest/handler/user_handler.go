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

func (handler UserHandler) Register(r *router.Router) {
	r.Get("/", handler.test)
}

func (handler UserHandler) test(rw http.ResponseWriter, r *http.Request) {
	// template, _ :=  template.HTML(`<div>working</div>`)
}

func (handler UserHandler) register(rw http.ResponseWriter, r *http.Request) {

}

func (handler UserHandler) signin(rw http.ResponseWriter, r *http.Request) {

}
