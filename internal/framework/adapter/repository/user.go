package repository

import (
	"log"

	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/port"
)

type User struct {
	logger  *log.Logger
	sqlDB   port.UserSql
	noSqlDB port.UserNoSql
}

func NewUser(logger *log.Logger, sqlDB port.UserSql, noSqlDB port.UserNoSql) *User {
	return &User{
		logger:  logger,
		sqlDB:   sqlDB,
		noSqlDB: noSqlDB,
	}
}
