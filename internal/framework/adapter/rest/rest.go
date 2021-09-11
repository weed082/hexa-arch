package rest

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/rest/router"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/rest/server"
)

type Adapter struct {
}

func NewAdapter() *Adapter {
	return &Adapter{}
}

func (rest Adapter) Run() {
	r := router.New()
	r.Get("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(rw, "home")
	})
	r.Get("/test/:id([0-9]+)", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(rw, router.URLParam(r))
	})
	server := server.New(r, ":5000")
	log.Fatal(server.ListenAndServe())
}
