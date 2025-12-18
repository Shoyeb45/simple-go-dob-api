package core

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type ErrorCode string

const (
	ErrCodeValidation   ErrorCode = "VALIDATION_ERROR"
	ErrCodeNotFound     ErrorCode = "NOT_FOUND"
	ErrCodeUnauthorized ErrorCode = "UNAUTHORIZED"
	ErrCodeBadRequest   ErrorCode = "BAD_REQUEST_ERROR"
	ErrCodeInternal     ErrorCode = "INTERNAL_SERVER_ERROR"
)

// AppError represents a structured application error.
type AppError struct {
	Code       ErrorCode              `json:"code"`
	Message    string                 `json:"message"`
	StatusCode int                    `json:"-"`
	Details    map[string]interface{} `json:"details,omitempty"`
	Internal   error                  `json:"-"`
}

func (e *AppError) Error() string {
	if e.Internal != nil {
		return fmt.Sprintf("[%s] %s: %v", e.Code, e.Message, e.Internal)
	}
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

func (e *AppError) WithDetails(key string, value interface{}) *AppError {
	if e.Details == nil {
		e.Details = make(map[string]interface{})
	}
	e.Details[key] = value

	return e
}

// WithInternal adds the underlying error for logging
func (e *AppError) WithInternal(err error) *AppError {
	e.Internal = err
	return e
}

// for logging structured error.
func (e *AppError) LogFields() []zap.Field {
	fields := []zap.Field{
		zap.String("error_code", string(e.Code)),
		zap.String("message", e.Message),
		zap.Int("status_code", e.StatusCode),
	}

	if e.Internal != nil {
		fields = append(fields, zap.Error(e.Internal))
	}

	if len(e.Details) > 0 {
		fields = append(fields, zap.Any("details", e.Details))
	}

	return fields
}

func NewValidationError(message string) *AppError {
	return &AppError{
		Code:       ErrCodeValidation,
		Message:    message,
		StatusCode: http.StatusBadRequest,
	}
}

func NewNotFoundError(message string) *AppError {
	return &AppError{
		Code:       ErrCodeNotFound,
		Message:    message,
		StatusCode: http.StatusNotFound,
	}
}

func NewUnauthorizedError(message string) *AppError {
	return &AppError{
		Code:       ErrCodeNotFound,
		Message:    message,
		StatusCode: http.StatusUnauthorized,
	}
}

func NewInternalError(message string) *AppError {
	return &AppError{
		Code:       ErrCodeInternal,
		Message:    message,
		StatusCode: http.StatusInternalServerError,
	}
}

func NewBadRequestError(message string) *AppError {
	return &AppError{
		Code: ErrCodeBadRequest,
		Message: message,
		StatusCode: http.StatusBadRequest,
	};
}

// Middleware to handle the error
func ErrorHandler(logger *zap.Logger) fiber.ErrorHandler {
	return func(c *fiber.Ctx, err error) error {
		// Default to 500 Internal Server Error
		code := fiber.StatusInternalServerError
		response := fiber.Map{
			"success": false,
			"error": fiber.Map{
				"code":    ErrCodeInternal,
				"message": "Internal server error",
			},
		}

		// Check if it's an AppError
		if appErr, ok := err.(*AppError); ok {
			code = appErr.StatusCode
			response["error"] = fiber.Map{
				"code":    appErr.Code,
				"message": appErr.Message,
				"details": appErr.Details,
			}

			// Log with structured fields
			logLevel := getLogLevel(appErr.StatusCode)
			fields := append(appErr.LogFields(),
				zap.String("path", c.Path()),
				zap.String("method", c.Method()),
				zap.String("ip", c.IP()),
			)

			switch logLevel {
			case "error":
				logger.Error("Application error", fields...)
			case "warn":
				logger.Warn("Application warning", fields...)
			default:
				logger.Info("Application error", fields...)
			}

		} else if fiberErr, ok := err.(*fiber.Error); ok {
			// Handle Fiber's built-in errors
			code = fiberErr.Code
			response["error"] = fiber.Map{
				"code":    ErrCodeBadRequest,
				"message": fiberErr.Message,
			}

			logger.Warn("Fiber error",
				zap.Int("status_code", fiberErr.Code),
				zap.String("message", fiberErr.Message),
				zap.String("path", c.Path()),
				zap.String("method", c.Method()),
			)
		} else {
			// Log unexpected errors
			logger.Error("Unexpected error",
				zap.Error(err),
				zap.String("path", c.Path()),
				zap.String("method", c.Method()),
				zap.String("ip", c.IP()),
			)
		}

		// Send JSON response
		return c.Status(code).JSON(response)
	}
}

// getLogLevel determines the appropriate log level based on status code
func getLogLevel(statusCode int) string {
	switch {
	case statusCode >= 500:
		return "error"
	case statusCode >= 400:
		return "warn"
	default:
		return "info"
	}
}
