package chat

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

type Client struct {
	conn net.Conn
}

func NewClient(conn net.Conn) *Client {
	return &Client{
		conn: conn,
	}
}

func (client *Client) Connect() {
	for {
		msg, err := bufio.NewReader(client.conn).ReadString('\n')
		if err != nil {
			log.Printf("bufio failed to initialized: %s", err.Error())
			return
		}
		msg = strings.Trim(msg, "\r\n")
		args := strings.Split(msg, " ")
		client.handleMessage(args)
	}
}

func (client *Client) handleMessage(args []string) {
	cmd := strings.TrimSpace(args[0])
	switch cmd {
	case "/nick":
	case "/join":
	case "/rooms":
	case "/msg":
	case "/quit":
	default:
		client.err(fmt.Errorf("unknown command: %s", cmd))
	}
}

func (client *Client) err(err error) {
}
