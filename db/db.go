package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func ConnectDatabase() *sqlx.DB {
	db, err := sqlx.Connect("postgres", "host=localhost port=5432 user=goapidbadmin dbname=goapidb password=s3cr3t sslmode=disable")
	if err != nil {
		fmt.Println(err)
		panic("Can't connect to the database")
	}
	DB = db
	return DB
}

func GetDatabase() *sqlx.DB {
	return DB
}

func CloseDatabase() {
	err := DB.Close()
	if err != nil {
		panic("Can't close connection to the database")
	}
}
