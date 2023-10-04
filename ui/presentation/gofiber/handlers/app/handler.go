package app

import "github.com/gofiber/fiber/v2"

func InitAppRoutes(app *fiber.App) {
	app.Get("/health", health)
}

func health(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"health": "ok",
	})
}
