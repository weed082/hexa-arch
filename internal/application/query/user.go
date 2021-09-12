package query

import (
	"github.com/ByungHakNoh/hexagonal-microservice/internal/framework/port"
)

type User struct {
	mongoDB port.Repository
}

func NewUserQuery(mongoDB port.Repository) *User {
	return &User{mongoDB: mongoDB}
}

func (query User) GetUserList() {

}
