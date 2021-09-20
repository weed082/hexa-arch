package handler

import (
	"log"
	"net/http"

	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/server/rest/core/router"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/port"
)

type TemplateHandler struct {
	app port.TemplateApp
}

func NewTemplateHandler(app port.TemplateApp) *TemplateHandler {
	return &TemplateHandler{app}
}

func (handler TemplateHandler) Register(r *router.Router) {
	r.Get("/", handler.test)
}

func (handler TemplateHandler) test(w http.ResponseWriter, r *http.Request) {
	tmpl, pageData, err := handler.app.RenderPage()
	if err != nil {
		log.Fatalf("render page failure: %v", err)
	}

	err = tmpl.Execute(w, pageData)
	if err != nil {
		log.Fatalf("excute page template failure: %v", err)
	}
}
