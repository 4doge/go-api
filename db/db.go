package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func ConnectDatabase() *gorm.DB {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=goapidbadmin dbname=goapidb password=s3cr3t sslmode=disable")
	if err != nil {
		fmt.Println(err)
		panic("Can't connect to the database")
	}
	DB = db
	return DB
}

func GetDatabase() *gorm.DB {
	return DB
}

func CloseDatabase() {
	err := DB.Close()
	if err != nil {
		panic("Can't close connection to the database")
	}
}
