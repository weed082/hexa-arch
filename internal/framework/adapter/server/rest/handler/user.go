package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/server/rest/router"
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

func (h *User) Register(r port.Router) {
	r.Get("/test/:param(param|sub)", h.test)
	r.Post("/user/create", h.create)
}

func (h *User) test(w http.ResponseWriter, r *http.Request) {
	urlParams := router.UrlParam(r)
	user := model.User{Idx: 3, Name: urlParams["param"]}
	jsonBytes, err := json.Marshal(user)
	if err != nil {
		h.logger.Printf("json marshal failed : %s", err)
	}
	fmt.Fprintln(w, string(jsonBytes))
}

func (h *User) create(w http.ResponseWriter, r *http.Request) {

	reqBody := struct {
		TestStr string `json:"test_str"`
	}{}
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res, err := json.Marshal(reqBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, string(res))
}
