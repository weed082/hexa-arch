package port

import "net/http"

type UserApp interface {
	Register(rw http.ResponseWriter, r *http.Request)
	Signin(rw http.ResponseWriter, r *http.Request)
}
