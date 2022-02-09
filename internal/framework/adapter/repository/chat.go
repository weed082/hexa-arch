package repository

import (
	"log"

	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/port"
)

type Chat struct {
	logger *log.Logger
	sql    port.ChatSql
}

func NewChat(logger *log.Logger, sql port.ChatSql) *Chat {
	return &Chat{
		logger: logger,
		sql:    sql,
	}
}

func (r *Chat) CreateRoom() (int, error) {
	return 1, nil
}

func (r *Chat) UploadMsg() error {
	return nil
}
