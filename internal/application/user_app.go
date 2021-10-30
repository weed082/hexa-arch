package application

import "github.com/ByungHakNoh/hexagonal-microservice/internal/framework/port"

type UserApp struct {
	repo port.UserRepo
}

func NewUserApp(repo port.UserRepo) *UserApp {
	return &UserApp{repo}
}

func (app *UserApp) Register() {

}

func (app *UserApp) Signin() {

}
