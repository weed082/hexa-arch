package repository

import (
	"log"
)

type userSql interface {
}

type userNosql interface {
}

type User struct {
	logger  *log.Logger
	sqlDB   userSql
	noSqlDB userNosql
}

func NewUser(logger *log.Logger, sqlDB userSql, noSqlDB userNosql) *User {
	return &User{
		logger:  logger,
		sqlDB:   sqlDB,
		noSqlDB: noSqlDB,
	}
}
