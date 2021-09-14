package main

import (
	"os"

	"github.com/ByungHakNoh/hexagonal-microservice/internal/application"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/repository"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/server/rest"
)

func main() {
	// env
	dbaseDriver := os.Getenv("DB_DRIVER")
	dsourceName := os.Getenv("DS_NAME")

	// repository
	mongoDB := repository.NewMongoDB()
	mysqlDB := repository.NewMysql(dbaseDriver, dsourceName)

	// application
	userApp := application.NewUserApp(mongoDB, mysqlDB)

	// server
	server := rest.NewRestAdapter(userApp)
	server.Run()
}
