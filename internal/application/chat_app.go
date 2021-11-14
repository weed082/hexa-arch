package application

import (
	"log"

	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/port"
)

type ChatApp struct {
	repo port.ChatRepo
}

func NewChatApp(repo port.ChatRepo) *ChatApp {
	return &ChatApp{
		repo: repo,
	}
}

func (chatApp *ChatApp) CreateRoom(client interface{}, rooms map[int][]interface{}) (int, error) {
	roomIdx, err := chatApp.repo.CreateRoom()
	if err != nil {
		return 0, err
	}
	rooms[roomIdx] = []interface{}{client}
	log.Printf("room idx: %d, client count: %d", roomIdx, len(rooms))
	return roomIdx, nil
}

func (chatApp *ChatApp) RemoveRoom(roomIdx int, rooms map[int][]interface{}) error {
	return nil
}

func (chatApp *ChatApp) JoinRoom(client interface{}, clients []interface{}) error {
	clients = append(clients, client)
	log.Printf("client count: %d", len(clients))
	return nil
}

func (chatApp *ChatApp) ExitRoom(roomIdx, userIdx, removeIdx int, rooms map[int][]interface{}) error {
	clients := rooms[roomIdx]
	clients = append(clients[:removeIdx], clients[removeIdx+1:]...)
	if len(clients) == 0 {
		delete(rooms, roomIdx)
	}
	log.Printf("exit user : %d, room count: %d", userIdx, len(rooms))
	return nil
}
