package posts

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListPosts(c *gin.Context) {
	posts, _ := GetAllPosts()
	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})

}

func CreatePost(c *gin.Context) {
	// TODO: investigate how we can accept the JSON body and bind it to the Post struct
	// var post Post
	// err := c.Bind(&post)
	// post, _ := CreateNewPost("title", "body")
	// fmt.Println(post)
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
	})
}
func RetrievePost(c *gin.Context) {
	// TODO: check if post exist, return 404 if not
	post, _ := GetSinglePost(c.Param("id"))
	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func UpdatePost(c *gin.Context) {
	// TODO
	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func RemovePost(c *gin.Context) {
	// TODO
	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}
