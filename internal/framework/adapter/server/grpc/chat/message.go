package chat

import "github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/server/grpc/chat/pb"

type Message struct {
	msg *pb.MsgRes
}

func (m *Message) GetRoomIdx() int {
	return int(m.msg.RoomIdx)
}
