package handlers

import (
	"indonesia-holiday-api/internal/config"
	"indonesia-holiday-api/internal/models"

	"github.com/gofiber/fiber/v2"
)

func GetHolidays(c *fiber.Ctx) error {
	year := c.Query("year")
	month := c.Query("month")

	query := "SELECT id, date, name, type FROM holidays WHERE 1=1"
	args := []interface{}{}

	if year != "" {
		query += " AND strftime('%Y', date) = ?"
		args = append(args, year)
	}

	if month != "" {
		query += " AND strftime('%m', date) = ?"
		if len(month) == 1 {
			month = "0" + month
		}
		args = append(args, month)
	}

	query += " ORDER BY date ASC"

	rows, err := config.DB.Query(query, args...)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	defer rows.Close()

	holidays := []models.Holiday{}
	
	for rows.Next() {
		var h models.Holiday
		if err := rows.Scan(&h.ID, &h.Date, &h.Name, &h.Type); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})	
		}
		holidays = append(holidays, h)
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data": holidays,
	})
}

func GetNextHoliday(c *fiber.Ctx) error {
	query := `
		SELECT id, date, name, type
		FROM holidays
		WHERE date >= date('now')
		ORDER BY id date ASC
		LIMIT 1
	`

	row := config.DB.QueryRow(query)

	var h models.Holiday
	err := row.Scan(&h.ID, &h.Date, &h.Name, &h.Type)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "No upcoming holiday found",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data": h,
	})
}
