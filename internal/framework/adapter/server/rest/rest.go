package rest

import (
	"log"
	"net/http"

	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/server/rest/core/router"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/server/rest/core/server"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/server/rest/handler"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/port"
)

type Rest struct {
	Server      *http.Server
	userApp     port.UserApp
	templateApp port.TemplateApp
}

func NewRestAdapter(userApp port.UserApp, templateApp port.TemplateApp) *Rest {
	return &Rest{userApp: userApp, templateApp: templateApp}
}

func (r *Rest) Run(port string) {
	router := router.New()
	handler.NewUserHandler(r.userApp).Register(router)
	handler.NewTemplateHandler(r.templateApp).Register(router)

	r.Server = server.New(router, ":"+port)
	log.Println(r.Server.ListenAndServe())
}
