package rest

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/server/rest/handler"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/server/rest/router"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/server/rest/server"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/port"
)

type Rest struct {
	logger   *log.Logger
	server   *http.Server
	userApp  port.User
	chatApp  port.Chat
	chatPool port.Consumer
}

func NewRestAdapter(
	logger *log.Logger,
	userApp port.User,
	chatApp port.Chat,
	chatPool port.Consumer,
) *Rest {
	return &Rest{
		logger:   logger,
		userApp:  userApp,
		chatApp:  chatApp,
		chatPool: chatPool,
	}
}

func (r *Rest) Run(port string) {
	router := router.New() // custom router
	// handlers
	handler.NewUser(r.logger, r.userApp).Register(router)
	handler.NewChat(r.logger, r.chatApp).Register(router)

	// serve rest server
	server := server.NewServer(router, ":"+port)
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		r.logger.Fatalf("rest server error: %s", err)
	}
}

func (r *Rest) Stop() {
	ctx, restCancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
	defer restCancelFunc()
	if err := r.server.Shutdown(ctx); err != nil {
		r.logger.Printf("shutting down rest server failed: %s", err)
	}
}
