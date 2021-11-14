package rest

import (
	"log"
	"net/http"

	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/server/rest/core/router"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/server/rest/core/server"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/server/rest/handler"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/port"
)

type Rest struct {
	Server  *http.Server
	userApp port.UserApp
}

func NewRestAdapter(userApp port.UserApp) *Rest {
	return &Rest{userApp: userApp}
}

func (r *Rest) Run(port string) {
	router := router.New()
	handler.NewUserHandler(r.userApp).Register(router)

	r.Server = server.New(router, ":"+port)
	log.Println(r.Server.ListenAndServe())
}
