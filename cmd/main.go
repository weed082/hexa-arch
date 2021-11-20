package main

import (
	"log"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
	"time"

	"github.com/ByungHakNoh/hexagonal-microservice/external/pool"
	"github.com/ByungHakNoh/hexagonal-microservice/external/router"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/application"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/repository"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/repository/mongo_db"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/repository/mysql"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/server/grpc"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/server/rest"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/port"
)

const (
	DB_DRIVER      = "mysql"
	DB_SOURCE_NAME = "root:Admin123@/go_arch"
)

var (
	// logger
	logger = log.New(os.Stdout, "LOG", log.LstdFlags|log.Llongfile)

	// server
	Grpc *grpc.Grpc
	Rest *rest.Rest

	// database
	mysqlDB = mysql.NewMysql(logger, DB_DRIVER, DB_SOURCE_NAME)
	mongoDB = mongo_db.NewMongoDB(logger, 5*time.Second)

	// chat wait group
	chatWg = &sync.WaitGroup{}
)

func main() {
	logger.Printf("cpu : %d", runtime.GOMAXPROCS(runtime.NumCPU()))
	go runRest()
	go runGrpc()
	gracefulShutdown() // block until grpc and rest server finishes
}

//! run rest server
func runRest() {
	// repository
	userRepo := repository.NewUser(logger, mysqlDB, mongoDB)
	// application
	userApp := application.NewUser(logger, userRepo)
	// router
	router := router.New()
	// rest
	Rest = rest.NewRestAdapter(logger, router, userApp)
	Rest.Run("8080")
}

//! run grpc server
func runGrpc() {
	chatPool := pool.NewWorkerPool(logger, chatWg, 100)
	defer chatPool.Stop()
	chatPool.Generate(1)
	// repository
	chatRepo := repository.NewChat(logger, mysqlDB)
	// application
	chatApp := application.NewChat(logger, map[int]port.Client{}, chatRepo)
	// grpc
	Grpc = grpc.NewServer(logger, chatPool, chatApp)
	Grpc.Run("9000")
}

//! grcefully shutdown in order
func gracefulShutdown() {
	//* listen exit signal from os
	terminationChan := make(chan os.Signal, 1)
	signal.Notify(terminationChan, os.Interrupt, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-terminationChan

	//* close db
	defer mysqlDB.Disconnect()
	defer mongoDB.Disconnect()

	//* close rest server
	logger.Println("gracefully shutdown Rest") //! need to log before call graceful shutdown or race condition problem occur
	Rest.Stop()

	//* close grpc server
	logger.Println("gracefully shutdown Grpc") //! need to log before call graceful shutdown or race condition problem occur
	Grpc.Stop()

	chatWg.Wait() //* wait for chat pool to finish
}
