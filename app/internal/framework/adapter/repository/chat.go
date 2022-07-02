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

/** -------------------- room -------------------- */
func (r *Chat) CreateRoom() (int, error) {
	return 1, nil
}

func (r *Chat) Join() error {
	return nil
}

func (r *Chat) Exit() error {
	return nil
}

/** -------------------- msg -------------------- */
func (r *Chat) UploadMsg() error {
	return nil
}

func (r *Chat) UpdateMsg() error {
	return nil
}
