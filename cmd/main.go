package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ByungHakNoh/hexagonal-microservice/external"
)

func main() {
	r := external.NewRouter()
	r.Use(middlewareTest) // 미들웨어 사용하기
	r.Get("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(rw, "home")
	})
	r.Get("/test/:id([0-9]+)", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(rw, external.URLParam(r))
	})
	server := external.NewServer(r, ":5000")
	log.Fatal(server.ListenAndServe())
}

// 미들웨어 만들기
func middlewareTest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(rw, "middle")
		next.ServeHTTP(rw, r)
	})
}
