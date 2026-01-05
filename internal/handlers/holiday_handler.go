package handlers

import (
	"database/sql"
	"indonesia-holiday-api/internal/config"
	"indonesia-holiday-api/internal/models"

	"github.com/gofiber/fiber/v2"
)

func GetHolidays(c *fiber.Ctx) error {
	rows, err := config.DB.Query(
		"SELECT id, date, name, type FROM holidays ORDER BY date ASC",
	)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	defer rows.Close()

	holidays := []models.Holiday{}
	
	for rows.Next() {
		var h models.Holiday
		err := rows.Scan(&h.ID, &h.Date, &h.Name, &h.Type);
		if err != nil {
			if err == sql.ErrNoRows {
				continue
			}
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
		holidays = append(holidays, h)
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data": holidays,
	})
}
