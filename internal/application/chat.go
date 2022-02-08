package application

import (
	"log"

	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/port"
)

type Chat struct {
	logger *log.Logger
	pool   port.WorkerPool
	repo   port.ChatRepo
	rooms  map[int][]port.Client
}

func NewChat(logger *log.Logger, rooms map[int]port.Client, pool port.WorkerPool, repo port.ChatRepo) *Chat {
	pool.Generate(1) // only need single worker for sync
	return &Chat{
		logger: logger,
		pool:   pool,
		repo:   repo,
		rooms:  make(map[int][]port.Client),
	}
}

//! ----------- 1) chat db related -----------

//! ----------- 2) chat function related -----------
//* (1) rooms

func (c *Chat) CreateRoom() {

}

func (c *Chat) JoinRoom(roomIdx int, client port.Client) {
	c.pool.RegisterJob(func() {
		c.rooms[roomIdx] = append(c.rooms[roomIdx], client)
		c.logger.Println(len(c.rooms[roomIdx]))
	})
}

func (c *Chat) ExitRoom(roomIdx, userIdx int) {
	c.pool.RegisterJob(func() {
		clients := c.rooms[roomIdx]
		for i, client := range clients {
			if client.GetUserIdx() == userIdx {
				c.rooms[roomIdx] = append(clients[:i], clients[i+1:]...)
				if len(c.rooms[roomIdx]) == 0 {
					delete(c.rooms, roomIdx)
				}
				c.logger.Printf("current clients : %d, current rooms : %d", len(c.rooms[roomIdx]), len(c.rooms))
				return
			}
		}
		c.logger.Println("no match user idx in the chat room")
	})
}

func (c *Chat) ExitAllRooms(roomIdxs *[]int, client port.Client) {
	c.pool.RegisterJob(func() {
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
	})
}

//* (2) message
func (c *Chat) BroadcastMsg(msg port.Message) {
	c.pool.RegisterJob(func() {
		roomIdx := msg.GetRoomIdx()

		for _, client := range c.rooms[roomIdx] {
			err := client.SendMsg(msg)
			if err != nil {
				c.logger.Printf("sending message error: %s", err)
				continue
			}
		}
	})
}
