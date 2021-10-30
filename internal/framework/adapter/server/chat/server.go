package chat

import (
	"log"
	"net"
)

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (server *Server) Run(port string) {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("unable to start server: %s", err.Error())
	}
	defer listener.Close()
	log.Printf("started server on: %s", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("unable to accept connection: %s", err.Error())
			continue
		}
		server.newClient(conn)
	}
}

func (server *Server) newClient(conn net.Conn) {
	log.Printf("new client has connected: %s", conn.RemoteAddr().String())
}
