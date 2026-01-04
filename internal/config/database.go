package config

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func ConnectDB() {
	db, err := sql.Open("sqlite", "holidays.db")
	if err != nil {
		log.Fatal(err)
	}

	DB = db
	log.Println("Database connected")
}