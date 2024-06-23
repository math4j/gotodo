package main

import "github.com/math4j/gotodo/internal/core/server"

func main() {

	server := server.NewServer(server.WithPort(3000))
	server.Run()
}
