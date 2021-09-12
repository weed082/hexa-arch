package repository

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoDB struct {
	ctx           context.Context
	ctxCancelFunc context.CancelFunc
	client        *mongo.Client
}

func NewMongoDB() *MongoDB {
	ctx, ctxCancelFunc := context.WithTimeout(context.Background(), 5*time.Second) // repository context
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:example@localhost:27017"))
	if err != nil {
		log.Fatalf("mongoDB new client failed : %v", err)
	}

	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("mongoDB client connect failed : %v", err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatalf("mongoDB ping failed : %v", err)
	}
	return &MongoDB{
		ctx:           ctx,
		ctxCancelFunc: ctxCancelFunc,
		client:        client,
	}
}

// disconnect to mongoDB
func (mongoDB MongoDB) Disconnect() {
	err := mongoDB.client.Disconnect(mongoDB.ctx)
	if err != nil {
		log.Fatalf("mongoDB disconnect failed : %v", err)
	}
	mongoDB.ctxCancelFunc()
}
