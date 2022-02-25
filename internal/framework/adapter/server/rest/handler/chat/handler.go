package chat

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/server/rest/handler"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/port"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/model/chat"
	"github.com/gorilla/websocket"
	"github.com/mitchellh/mapstructure"
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
	r.Get("/chat", h.chat)
}

//* ws connection
func (h *Chat) chat(w http.ResponseWriter, r *http.Request) {
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
		return
	}
	defer conn.Close()

	var client *Client
	var req chat.Req
	for {
		// read, write req through ws
		err := conn.ReadJSON(&req)
		if err != nil {
			h.logger.Printf("ws read msg failed: %s", err)
		}
		switch req.Type {
		case CONNECT_REQ:
			if client = h.connect(req.Body, conn); client == nil {
				break
			}
		case DISCONNECT_REQ:
			h.app.Disconnect(client)
		case CREATE_ROOM_REQ:
			h.createRoom(client)
		case JOIN_ROOM_REQ:
			h.joinRoom(req.Body, client)
		case EXIT_ROOM_REQ:
			// h.app.Exit(int(req.RoomIdx), client.GetUserIdx())
		case MSG_REQ:
			// h.app.SendMsg(&req)
		}
	}
}

// ws connect
func (h *Chat) connect(body interface{}, conn *websocket.Conn) *Client {
	reqData := &struct {
		userIdx int    `mapstruct:"userIdx"`
		name    string `mapstruct:"name"`
	}{}
	if mapstructure.Decode(body, reqData) != nil {
		conn.WriteJSON(&chat.Res{Code: 400, Body: "wrong request body"})
		return nil
	}
	client := &Client{reqData.userIdx, reqData.name, conn}
	h.app.Connect(client)                       // connect to rooms that client was participated in
	h.app.SendRes(client, &chat.Res{Code: 200}) // send success msg to client
	return client
}

// create chat room
func (h *Chat) createRoom(client *Client) {
	roomIdx, err := h.app.CreateRoom()
	if err != nil {
		h.logger.Printf("create room failed: %s", err)
		h.app.SendRes(client, &chat.Res{Code: 500, Body: "create room failed"})
		return
	}
	h.app.JoinRoom(roomIdx, client)
}

// join chat room
func (h *Chat) joinRoom(body interface{}, client *Client) {
	reqData := &struct {
		roomIdx int
	}{}
	if mapstructure.Decode(body, reqData) != nil {
		h.app.SendRes(client, &chat.Res{Code: 400, Body: "wrong request body"})
		return
	}
	err := h.app.JoinRoom(reqData.roomIdx, client)
	if err != nil {
		h.logger.Printf("join room failed: %s", err)
		h.app.SendRes(client, &chat.Res{Code: 500, Body: "join room failed"})
		return
	}
}

// exit room
func (h *Chat) exitRoom(body interface{}, client *Client) {
	reqData := &struct {
		roomIdx int
	}{}
	if mapstructure.Decode(body, reqData) != nil {
		h.app.SendRes(client, &chat.Res{Code: 400, Body: "wrong reqest body"})
		return
	}

	err := h.app.ExitRoom(reqData.roomIdx, client)
	if err != nil {
		h.app.SendRes(client, &chat.Res{Code: 400, Body: fmt.Sprintf("exit room failed: %s", err)})
		return
	}
	h.app.SendRes(client, &chat.Res{Code: 200, Body: "success"})
}
