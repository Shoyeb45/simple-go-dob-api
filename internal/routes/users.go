package routes

import (
	"github.com/Shoyeb45/simple-go-dob-api/internal/handler"
	"github.com/gofiber/fiber/v2"
)

func Register(app *fiber.App, h *handler.UserHandler) {
	app.Post("/users", h.CreateUser);
	app.Get("/users/:id", h.GetUser);
	app.Put("/users/:id", h.UpdateUser);
	app.Delete("users/:id", h.DeleteUser);
}