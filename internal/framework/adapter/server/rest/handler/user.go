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

// testing the post req, res
type SomeStruct struct {
	Test string
}

func (h *User) create(w http.ResponseWriter, r *http.Request) {
	var someStruct SomeStruct
	err := json.NewDecoder(r.Body).Decode(&someStruct)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res, err := json.Marshal(someStruct)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, string(res))
}
