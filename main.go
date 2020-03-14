package main

import (
	"github.com/gin-gonic/gin"
	"go-api/db"
)

func main() {
	db.ConnectDatabase()
	defer db.CloseDatabase()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	err := r.Run()
	if err != nil {
		panic("Can't run server")
	}
}
