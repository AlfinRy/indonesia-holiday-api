package models

import (
	"log"

	"indonesia-holiday-api/internal/config"
)

func Migrate() {
	query := `
	CREATE TABLE IF NOT EXISTS holidays (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		date TEXT,
		name TEXT,
		type TEXT
	);
	`

	_, err := config.DB.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}