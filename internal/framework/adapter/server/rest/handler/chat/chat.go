package chat

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/server/rest/handler"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/port"
	"github.com/gorilla/websocket"
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

func (h *Chat) Register(r handler.Router) {
	r.Get("/chat", h.connect)
}

//* ws connection
func (h *Chat) connect(w http.ResponseWriter, r *http.Request) {
	up := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	// upgrade to ws conn
	conn, err := up.Upgrade(w, r, nil)
	if err != nil {
		h.logger.Printf("ws connect failed: %s", err)
	}
	// read, write msg through ws
	for {
		_, p, err := conn.ReadMessage()
		fmt.Println(string(p))
		if err != nil {
			h.logger.Printf("ws read msg failed: %s", err)
			break
		}
		if err != nil {
			h.logger.Printf("ws write msg failed: %s", err)
			break
		}
	}
}
