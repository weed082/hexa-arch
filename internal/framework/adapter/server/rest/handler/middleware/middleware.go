package middlleware

import (
	"log"
	"net/http"

	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/server/rest/handler"
)

type Middleware struct {
	logger *log.Logger
}

func NewServerHeader(logger *log.Logger) *Middleware {
	return &Middleware{
		logger: logger,
	}
}

func (m Middleware) Register(r handler.Router) {
	r.Use(m.setHeader)
}

func (m *Middleware) setHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		next.ServeHTTP(w, r)
	})
}
