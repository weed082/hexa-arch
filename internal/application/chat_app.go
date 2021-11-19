package application

import (
	"errors"
	"log"
	"sync"

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

func (chatApp *ChatApp) CreateRoom(mtx *sync.Mutex, client port.Client, rooms map[int][]port.Client) (int, error) {
	roomIdx, err := chatApp.repo.CreateRoom()
	if err != nil {
		return 0, err
	}
	mtx.Lock()
	rooms[roomIdx] = []port.Client{client}
	mtx.Unlock()
	log.Printf("room idx: %d, client count: %d", roomIdx, len(rooms))
	return roomIdx, nil
}

func (chatApp *ChatApp) RemoveRoom(mtx *sync.Mutex, roomIdx int, rooms map[int][]port.Client) error {
	return nil
}

func (chatApp *ChatApp) JoinRoom(mtx *sync.Mutex, roomIdx int, client port.Client, rooms map[int][]port.Client) error {
	mtx.Lock()
	rooms[roomIdx] = append(rooms[roomIdx], client)
	mtx.Unlock()
	log.Printf("client count: %d", len(rooms[roomIdx]))
	return nil
}

func (chatApp *ChatApp) ExitRoom(mtx *sync.Mutex, roomIdx, userIdx int, rooms map[int][]port.Client) error {
	clients := rooms[roomIdx]
	for index, client := range clients {
		if client.GetUserIdx() != userIdx {
			continue
		}
		mtx.Lock()
		clients = append(clients[:index], clients[index+1:]...)
		if len(clients) == 0 {
			delete(rooms, roomIdx)
		}
		mtx.Unlock()
		return nil
	}
	return errors.New("no match user idx in the chat room")
}
