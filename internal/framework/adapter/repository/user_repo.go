package repository

import "github.com/ByungHakNoh/hexagonal-microservice/internal/framework/port"

type UserRepo struct {
	sqlDB   port.UserSql
	noSqlDB port.UserNoSql
}

func NewUserRepo(sqlDB port.UserSql, noSqlDB port.UserNoSql) *UserRepo {
	return &UserRepo{
		sqlDB:   sqlDB,
		noSqlDB: noSqlDB,
	}
}
