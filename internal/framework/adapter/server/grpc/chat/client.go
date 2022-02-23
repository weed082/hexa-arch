package chat

import (
	"errors"

	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/server/grpc/chat/pb"
)

type Client struct {
	userIdx int
  name string
	stream  pb.ChatService_ChatServiceServer
}

func (c *Client) GetUserIdx() int {
	return c.userIdx
}

func (c *Client) SendMsg(msg interface{}) error {
	msgRes, ok := msg.(*pb.MsgRes)
	if !ok {
		return errors.New("type assertion to pb struct failed")
	}
	return c.stream.Send(msgRes)
}
