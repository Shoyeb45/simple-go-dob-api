package routes

import (
	"github.com/Shoyeb45/simple-go-dob-api/internal/handler"
	"github.com/gofiber/fiber/v2"
)

func RegisterUserRoutes(app *fiber.App, h *handler.UserHandler) {
	// Create a user
	app.Post("/users", h.CreateUser);
	// Get all the users
	app.Get("/users", h.GetAllUsers);
	// Get user by id
	app.Get("/users/:id", h.GetUser);
	app.Put("/users/:id", h.UpdateUser);
	app.Delete("users/:id", h.DeleteUser);
}