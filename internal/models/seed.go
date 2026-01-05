package models

import (
	"log"

	"indonesia-holiday-api/internal/config"
)

func SeedHolidays() {
	var count int
	config.DB.QueryRow("SELECT COUNT(*) FROM holidays").Scan(&count)

	if count > 0 {
		log.Println("Holidays data already exists")
		return
	}

	query := `
	INSERT INTO holidays (date, name, type) VALUES
	('2025-01-01', 'Tahun Baru Masehi', 'national'),
	('2025-02-10', 'Isra Miraj Nabi Muhammad SAW', 'national'),
	('2025-03-29', 'Hari Raya Nyepi', 'national'),
	('2025-04-10', 'Idul Fitri 1446 H', 'national'),
	('2025-12-25', 'Hari Raya Natal', 'national');
	`

	_, err := config.DB.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Holidays seed inserted")
}
