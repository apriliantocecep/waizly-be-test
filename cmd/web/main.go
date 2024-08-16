package main

import (
	"waizly/internal/bootstrap"
	"waizly/internal/server"
)

func main() {
	app := server.InitializeServer()

	boot := bootstrap.NewBootstrap(app)
	boot.StartServer()
}
