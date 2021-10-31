package main

import (
	"github.com/ByungHakNoh/hexagonal-microservice/internal/application"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/repository"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/repository/mongo_db"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/repository/mysql"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/server/rest"
)

const (
	dbaseDriver = "mysql"
	dsourceName = "root:Admin123@/test"
)

func main() {
	// DB
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

	// rest
	rest := rest.NewRestAdapter(userApp, templateApp)
	rest.Run("8080")
}
