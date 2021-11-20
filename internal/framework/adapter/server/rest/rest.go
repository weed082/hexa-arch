package rest

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/server/rest/core/router"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/server/rest/core/server"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/server/rest/handler"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/port"
)

type Rest struct {
	server  *http.Server
	userApp port.User
}

func NewRestAdapter(userApp port.User) *Rest {
	return &Rest{userApp: userApp}
}

func (r *Rest) Run(port string) {
	router := router.New()
	handler.NewUserHandler(r.userApp).Register(router)

	r.server = server.New(router, ":"+port)
	err := r.server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("rest server error: %s", err)
	}
	log.Println(err)
}

func (r *Rest) Stop() {
	ctx, restCancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
	defer restCancelFunc()
	if err := r.server.Shutdown(ctx); err != nil {
		log.Printf("shutting down rest server failed: %s", err)
	}
}
