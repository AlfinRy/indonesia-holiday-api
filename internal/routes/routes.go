package routes

import (
	"indonesia-holiday-api/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func Register(app *fiber.App) {
	app.Get("/api/holidays", handlers.GetHolidays)
}