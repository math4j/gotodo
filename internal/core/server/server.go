package server

import (
	"fmt"
	"log/slog"
)

type Opt func(*Server)

type Server struct {
	Port int
}

func NewServer(opts ...Opt) *Server {

	server := defaultConfig()

	for _, opt := range opts {
		opt(server)
	}

	return server
}

func defaultConfig() *Server {
	return &Server{
		Port: 8080,
	}
}

func WithPort(port int) func(server *Server) {

	return func(server *Server) {
		server.Port = port
	}
}

func (s *Server) Run() {

	slog.Info(fmt.Sprintf("server running at port %d", s.Port))
}
