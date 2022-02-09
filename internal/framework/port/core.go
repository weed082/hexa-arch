package port

import "net/http"

type WorkerPool interface {
	Generate(count int)
	Stop()
	RegisterJob(callback func())
}

type Router interface {
	Use(...func(http.Handler) http.Handler)
	Get(path string, handleFunc http.HandlerFunc)
	Post(path string, handleFunc http.HandlerFunc)
	Put(path string, handleFunc http.HandlerFunc)
	Delete(path string, handleFunc http.HandlerFunc)
	ServeHTTP(rw http.ResponseWriter, r *http.Request)
}
