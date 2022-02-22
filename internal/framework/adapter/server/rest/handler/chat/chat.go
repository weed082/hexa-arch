package chat

import (
	"log"
	"net/http"

	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/server/rest/handler"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/port"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/model/chat"
	"github.com/gorilla/websocket"
)

const (
	CONNECT_REQ     = 1
	DISCONNECT_REQ  = 2
	CREATE_ROOM_REQ = 3
	JOIN_ROOM_REQ   = 4
	EXIT_ROOM_REQ   = 5
	MSG_REQ         = 6
)

const (
	TEXT_MSG  = 1
	IMG_MSG   = 2
	FILE_MSG  = 3
	VIDEO_MSG = 4
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
	for {
		// read, write req through ws
		var req chat.Req
		err := conn.ReadJSON(&req)
		if err != nil {
			h.logger.Printf("ws read msg failed: %s", err)
			break
		}
		h.logger.Println(req) // TODO: test
		var client *Client
		switch req.Type {
		case CONNECT_REQ:
			// client = &Client{int(req.UserIdx), conn}
			// h.app.ConnectAll(client)
		case DISCONNECT_REQ:
			h.app.DisconnectAll(client)
		case CREATE_ROOM_REQ:
			roomIdx, err := h.app.CreateRoom()
			if err != nil {
				h.logger.Printf("create room failed: %s", err)
				continue
			}
			h.app.Join(roomIdx, client)
		case JOIN_ROOM_REQ:
			// h.app.Join(int(req.RoomIdx), client)
		case EXIT_ROOM_REQ:
			// h.app.Exit(int(req.RoomIdx), client.GetUserIdx())
		case MSG_REQ:
			// h.app.SendMsg(&req)
		}
	}
}
