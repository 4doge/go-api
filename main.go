package main

import (
	"go-api/config"
	"go-api/db"
	"go-api/posts"
	"go-api/server"
)

func main() {
	config.Init("dev")

	db.ConnectDatabase()
	db.GetDatabase().AutoMigrate(&posts.Post{})
	defer db.CloseDatabase()

	server.Serve()
}
