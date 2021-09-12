package repository

import (
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/repository/mongoDB"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/repository/mysqlDB"
	"go.mongodb.org/mongo-driver/bson"
)

type Repository struct {
	mongoDB mongoDB.MongoDB
	mysqlDB mysqlDB.MysqlDB
}

func New() *Repository {
	return &Repository{
		mongoDB: *mongoDB.New(),
		mysqlDB: *mysqlDB.New(),
	}
}

func (repository Repository) Disconnect() {
	repository.mongoDB.Disconnect()
	repository.mysqlDB.Disconnect()
}

func (repository Repository) Test() ([]string, error) {
	return repository.mongoDB.Client.ListDatabaseNames(repository.mongoDB.Ctx, bson.D{})
}
