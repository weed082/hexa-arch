package main

import (
	"os"

	"github.com/ByungHakNoh/hexagonal-microservice/internal/application"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/mongo_db"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/mysql"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/rest"
)

func main() {
	// env
	dbaseDriver := os.Getenv("DB_DRIVER")
	dsourceName := os.Getenv("DS_NAME")

	// repository
	mongoDB := mongo_db.NewMongoDB()
	mysqlDB := mysql.NewMysql(dbaseDriver, dsourceName)

	// application
	userApp := application.New(mongoDB, mysqlDB)

	// server
	server := rest.NewRestAdapter(userApp)
	server.Run()
}
