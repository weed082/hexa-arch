package handler

import (
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

}

func (handler UserHandler) register() {

}

func (handler UserHandler) signin() {

}
