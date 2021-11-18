package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
	"time"

	"github.com/ByungHakNoh/hexagonal-microservice/internal/application"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/core/concurrency"
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
	// worker pool
	wg              = &sync.WaitGroup{}
	ctx, cancelFunc = context.WithCancel(context.Background())

	wp = concurrency.NewWorkerPool(wg, ctx)
)

func main() {
	log.Printf("cpu : %d", runtime.GOMAXPROCS(runtime.NumCPU()))
	terminationChan := make(chan os.Signal, 1)
	signal.Notify(terminationChan, os.Interrupt, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)

	go runRest()
	go runGrpc()
	go wp.Start()

	<-terminationChan
	cancelFunc()
	gracefulShutdown()
	wg.Wait()
}

//! run rest server
func runRest() {
	// repository
	userRepo := repository.NewUserRepo(mysqlDB, mongoDB)
	// application
	userApp := application.NewUserApp(userRepo)
	// rest
	Rest = rest.NewRestAdapter(userApp)
	Rest.Run("8080")
}

//! run grpc server
func runGrpc() {
	// repository
	chatRepo := repository.NewChatRepo(mysqlDB)
	// application
	chatApp := application.NewChatApp(chatRepo)
	// grpc
	Grpc = grpc.NewServer(wp, chatApp)
	Grpc.Run("9000")
	log.Println("grpc shut down")
}

//! shutdown rest, grpc gracefully + close db
func gracefulShutdown() {
	ctx, restCancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
	defer mysqlDB.Disconnect()
	defer mongoDB.Disconnect()
	defer restCancelFunc()
	if err := Rest.Server.Shutdown(ctx); err != nil {
		log.Printf("shutting down rest server failed: %s", err)
	}
	Grpc.Server.GracefulStop()
}
