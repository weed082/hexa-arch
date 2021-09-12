package mongoDB

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoDB struct {
	Client        *mongo.Client
	Ctx           context.Context
	ctxCancelFunc context.CancelFunc
}

func New() *MongoDB {
	ctx, ctxCancelFunc := context.WithTimeout(context.Background(), 5*time.Second) // repository context
	client := connect(ctx)
	return &MongoDB{
		Client:        client,
		Ctx:           ctx,
		ctxCancelFunc: ctxCancelFunc,
	}
}

// connect mongoDB
func connect(ctx context.Context) *mongo.Client {
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
	return client
}

// disconnect to mongoDB
func (db MongoDB) Disconnect() {
	err := db.Client.Disconnect(db.Ctx)
	if err != nil {
		log.Fatalf("mongoDB disconnect failed : %v", err)
	}
	db.ctxCancelFunc()
}
