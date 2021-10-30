package main

import (
	"github.com/ByungHakNoh/hexagonal-microservice/internal/application"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/repository"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/repository/mongo_db"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/repository/mysql"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/server/rest"
)

func main() {
	// env
	dbaseDriver := "mysql"
	dsourceName := "root:Admin123@/test"

	// database
	mongoDB := mongo_db.NewMongoDB()
	mysqlDB := mysql.NewMysql(dbaseDriver, dsourceName)
	defer mongoDB.Disconnect()
	defer mysqlDB.Disconnect()

	// repository
	userRepo := repository.NewUserRepo(mysqlDB, mongoDB)
	templateRepo := repository.NewTemplateRepo(mysqlDB)

	// application
	userApp := application.NewUserApp(userRepo)
	templateApp := application.NewTemplateApp(templateRepo)

	// rest server
	server := rest.NewRestAdapter(userApp, templateApp)
	server.Run(":8080")
}
