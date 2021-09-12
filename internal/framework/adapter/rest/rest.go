package rest

import (
	"log"

	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/rest/router"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/rest/server"
)

type Adapter struct {
}

func NewRestAdapter() *Adapter {
	return &Adapter{}
}

func (rest Adapter) Run() {
	r := router.New()
	httpServer := server.New(r, ":5000")
	log.Fatal(httpServer.ListenAndServe())
}
