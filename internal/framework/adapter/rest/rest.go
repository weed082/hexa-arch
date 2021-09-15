package rest

import (
	"log"

	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/rest/router"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/rest/server"
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
	httpServer := server.New(r, ":8080")
	log.Fatal(httpServer.ListenAndServe())
}
