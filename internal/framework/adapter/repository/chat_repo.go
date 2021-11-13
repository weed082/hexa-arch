package repository

import "github.com/ByungHakNoh/hexagonal-microservice/internal/framework/port"

var count = 0

type ChatRepo struct {
	sql port.ChatSql
}

func NewChatRepo(sql port.ChatSql) *ChatRepo {
	return &ChatRepo{
		sql: sql,
	}
}

func (r *ChatRepo) CreateRoom() (int, error) {
	count++
	return count, nil
}

func (r *ChatRepo) SendMessage() error {
	return nil
}
