package application

import "github.com/ByungHakNoh/hexagonal-microservice/internal/framework/port"

type ChatApp struct {
	repo port.ChatRepo
}

func NewChatApp(repo port.ChatRepo) *ChatApp {
	return &ChatApp{
		repo: repo,
	}
}

func (c *ChatApp) CreateRoom(rooms map[int][]interface{}, client interface{}) error {
	roomIdx, err := c.repo.CreateRoom()
	if err != nil {
		return err
	}
	rooms[roomIdx] = []interface{}{client}

	return nil
}

func (c *ChatApp) JoinRoom(rooms map[int][]interface{}, client interface{}) {

}

func (c *ChatApp) RemoveRoom(rooms map[int][]interface{}, client interface{}) {

}

func (c *ChatApp) ExitRoom(rooms map[int][]interface{}, client interface{}) {

}

func (c *ChatApp) SendMessage(rooms map[int][]interface{}, client interface{}) {

}
