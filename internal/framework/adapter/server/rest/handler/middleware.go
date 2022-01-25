package handler

import (
	"log"
	"net/http"

	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/port"
)

type Middleware struct {
	logger *log.Logger
}

func NewServerHeader(logger *log.Logger) *Middleware {
	return &Middleware{
		logger: logger,
	}
}

func (m Middleware) Register(r port.Router) {
	r.Use(m.serverHeader)
}

func (m *Middleware) serverHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		next.ServeHTTP(w, r)
	})
}
