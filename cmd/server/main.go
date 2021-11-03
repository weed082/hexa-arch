package main

import (
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/server/grpc"
)

func main() {
	server := grpc.NewServer()
	server.Run("9000")
}
