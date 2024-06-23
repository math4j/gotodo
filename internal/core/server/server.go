package server

import (
	"fmt"
	"log/slog"

	"github.com/labstack/echo/v4"
	"github.com/math4j/gotodo/internal/core/application"
)

type Opt func(*Server)

type Server struct {
	Addr string
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
		Addr: ":8080",
	}
}

func WithPort(port int) func(server *Server) {

	addr := fmt.Sprintf(":%d", port)
	return func(server *Server) {
		server.Addr = addr
	}
}

func (s *Server) Run() {

	health := application.NewHealth()
	app := echo.New()

	app.GET("/api/v1/health", health.GetHealth)

	slog.Info(fmt.Sprintf("server starting at port %s", s.Addr))
	app.Start(s.Addr)
}
