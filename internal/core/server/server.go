package server

import (
	"fmt"
	"log/slog"

	"github.com/labstack/echo/v4"
	"github.com/math4j/gotodo/internal/config"
	"github.com/math4j/gotodo/internal/core/application"
	"github.com/math4j/gotodo/internal/infra/repository/mysql"
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

	db := config.NewDB()
	repository := mysql.NewTaskRepository(db)
	health := application.NewHealth()
	task := application.NewTask(repository)
	app := echo.New()

	app.GET("/api/v1/health", health.GetHealth)

	app.GET("/api/v1/task", task.Get)
	app.GET("/api/v1/task/get-all", task.GetAll)
	app.POST("/api/v1/task/create", task.Create)
	app.DELETE("/api/v1/task/delete", task.Delete)

	slog.Info(fmt.Sprintf("server starting at port %s", s.Addr))
	app.Start(s.Addr)
}
