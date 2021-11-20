package application

import (
	"errors"
	"log"
	"sync"

	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/server/grpc/chat/pb"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/port"
)

type Chat struct {
	mtx   *sync.RWMutex
	rooms map[int][]port.Client
	repo  port.ChatRepo
}

func NewChat(mtx *sync.RWMutex, rooms map[int]port.Client, repo port.ChatRepo) *Chat {
	return &Chat{
		mtx:   mtx,
		rooms: map[int][]port.Client{},
		repo:  repo,
	}
}

//! ----------- 1) Room -----------
func (c *Chat) CreateRoom(client port.Client) (int, error) {
	roomIdx, err := c.repo.CreateRoom()
	if err != nil {
		return 0, err
	}
	c.mtx.Lock()
	defer c.mtx.Unlock()
	c.rooms[roomIdx] = []port.Client{client}
	log.Printf("room idx: %d, client count: %d", roomIdx, len(c.rooms))
	return roomIdx, nil
}

func (c *Chat) JoinRoom(roomIdx int, client port.Client) error {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	c.rooms[roomIdx] = append(c.rooms[roomIdx], client)
	return nil
}

func (c *Chat) ExitRoom(roomIdx, userIdx int) error {
	clients := c.rooms[roomIdx]
	for index, client := range clients {
		if client.GetUserIdx() != userIdx {
			continue
		}
		c.mtx.Lock()
		defer c.mtx.Unlock()
		clients = append(clients[:index], clients[index+1:]...)
		if len(clients) == 0 {
			delete(c.rooms, roomIdx)
		}
		return nil
	}
	return errors.New("no match user idx in the chat room")
}

func (c *Chat) ExitAllRooms(client port.Client) error {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	clients := c.rooms[1]
	for index, existClient := range clients {
		if existClient == client {
			log.Println("erased")
			clients = append(clients[:index], clients[index+1:]...)
			if len(clients) == 0 {
				delete(c.rooms, 1)
			}
			log.Println(c.rooms)
		}
	}
	return nil
}

//! ----------- 2) Msg -----------
func (c *Chat) BroadcastMsg(msg *pb.MsgRes) {
	c.mtx.RLock()
	defer c.mtx.RUnlock()
	for _, c := range c.rooms[int(msg.RoomIdx)] {
		err := c.SendMsg(msg)
		if err != nil {
			log.Printf("sending message error: %s", err)
			continue
		}
	}
}