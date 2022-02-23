package chat

import "github.com/gorilla/websocket"

type Client struct {
	userIdx int
  name string
	conn    *websocket.Conn
}

func (c *Client) GetUserIdx() int {
	return c.userIdx
}

func (c *Client) SendMsg(msg interface{}) error {
	err := c.conn.WriteJSON(msg)
	if err != nil {
		return err
	}
	return nil
}
