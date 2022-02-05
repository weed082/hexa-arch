package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/port"
)

type Chat struct {
	logger *log.Logger
	app    port.Chat
}

func NewChat(logger *log.Logger, app port.Chat) *Chat {
	return &Chat{
		logger: logger,
		app:    app,
	}
}

func (h *Chat) Register(r port.Router) {
	r.Get("/chat", h.connect)
}

func (h *Chat) connect(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "yes")
}
