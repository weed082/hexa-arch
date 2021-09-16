package main

import (
	"os"

	"github.com/ByungHakNoh/hexagonal-microservice/internal/application"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/repository"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/repository/mongo_db"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/repository/mysql"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/rest"
)

func main() {
	// env
	dbaseDriver := os.Getenv("DB_DRIVER")
	dsourceName := os.Getenv("DS_NAME")

	// database
	mongoDB := mongo_db.NewMongoDB()
	mysqlDB := mysql.NewMysql(dbaseDriver, dsourceName)
	defer mongoDB.Disconnect()
	defer mysqlDB.Disconnect()

	// repository
	userRepo := repository.NewUserRepo(mysqlDB, mongoDB)

	// application
	userApp := application.NewUserApp(userRepo)

	// server
	server := rest.NewRestAdapter(userApp)
	server.Run()
}

func runRestServer() {

}
