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
	dbaseDriver = "mysql"
	dsourceName = "root:Admin123@/go_arch"
)

// server
var (
	Grpc *grpc.Grpc
	Rest *rest.Rest
)

var wg = &sync.WaitGroup{}

func main() {
	mysqlDB := mysql.NewMysql(dbaseDriver, dsourceName)
	mongoDB := mongo_db.NewMongoDB()
	defer mysqlDB.Disconnect()
	defer mongoDB.Disconnect()

	wg.Add(2)
	go runRest(mysqlDB, mongoDB)
	go runGrpc(mysqlDB)

	terminationChan := make(chan os.Signal, 1)
	signal.Notify(terminationChan, os.Interrupt, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-terminationChan
	gracefulShutdown()
	wg.Wait()
}

func runRest(mysqlDB *mysql.Mysql, mongoDB *mongo_db.MongoDB) {
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

func runGrpc(mysqlDB *mysql.Mysql) {
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

func gracefulShutdown() {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelFunc()
	if err := Rest.Server.Shutdown(ctx); err != nil {
		log.Printf("shutting down rest server failed: %s", err)
	}
	Grpc.Server.GracefulStop()
}
