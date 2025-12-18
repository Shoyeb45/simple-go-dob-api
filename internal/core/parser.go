package core

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ParseBody[T any](c *fiber.Ctx) (*T, error) {
	var body T;

	// parse the body
	err := c.BodyParser(&body);

	// handle the error using our AppError class.
	if err != nil {
		return nil, NewBadRequestError("JSON Parsing error, Invalid JSON.").WithInternal(err);
	}

	validate := validator.New();

	err = validate.Struct(body);

	if err != nil {
		return nil, NewBadRequestError("Validation failed.").WithInternal(err);
	}

	return &body, nil;
}