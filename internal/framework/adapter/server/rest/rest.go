package rest

import (
	"log"

	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/rest/router"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/rest/server"
)

type Rest struct {
}

func NewRestAdapter() *Rest {
	return &Rest{}
}

func (rest Rest) Run() {
	r := router.New()
	httpServer := server.New(r, ":8080")
	log.Fatal(httpServer.ListenAndServe())
}
