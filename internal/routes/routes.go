package routes

import "github.com/gofiber/fiber/v2"

func Register(app *fiber.App) {
	app.Get("/api/holidays", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"data": []string{},
		})
	})
}