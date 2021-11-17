package application

import (
	"errors"
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

func (chatApp *ChatApp) CreateRoom(client port.Client, rooms map[int][]port.Client) (int, error) {
	roomIdx, err := chatApp.repo.CreateRoom()
	if err != nil {
		return 0, err
	}
	rooms[roomIdx] = []port.Client{client}
	log.Printf("room idx: %d, client count: %d", roomIdx, len(rooms))
	return roomIdx, nil
}

func (chatApp *ChatApp) RemoveRoom(roomIdx int, rooms map[int][]port.Client) error {
	return nil
}

func (chatApp *ChatApp) JoinRoom(roomIdx int, client port.Client, rooms map[int][]port.Client) error {
	rooms[roomIdx] = append(rooms[roomIdx], client)
	log.Printf("client count: %d", len(rooms[roomIdx]))
	return nil
}

func (chatApp *ChatApp) ExitRoom(roomIdx, userIdx int, rooms map[int][]port.Client) error {
	clients := rooms[roomIdx]
	for index, client := range clients {
		if client.GetUserIdx() != userIdx {
			continue
		}
		clients = append(clients[:index], clients[index+1:]...)
		if len(clients) == 0 {
			delete(rooms, roomIdx)
		}
		return nil
	}
	return errors.New("no match user idx in the chat room")
}
