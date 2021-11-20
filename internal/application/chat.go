package application

import (
	"errors"
	"log"

	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/port"
)

type Chat struct {
	logger *log.Logger
	rooms  map[int][]port.Client
	repo   port.ChatRepo
}

func NewChat(logger *log.Logger, rooms map[int]port.Client, repo port.ChatRepo) *Chat {
	return &Chat{
		logger: logger,
		rooms:  map[int][]port.Client{},
		repo:   repo,
	}
}

//! ----------- 1) Room -----------
func (c *Chat) CreateRoom(client port.Client) (int, error) {
	roomIdx, err := c.repo.CreateRoom()
	if err != nil {
		return 0, err
	}
	c.rooms[roomIdx] = []port.Client{client}
	c.logger.Printf("room idx: %d, client count: %d", roomIdx, len(c.rooms))
	return roomIdx, nil
}

func (c *Chat) JoinRoom(roomIdx int, client port.Client) error {
	c.rooms[roomIdx] = append(c.rooms[roomIdx], client)
	c.logger.Println(len(c.rooms[roomIdx]))
	return nil
}

func (c *Chat) ExitRoom(roomIdx, userIdx int) error {
	clients := c.rooms[roomIdx]
	for i, client := range clients {
		if client.GetUserIdx() != userIdx {
			continue
		}
		c.rooms[roomIdx] = append(clients[:i], clients[i+1:]...)
		if len(c.rooms[roomIdx]) == 0 {
			delete(c.rooms, roomIdx)
		}
		c.logger.Printf("current clients : %d, current rooms : %d", len(c.rooms[roomIdx]), len(c.rooms))
		return nil
	}
	return errors.New("no match user idx in the chat room")
}

func (c *Chat) ExitAllRooms(roomIdxs *[]int, client port.Client) {
	for _, roomIdx := range *roomIdxs {
		clients := c.rooms[roomIdx]
		for i, existClient := range clients {
			if existClient == client {
				c.rooms[roomIdx] = append(clients[:i], clients[i+1:]...) // delete user
				if len(c.rooms[roomIdx]) == 0 {
					delete(c.rooms, roomIdx) // delete room
				}
				c.logger.Printf("room len : %d, clients len : %d, index : %d", len(c.rooms), len(c.rooms[roomIdx]), i)
				break
			}
		}
	}
}

//! ----------- 2) Msg -----------
func (c *Chat) BroadcastMsg(msg port.Message) {
	for _, client := range c.rooms[msg.GetRoomIdx()] {
		err := client.SendMsg(msg)
		if err != nil {
			c.logger.Printf("sending message error: %s", err)
			continue
		}
	}
}
