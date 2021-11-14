package repository

import "github.com/ByungHakNoh/hexagonal-microservice/internal/framework/port"

type ChatRepo struct {
	sql port.ChatSql
}

func NewChatRepo(sql port.ChatSql) *ChatRepo {
	return &ChatRepo{
		sql: sql,
	}
}

func (r *ChatRepo) CreateRoom() (int, error) {
	return 1, nil
}

func (r *ChatRepo) SendMessage() error {
	return nil
}
