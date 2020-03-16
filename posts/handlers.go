package posts

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListPosts(c *gin.Context) {
	posts := GetAllPosts()
	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})

}

func CreatePost(c *gin.Context) {
	var post Post
	_ = c.Bind(&post)
	id := CreateNewPost(post.Title, post.Body)
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"id": id,
	})
}

func RetrievePost(c *gin.Context) {
	post := GetSinglePost(c.Param("id"))
	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func UpdatePost(c *gin.Context) {
	var post Post
	_ = c.Bind(&post)
	UpdateSinglePost(c.Param("id"), post.Title, post.Body)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func RemovePost(c *gin.Context) {
	RemoveSinglePost(c.Param("id"))
	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}
