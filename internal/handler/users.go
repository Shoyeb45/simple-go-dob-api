package handler

import (
	"github.com/Shoyeb45/simple-go-dob-api/internal/service"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	service *service.UserService
};


func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	return nil;
}

func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	return nil;
}

func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	return nil;

}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	return nil;
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service};
}