package main

import (
	"github.com/ByungHakNoh/hexagonal-microservice/internal/application"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/repository"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/repository/mysql"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/server/grpc"
)

const (
	dbaseDriver = "mysql"
	dsourceName = "root:Admin123@/go_arch"
)

func main() {
	chatSql := mysql.NewMysql(dbaseDriver, dsourceName)
	chatRepo := repository.NewChatRepo(chatSql)
	chatApp := application.NewChatApp(chatRepo)
	server := grpc.NewServer(chatApp)
	server.Run("9000")
}
