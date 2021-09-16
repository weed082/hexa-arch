package rest

import (
	"log"

	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/server/rest/core/router"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/server/rest/core/server"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/server/rest/handler"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/port"
)

type Rest struct {
	userApp port.UserApp
}

func NewRestAdapter(userApp port.UserApp) *Rest {
	return &Rest{userApp: userApp}
}

func (rest Rest) Run() {
	r := router.New()
	handler.NewUserHandler(rest.userApp).Register(r)
	httpServer := server.New(r, ":8080")
	log.Fatal(httpServer.ListenAndServe())
}
