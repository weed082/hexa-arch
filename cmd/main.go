package main

import (
	"fmt"
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
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/repository/mongo"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/repository/mysql"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/server/grpc"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/server/rest"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/port"
)

var (
	// mysql env
	mysqlUser     = os.Getenv("MYSQL_USER")
	mysqlPassword = os.Getenv("MYSQL_PASSWORD")
	mysqlHost     = os.Getenv("MYSQL_HOST")
	mysqlDatabase = os.Getenv("MYSQL_DATABASE")
	// mongo env
	mongoUser     = os.Getenv("MONGO_USER")
	mongoPassword = os.Getenv("MONGO_PASSWORD")
	mongoHost     = os.Getenv("MONGO_HOST")

	// port env
	restPort = os.Getenv("REST_PORT")
	grpcPort = os.Getenv("GRPC_PORT")
)

var (
	// logger
	logger = log.New(os.Stdout, "LOG", log.LstdFlags|log.Llongfile)
	// server
	Grpc *grpc.Grpc
	Rest *rest.Rest
	// database
	mysqlDB *mysql.Mysql
	mongoDB *mongo.MongoDB
	// chat wait group
	chatWg = &sync.WaitGroup{}
)

func init() {
	// init mysql
	dbSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s", mysqlUser, mysqlPassword, mysqlHost, mysqlDatabase)
	mysqlDB = mysql.NewMysql(logger, "mysql", dbSourceName)

	// init mongo
	applyUri := fmt.Sprintf("mongodb://%s:%s@%s", mongoUser, mongoPassword, mongoHost)
	mongoDB = mongo.NewMongoDB(logger, applyUri, 5*time.Second)
}

func main() {
	logger.Printf("csdfpu : %d", runtime.GOMAXPROCS(runtime.NumCPU()))
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
	Rest.Run(restPort)
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
	Grpc.Run(grpcPort)
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
