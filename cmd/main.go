package main

import (
	"indonesia-holiday-api/internal/routes"

	"github.com/gofiber/fiber/v2"
)


func main() {
	app := fiber.New()

	routes.Register(app)
	
	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.JSON(fiber.Map{
	// 		"message": "Indonesia Holiday API is running",
	// 	})
	// })

	app.Listen(":3000")
}