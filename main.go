package main

import (
	"github.com/gin-gonic/gin"
	"go-api/db"
)

func main() {
	db.ConnectDatabase()
	defer db.CloseDatabase()

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
