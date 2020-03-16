package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"go-api/config"
)

var DB *gorm.DB

func ConnectDatabase() *gorm.DB {
	c := config.GetConfig()
	dataSource := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		c.GetString("database.host"),
		c.GetString("database.port"),
		c.GetString("database.user"),
		c.GetString("database.name"),
		c.GetString("database.password"),
	)
	db, err := gorm.Open("postgres", dataSource)
	if err != nil {
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
