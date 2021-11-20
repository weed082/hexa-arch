package chat

import "github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/server/grpc/chat/pb"

type Client struct {
	userIdx int
	stream  pb.ChatService_ChatServiceServer
}

func (c *Client) GetUserIdx() int {
	return c.userIdx
}

func (c *Client) SendMsg(msg *pb.MsgRes) error {
	return c.stream.Send(msg)
}
