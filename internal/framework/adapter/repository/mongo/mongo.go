package mongo

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoDB struct {
	logger        *log.Logger
	ctx           context.Context
	ctxCancelFunc context.CancelFunc
	client        *mongo.Client
}

func NewMongoDB(logger *log.Logger, applyUri string, ctxTimeout time.Duration) *MongoDB {
	ctx, ctxCancelFunc := context.WithTimeout(context.Background(), ctxTimeout) // repository context
	client, err := mongo.NewClient(options.Client().ApplyURI(applyUri))
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
		logger:        logger,
		ctx:           ctx,
		ctxCancelFunc: ctxCancelFunc,
		client:        client,
	}
}

// disconnect to mongoDB
func (mongo *MongoDB) Disconnect() {
	err := mongo.client.Disconnect(mongo.ctx)
	if err != nil {
		mongo.logger.Fatalf("mongoDB disconnect failed : %v", err)
	}
	mongo.ctxCancelFunc()
}
