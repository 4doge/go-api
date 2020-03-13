package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	// Database
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=goapidbadmin dbname=goapidb password=s3cr3t sslmode=disable")
	if err != nil {
		fmt.Println(err)
		panic("Failed to connect database")
	}
	defer db.Close()
	//
	// // CRUD
	// type Product struct {
	// 	gorm.Model
	// 	Code  string
	// 	Price uint
	// }
	//
	// db.AutoMigrate(&Product{})
	// db.Create(&Product{Code: "example", Price: 123})

	// Server
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
