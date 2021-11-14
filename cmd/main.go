package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/ByungHakNoh/hexagonal-microservice/internal/application"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/repository"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/repository/mongo_db"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/repository/mysql"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/server/grpc"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/server/rest"
)

const (
	DB_DRIVER      = "mysql"
	DB_SOURCE_NAME = "root:Admin123@/go_arch"
)

var (
	// server
	Grpc *grpc.Grpc
	Rest *rest.Rest
	// database
	mysqlDB = mysql.NewMysql(DB_DRIVER, DB_SOURCE_NAME)
	mongoDB = mongo_db.NewMongoDB()
	// sync
	wg = &sync.WaitGroup{}
)

func main() {
	terminationChan := make(chan os.Signal, 1)
	signal.Notify(terminationChan, os.Interrupt, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)

	wg.Add(2)
	go runRest()
	go runGrpc()

	<-terminationChan
	gracefulShutdown()
	wg.Wait()
}

//! run rest server
func runRest() {
	defer wg.Done()
	// repository
	userRepo := repository.NewUserRepo(mysqlDB, mongoDB)
	templateRepo := repository.NewTemplateRepo(mysqlDB)
	// application
	userApp := application.NewUserApp(userRepo)
	templateApp := application.NewTemplateApp(templateRepo)
	// rest
	Rest = rest.NewRestAdapter(userApp, templateApp)
	Rest.Run("8080")
}

//! run grpc server
func runGrpc() {
	defer wg.Done()
	// repository
	chatRepo := repository.NewChatRepo(mysqlDB)
	// application
	chatApp := application.NewChatApp(chatRepo)
	// grpc
	Grpc = grpc.NewServer(chatApp)
	Grpc.Run("9000")
	log.Println("grpc shut down")
}

//! shutdown rest, grpc gracefully + close db
func gracefulShutdown() {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
	defer mysqlDB.Disconnect()
	defer mongoDB.Disconnect()
	defer cancelFunc()
	if err := Rest.Server.Shutdown(ctx); err != nil {
		log.Printf("shutting down rest server failed: %s", err)
	}
	Grpc.Server.GracefulStop()
}
