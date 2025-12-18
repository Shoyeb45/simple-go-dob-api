package handler

import (
	"strconv"
	"time"

	"github.com/Shoyeb45/simple-go-dob-api/internal/core"
	"github.com/Shoyeb45/simple-go-dob-api/internal/mapper"
	"github.com/Shoyeb45/simple-go-dob-api/internal/models"
	"github.com/Shoyeb45/simple-go-dob-api/internal/service"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	service *service.UserService
};




func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	body, err := core.ParseBody[models.UserCreate](c);

	if err != nil {
		return err;
	}
	
	// parse the dob manually
	dob, err := time.Parse("2006-01-02", body.Dob);

	if err != nil {
		return core.NewBadRequestError("Invalid Date of Birth Provided.").WithInternal(err);
	}

	createdUser, err := h.service.CreateUser(c.Context(), body.Name, dob);

	if err != nil {
		return core.NewInternalError("Failed to create user.").WithInternal(err);
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data": mapper.UserToResponse(createdUser),
	});
}


func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	id := c.Params("id");

	idInNum, err := strconv.ParseInt(id, 10, 64);

	if err != nil {
		return core.NewBadRequestError("Not valid id given, id: " + id).WithInternal(err);
	}

	user, err := h.service.GetUser(c.Context(), idInNum);

	if err != nil {
		return core.NewNotFoundError("No user found with given ID.").WithInternal(err).WithDetails("id", id);
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": mapper.UserToWithAgeResponse(user), 
	});
}

func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id");

	idInNum, err := strconv.ParseInt(id, 10, 64);

	if err != nil {
		return core.NewBadRequestError("Not valid id given").WithInternal(err).WithDetails("id", id);
	}

	body, err := core.ParseBody[models.UserCreate](c);

	if err != nil {
		return err;
	}

	// first check if the user is present or not
	_, err = h.service.GetUser(c.Context(), idInNum);

	if err != nil {
		return core.NewNotFoundError("No user is present with given ID.").WithDetails("id", id).WithInternal(err);
	}
	updatedUser, err := h.service.UpdateUser(c.Context(), idInNum, *body);

	if err != nil {
		return err;
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": mapper.UserToResponse(updatedUser),
	});
}


func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id");

	idInNum, err := strconv.ParseInt(id, 10, 64);

	if err != nil {
		return core.NewBadRequestError("Not valid id given, id: " + id).WithInternal(err);
	}

	// Verify if user exists 
	_, err = h.service.GetUser(c.Context(), idInNum);

	if err != nil {
		return core.NewNotFoundError("No User found with given Id.").WithDetails("id", id).WithInternal(err);
	}

	err = h.service.DeleteUser(c.Context(), idInNum);

	if err != nil {
		return core.NewInternalError("Failed to delete the user.").WithInternal(err);
	}

	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{});
}


func (h *UserHandler) GetAllUsers(c *fiber.Ctx) error {
	users, err := h.service.GetAllUsers(c.Context());

	if err != nil {
		return core.NewInternalError("Failed to retrieve users.").WithInternal(err);
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": mapper.UsersToResponse(users),
	});
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service};
}