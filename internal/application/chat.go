package application

import (
	"log"

	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/port"
)

type Chat struct {
	logger *log.Logger
	repo   port.ChatRepo
	pool   port.WorkerPool
	rooms  map[int][]port.ChatClient
}

func NewChat(logger *log.Logger, repo port.ChatRepo, pool port.WorkerPool) *Chat {
	pool.Generate(1) // only need single worker for sync
	return &Chat{
		logger: logger,
		repo:   repo,
		pool:   pool,
		rooms:  make(map[int][]port.ChatClient),
	}
}

//! ----------- 1. chat room -----------
// create a chat room
func (c *Chat) CreateRoom() (int, error) {
	roomIdx := 1 // TODO: need to get room idx from db
	return roomIdx, nil
}

// join room
func (c *Chat) JoinRoom(roomIdx int, client port.ChatClient) error {
	c.ConnectRoom(roomIdx, client) // connect room
	return nil
}

// exist a room
func (c *Chat) ExitRoom(roomIdx, userIdx int) {
}

/** ----------- 3. message ----------- */
//* send message to all clients that are participated in the chat room given
func (c *Chat) SendMsg(msg port.ChatMsg) {
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

/** ----------- ?. pool job for chat room ----------- */
// connect the client to the chat room
func (c *Chat) ConnectRoom(roomIdx int, client port.ChatClient) {
	c.pool.RegisterJob(func() {
		c.rooms[roomIdx] = append(c.rooms[roomIdx], client)
	})
}

// connect rooms
func (c *Chat) ConnectRooms(rooms []int, client port.ChatClient) {
	c.pool.RegisterJob(func() {
		for _, roomIdx := range rooms {
			c.rooms[roomIdx] = append(c.rooms[roomIdx], client)
		}
		c.logger.Println(len(c.rooms))
	})
}

func (c *Chat) DisconnectRoom(roomIdx, userIdx int) {
	c.pool.RegisterJob(func() {
		clients := c.rooms[roomIdx]
		for i, client := range clients {
			if client.GetUserIdx() == userIdx {
				c.rooms[roomIdx] = append(clients[:i], clients[i+1:]...)
				if len(c.rooms[roomIdx]) == 0 {
					delete(c.rooms, roomIdx)
				}
				c.logger.Printf("current clients : %d, current rooms : %d", len(c.rooms[roomIdx]), len(c.rooms))
				break
			}
		}
		c.logger.Println("no match user idx in the chat room")
	})

}

// disconnect the client from all chat rooms that are participated in
func (c *Chat) DisconnectRooms(client port.ChatClient) {
	c.pool.RegisterJob(func() {
		for roomIdx, clients := range c.rooms {
			for i, existClient := range clients {
				if client == existClient {
					c.rooms[roomIdx] = append(clients[:i], clients[i+1:]...) // delete user
					if len(c.rooms[roomIdx]) == 0 {
						delete(c.rooms, roomIdx) // delete room
					}
					break
				}
			}
		}
	})
}
