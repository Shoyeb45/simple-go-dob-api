package routes

import "github.com/gofiber/fiber/v2"

func RegisterHealthRoutes(app *fiber.App) {
	// Root route
	app.Get("/", func (c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "The server is running.",
		});
	});

	// Health route
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "The server is healthy and running fine.",
		})
	});
}
