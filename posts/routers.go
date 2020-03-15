package posts

import "github.com/gin-gonic/gin"

func RegisterPostsRoutes(group *gin.RouterGroup) {
	group.GET("/", ListPosts)
	group.POST("/", CreatePost)
	group.GET("/:id", RetrievePost)
	group.PUT("/:id", UpdatePost)
	group.DELETE("/:id", RemovePost)
}
