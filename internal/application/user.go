package application

import (
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/port"
)

type UserApp struct {
	mysqlDB port.UserSql
	mongoDB port.UserNoSql
}

func NewUserApp(mysqlDB port.UserSql, mongoDB port.UserNoSql) *UserApp {
	return &UserApp{mysqlDB: mysqlDB, mongoDB: mongoDB}
}
