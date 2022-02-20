package repository

import (
	"log"
)

type chatSql interface {
}

type Chat struct {
	logger *log.Logger
	sql    chatSql
}

func NewChat(logger *log.Logger, sql chatSql) *Chat {
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
