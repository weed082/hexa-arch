package application

import (
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/port"
)

type Application struct {
	mysqlDB port.UserSql
	mongoDB port.UserNoSql
}

func New(mysqlDB port.UserSql, mongoDB port.UserNoSql) *Application {
	return &Application{mysqlDB: mysqlDB, mongoDB: mongoDB}
}
