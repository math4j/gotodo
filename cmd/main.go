package main

import (
	"github.com/math4j/gotodo/internal/config"
	"github.com/math4j/gotodo/internal/core/server"
)

func main() {

	config.SetUpLog()
	server := server.NewServer(server.WithPort(3000))
	server.Run()
}
