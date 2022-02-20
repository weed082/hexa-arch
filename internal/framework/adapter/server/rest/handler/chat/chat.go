package chat

import (
	"log"
	"net/http"

	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/server/rest/handler"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/port"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/model"
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
	wsUpgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	// check updgrader
	wsUpgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	// upgrade to ws conn
	conn, err := wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		h.logger.Printf("ws connect failed: %s", err)
	}
	// read, write msg through ws
	for {
    var msg model.Msg
		err := conn.ReadJSON(&msg)
		if err != nil {
			h.logger.Printf("ws read msg failed: %s", err)
			break
		}
		if err != nil {
			h.logger.Printf("ws write msg failed: %s", err)
			break
		}
		h.logger.Printf("id: %v, name: %s", msg.Id, msg.Name)
	}
}

func (h *Chat) sendMsg() {

}
