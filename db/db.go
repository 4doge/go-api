package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go-api/config"
)

var DB *sqlx.DB

func ConnectDatabase() *sqlx.DB {
	c := config.GetConfig()
	// TODO: refactor this crazy connection to database string
	dataSource := "host=" + c.GetString("database.host") + " port=" + c.GetString("database.port") + " user=" + c.GetString("database.user") + " dbname=" + c.GetString("database.name") + " password=" + c.GetString("database.password") + " sslmode=disable"
	db, err := sqlx.Connect("postgres", dataSource)
	if err != nil {
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
