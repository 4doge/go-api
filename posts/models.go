package posts

import (
	"github.com/jinzhu/gorm"
	"go-api/db"
)

type Post struct {
	gorm.Model
	Title string `json:"title"`
	Body  string `json:"body"`
}

func GetAllPosts() []Post {
	database := db.GetDatabase()
	var posts []Post
	database.Select("id, title, body").Find(&posts)
	return posts
}

func CreateNewPost(title string, body string) uint {
	database := db.GetDatabase()
	post := Post{
		Title: title,
		Body:  body,
	}
	database.Create(&post)
	return post.ID
}

func GetSinglePost(id string) Post {
	database := db.GetDatabase()
	var post Post
	database.First(&post, id)
	return post
}

func UpdateSinglePost(id string, title string, body string) {
	database := db.GetDatabase()
	post := GetSinglePost(id)
	post.Title = title
	post.Body = body
	database.Save(&post)
}

func RemoveSinglePost(id string) {
	database := db.GetDatabase()
	post := GetSinglePost(id)
	database.Unscoped().Delete(&post)
}
