package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New();

	app.Get("/health", func (c *fiber.Ctx) error {
		return c.SendString("The server is running.");
	});

	app.Listen(":8080");

}