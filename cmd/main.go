package main

import (
	"fmt"
	"os"

	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/repository"
)

func main() {
	dbaseDriver := os.Getenv("DB_DRIVER")
	dsourceName := os.Getenv("DS_NAME")
	repository.NewMongoDB()
	// repository.NewMysql(dbaseDriver, dsourceName)
	fmt.Println(dbaseDriver, dsourceName)
}
