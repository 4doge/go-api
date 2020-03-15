package server

import (
	"github.com/gin-gonic/gin"
	"go-api/posts"
)

func Serve() {
	router := gin.Default()

	posts.RegisterPostsRoutes(router.Group("/posts"))

	err := router.Run()
	if err != nil {
		panic("Can't run server")
	}
}
