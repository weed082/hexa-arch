package rest

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/server/rest/handler"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/port"
)

type Rest struct {
	logger  *log.Logger
	server  *http.Server
	router  port.Router
	userApp port.User
}

func NewRestAdapter(logger *log.Logger, router port.Router, userApp port.User) *Rest {
	return &Rest{
		logger:  logger,
		router:  router,
		userApp: userApp,
	}
}

func (r *Rest) Run(port string) {
	handler.NewUser(r.logger, r.userApp).Register(r.router)
	r.server = r.NewServer(r.router, ":"+port)
	err := r.server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		r.logger.Fatalf("rest server error: %s", err)
	}
	r.logger.Println(err)
}

func (r *Rest) Stop() {
	ctx, restCancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
	defer restCancelFunc()
	if err := r.server.Shutdown(ctx); err != nil {
		r.logger.Printf("shutting down rest server failed: %s", err)
	}
}
