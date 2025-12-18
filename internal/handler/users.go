package handler

import (
	"github.com/Shoyeb45/simple-go-dob-api/internal/core"
	"github.com/Shoyeb45/simple-go-dob-api/internal/models"
	"github.com/Shoyeb45/simple-go-dob-api/internal/service"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	service *service.UserService
};



// Handler to create a user
func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	body, err := core.ParseBody[models.UserCreate](c);

	if err != nil {
		return err;
	}

	createdUser, err := h.service.CreateUser(c.Context(), body);

	if err != nil {
		return err;
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data": createdUser,
	});
}

// Get particular user and also output the age. 
func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	id := c.Params("id");
	
	user, err := h.service.GetUser(c.Context(), &id);

	if err != nil {
		return err;
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": user, 
	});
}

// Handler to handle the update of the user.
func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id");

	body, err := core.ParseBody[models.UserCreate](c);

	if err != nil {
		return err;
	}

	updatedUser, err := h.service.UpdateUser(c.Context(), &id, body);

	if err != nil {
		return err;
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": updatedUser,
	});
}

// Handler to delete the user
func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id");

	err := h.service.DeleteUser(c.Context(), &id);

	if err != nil {
		return err;
	}

	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{});
}

// Get all the users
func (h *UserHandler) GetAllUsers(c *fiber.Ctx) error {
	users, err := h.service.GetAllUsers(c.Context());

	if err != nil {
		return err;
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": users,
	});
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service};
}