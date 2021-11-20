package application

import (
	"errors"
	"log"

	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/server/grpc/chat/pb"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/port"
)

type Chat struct {
	rooms map[int][]port.Client
	repo  port.ChatRepo
}

func NewChat(rooms map[int]port.Client, repo port.ChatRepo) *Chat {
	return &Chat{
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
	c.rooms[roomIdx] = []port.Client{client}
	log.Printf("room idx: %d, client count: %d", roomIdx, len(c.rooms))
	return roomIdx, nil
}

func (c *Chat) JoinRoom(roomIdx int, client port.Client) error {
	c.rooms[roomIdx] = append(c.rooms[roomIdx], client)
	return nil
}

func (c *Chat) ExitRoom(roomIdx, userIdx int) error {
	clients := c.rooms[roomIdx]
	for index, client := range clients {
		if client.GetUserIdx() != userIdx {
			continue
		}
		clients = append(clients[:index], clients[index+1:]...)
		if len(clients) == 0 {
			delete(c.rooms, roomIdx)
		}
		return nil
	}
	return errors.New("no match user idx in the chat room")
}

func (c *Chat) ExitAllRooms(roomIdxs []int, client port.Client) error {
	for _, roomIdx := range roomIdxs {
		clients := c.rooms[roomIdx]
		for index, existClient := range c.rooms[roomIdx] {
			if existClient == client {
				log.Printf("len : %d, index : %d", len(clients), index)
				c.rooms[roomIdx] = append(clients[:index], clients[index+1:]...) // delete user

				if len(c.rooms[roomIdx]) == 0 {
					delete(c.rooms, roomIdx) // delete room
				}
				break
			}
		}
	}

	return nil
}

//! ----------- 2) Msg -----------
func (c *Chat) BroadcastMsg(msg *pb.MsgRes) {
	for _, c := range c.rooms[int(msg.RoomIdx)] {
		err := c.SendMsg(msg)
		if err != nil {
			log.Printf("sending message error: %s", err)
			continue
		}
	}
}
