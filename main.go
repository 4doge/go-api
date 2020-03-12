package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	// Database
	// db, err := gorm.Open("sqlite3", "test.db")
	// if err != nil {
	// 	panic("failed to connect database")
	// }
	// defer db.Close()
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
