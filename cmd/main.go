package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ByungHakNoh/hexagonal-microservice/core"
)

func main() {
	r := core.NewRouter()
	r.Get("/home", func(rw http.ResponseWriter, r *http.Request) {

	})
	r.Get("/test/:id/:test([0-9]+)", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println(core.URLParam(r))
	})
	server := core.NewServer(r, ":5000")
	log.Fatal(server.ListenAndServe())
}
