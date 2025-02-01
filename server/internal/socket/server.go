package socket

import "sync"

type Server struct {
	clients map[string]*Client
	rooms   map[string]*Room
	mut     sync.Mutex
}

func NewServer() *Server {
	return &Server{
		clients: make(map[string]*Client),
		rooms:   make(map[string]*Room),
	}
}
