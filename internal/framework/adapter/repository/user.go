package repository

import "github.com/ByungHakNoh/hexagonal-microservice/internal/framework/port"

type User struct {
	sqlDB   port.UserSql
	noSqlDB port.UserNoSql
}

func NewUser(sqlDB port.UserSql, noSqlDB port.UserNoSql) *User {
	return &User{
		sqlDB:   sqlDB,
		noSqlDB: noSqlDB,
	}
}
