package posts

import (
	"go-api/db"
)

type Post struct {
	// TODO: check if we need ID here, also investigate about bindings
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func GetAllPosts() ([]Post, error) {
	database := db.GetDatabase()
	var posts []Post
	query := "SELECT id, title, body FROM posts"
	err := database.Select(&posts, query)
	return posts, err
}

func CreateNewPost(title string, body string) (Post, error) {
	database := db.GetDatabase()
	var post Post
	query := "INSERT INTO posts (title, body) VALUES ($1, $2)"
	database.MustExec(query, title, body)
	// TODO: check the result and return a post instance and an error if any
	return post, nil
}

func GetSinglePost(id string) (Post, error) {
	database := db.GetDatabase()
	var post Post
	query := "SELECT id, title, body FROM posts WHERE id=$1"
	err := database.Get(&post, query, id)
	return post, err
}

func UpdateSinglePost(title string, body string, id string) (Post, error) {
	// TODO: check the result and return a post instance and an error if any
	database := db.GetDatabase()
	var post Post
	query := "UPDATE posts SET title=$1, body=$2 WHERE id=$3"
	database.MustExec(query, title, body, id)
	return post, nil
}
