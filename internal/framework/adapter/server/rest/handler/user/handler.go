package user

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/adapter/server/rest/handler"
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/port"
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

func (h *User) Register(r handler.Router) {
	r.Post("/user/create", h.create)
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
