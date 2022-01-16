package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ByungHakNoh/hexagonal-microservice/external/router"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/port"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/model"
)

type User struct {
	logger *log.Logger
	app    port.User
}

func NewUser(logger *log.Logger, app port.User) *User {
	return &User{
		logger: logger,
		app:    app,
	}
}

func (handler *User) Register(r port.Router) {
	r.Get("/test/:param", handler.test)
}

func (handler *User) test(w http.ResponseWriter, r *http.Request) {
	urlParams := router.UrlParam(r)
	user := model.User{Idx: 4, Name: urlParams["param"]}
	jsonBytes, err := json.Marshal(user)
	if err != nil {
		handler.logger.Printf("json marshal failed : %s", err)
	}
	fmt.Fprintln(w, string(jsonBytes))
}
