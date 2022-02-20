package server

import (
	"net/http"
	"time"
)

func NewServer(handler http.Handler, address string) *http.Server {
	return &http.Server{
		Addr:         address,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      handler,
	}
}
