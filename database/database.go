package database

import (
	"database/sql"
	"fmt"
	"gobot/config"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Init() {
	dbURL := config.GetString("db.url")
	var err error
	DB, err = sql.Open("postgres", dbURL)
	if err != nil {
		panic(err)
	}

	err = DB.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to database!")
}
