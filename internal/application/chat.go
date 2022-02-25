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

// put client to chat rooms
func (c *Chat) Connect(client port.ChatClient) {
	// TODO: get all rooms that client were participated in before
	rooms := []int{1, 2}
	c.connectRooms(rooms, client)
}

// remove client from all chat rooms
func (c *Chat) Disconnect(client port.ChatClient) {
	c.disconnectRooms(client)
}

// create a chat room
func (c *Chat) CreateRoom() (int, error) {
	roomIdx := 1 // TODO: db work
	return roomIdx, nil
}

// join room
func (c *Chat) JoinRoom(roomIdx int, client port.ChatClient) error {
	c.connectRoom(roomIdx, client) // connect room
	return nil
}

// exist a room
func (c *Chat) ExitRoom(roomIdx int, client port.ChatClient) error {
	//TODO: db work
	c.disconnectRoom(roomIdx, client.GetUserIdx()) // remove client from target room
	return nil
}

// send message to all clients that are participated in the chat room given
func (c *Chat) SendMsg(roomIdx int, msg interface{}) {
	c.pool.RegisterJob(func() {
		for _, client := range c.rooms[roomIdx] {
			err := client.SendMsg(msg)
			if err != nil {
				c.logger.Printf("sending message error: %s", err)
				continue
			}
		}
	})
}

// send res back to client
func (c *Chat) SendRes(client port.ChatClient, msg interface{}) {
	c.pool.RegisterJob(func() {
		client.SendMsg(msg)
	})
}

/** ----------- ?. chat pool job for chat room ----------- */
// connect the client to the chat room
func (c *Chat) connectRoom(roomIdx int, client port.ChatClient) {
	c.pool.RegisterJob(func() {
		c.rooms[roomIdx] = append(c.rooms[roomIdx], client)
	})
}

// connect rooms
func (c *Chat) connectRooms(rooms []int, client port.ChatClient) {
	c.pool.RegisterJob(func() {
		for _, roomIdx := range rooms {
			c.rooms[roomIdx] = append(c.rooms[roomIdx], client)
		}
		c.logger.Println(len(c.rooms))
	})
}

// remove client from target room
func (c *Chat) disconnectRoom(roomIdx, userIdx int) {
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
func (c *Chat) disconnectRooms(client port.ChatClient) {
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
