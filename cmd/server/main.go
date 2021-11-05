package main

import (
	"github.com/ByungHakNoh/hexagonal-microservice/internal/application"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/server/grpc"
)

func main() {
	var testRepo struct{}
	fileApp := application.NewFileApp(testRepo)
	server := grpc.NewServer(fileApp)
	server.Run("9000")
}
