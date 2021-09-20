package rest

import (
	"log"

	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/server/rest/core/router"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/server/rest/core/server"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/server/rest/handler"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/port"
)

type Rest struct {
	userApp     port.UserApp
	templateApp port.TemplateApp
}

func NewRestAdapter(userApp port.UserApp, templateApp port.TemplateApp) *Rest {
	return &Rest{userApp, templateApp}
}

func (rest Rest) Run() {
	r := router.New()
	handler.NewUserHandler(rest.userApp).Register(r)
	handler.NewTemplateHandler(rest.templateApp).Register(r)
	httpServer := server.New(r, ":8080")
	log.Fatal(httpServer.ListenAndServe())
}
