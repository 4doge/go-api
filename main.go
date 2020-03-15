package main

import (
	"go-api/config"
	"go-api/db"
	"go-api/server"
)

func main() {
	config.Init("dev")

	db.ConnectDatabase()
	defer db.CloseDatabase()

	server.Serve()
}
