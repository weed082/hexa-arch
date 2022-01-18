package application

import (
	"errors"
	"log"

	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/port"
)

type Chat struct {
	logger *log.Logger
	repo   port.ChatRepo
}

func NewChat(logger *log.Logger, rooms map[int]port.Client, repo port.ChatRepo) *Chat {
	return &Chat{
		logger: logger,
		repo:   repo,
	}
}

//! ----------- 1) chat db related -----------

//! ----------- 2) chat function related -----------
//* (1) rooms
func (c *Chat) JoinRoom(rooms map[int][]port.Client, roomIdx int, client port.Client) {
	rooms[roomIdx] = append(rooms[roomIdx], client)
	c.logger.Println(len(rooms[roomIdx]))
}

func (c *Chat) ExitRoom(rooms map[int][]port.Client, roomIdx, userIdx int) error {
	clients := rooms[roomIdx]
	for i, client := range clients {
		if client.GetUserIdx() != userIdx {
			continue
		}
		rooms[roomIdx] = append(clients[:i], clients[i+1:]...)
		if len(rooms[roomIdx]) == 0 {
			delete(rooms, roomIdx)
		}
		c.logger.Printf("current clients : %d, current rooms : %d", len(rooms[roomIdx]), len(rooms))
		return nil
	}
	return errors.New("no match user idx in the chat room")
}

func (c *Chat) ExitAllRooms(rooms map[int][]port.Client, roomIdxs *[]int, client port.Client) {
	for _, roomIdx := range *roomIdxs {
		clients := rooms[roomIdx]
		for i, existClient := range clients {
			if existClient == client {
				rooms[roomIdx] = append(clients[:i], clients[i+1:]...) // delete user
				if len(rooms[roomIdx]) == 0 {
					delete(rooms, roomIdx) // delete room
				}
				c.logger.Printf("room len : %d, clients len : %d, index : %d", len(rooms), len(rooms[roomIdx]), i)
				break
			}
		}
	}
}

//* (2) message
func (c *Chat) BroadcastMsg(rooms map[int][]port.Client, msg port.Message) {
	roomIdx := msg.GetRoomIdx()

	for _, client := range rooms[roomIdx] {
		err := client.SendMsg(msg)
		if err != nil {
			c.logger.Printf("sending message error: %s", err)
			continue
		}
	}
}
