package main

import (
	"indonesia-holiday-api/internal/config"
	"indonesia-holiday-api/internal/models"
	"indonesia-holiday-api/internal/routes"

	"github.com/gofiber/fiber/v2"
)


func main() {
	app := fiber.New()

	config.ConnectDB()
	models.Migrate()

	routes.Register(app)

	app.Listen(":3000")
}