package middlleware

import (
	"log"
	"net/http"

	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/server/rest/handler"
)

type Handler struct {
	logger *log.Logger
}

func New(logger *log.Logger) *Handler {
	return &Handler{
		logger: logger,
	}
}

func (m Handler) Register(r handler.Router) {
	r.Use(m.setHeader)
}

func (m *Handler) setHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		next.ServeHTTP(w, r)
	})
}
