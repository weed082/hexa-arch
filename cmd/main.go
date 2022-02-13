package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"

	"github.com/ByungHakNoh/hexagonal-microservice/internal/application"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/core/pool"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/repository"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/repository/mongo"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/repository/mysql"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/server/grpc"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/server/rest"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/port"
)

var (
	// logger
	logger = log.New(os.Stdout, "LOG", log.LstdFlags|log.Llongfile)
	// server
	grpcServer *grpc.Grpc
	restServer *rest.Rest
	// database
	mysqlDB = mysql.NewMysql(logger, "mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_DATABASE")))
	mongoDB = mongo.NewMongoDB(logger, fmt.Sprintf("mongodb://%s:%s@%s", os.Getenv("MONGO_USER"), os.Getenv("MONGO_PASSWORD"), os.Getenv("MONGO_HOST")))
	// chat worker pool
	chatPool = pool.NewWorkerPool(logger, chatWg, 100)
	chatWg   = &sync.WaitGroup{}
)

func main() {
	logger.Printf("cpu : %d", runtime.GOMAXPROCS(runtime.NumCPU()))
	chatRepo := repository.NewChat(logger, mysqlDB)
	chatApp := application.NewChat(logger, chatRepo, chatPool)

	go runRest(chatApp)
	go runGrpc(chatApp)
	gracefulShutdown() // block until grpc and rest server finishes
}

//! run rest server
func runRest(chatApp port.Chat) {
	// user
	userRepo := repository.NewUser(logger, mysqlDB, mongoDB)
	userApp := application.NewUser(logger, userRepo)
	// rest
	restServer = rest.NewRestAdapter(logger, userApp, chatApp)
	restServer.Run(os.Getenv("REST_PORT"))
}

//! run grpc server
func runGrpc(chatApp port.Chat) {
	// grpc
	grpcServer = grpc.NewServer(logger, chatApp)
	grpcServer.Run(os.Getenv("GRPC_PORT"))
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
	restServer.Stop()

	//* close grpc server
	logger.Println("gracefully shutdown Grpc") //! need to log before call graceful shutdown or race condition problem occur
	grpcServer.Stop()

	//* stop chat pool
	chatPool.Stop()
	chatWg.Wait()
}
