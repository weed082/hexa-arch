package repository

import "github.com/ByungHakNoh/hexagonal-microservice/internal/framework/port"

type Chat struct {
	sql port.ChatSql
}

func NewChat(sql port.ChatSql) *Chat {
	return &Chat{
		sql: sql,
	}
}

func (r *Chat) CreateRoom() (int, error) {
	return 1, nil
}

func (r *Chat) SendMessage() error {
	return nil
}
