package app

import (
	"github.com/Shoyeb45/simple-go-dob-api/internal/core"
	"github.com/Shoyeb45/simple-go-dob-api/internal/handler"
	"github.com/Shoyeb45/simple-go-dob-api/internal/logger"
	"github.com/Shoyeb45/simple-go-dob-api/internal/repository"
	"github.com/Shoyeb45/simple-go-dob-api/internal/routes"
	"github.com/Shoyeb45/simple-go-dob-api/internal/service"

	sqlcDb "github.com/Shoyeb45/simple-go-dob-api/db/sqlc"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

type App struct {
	Fiber *fiber.App
}

func New (db *pgxpool.Pool) *App {
	app := fiber.New(fiber.Config{
		ErrorHandler: core.ErrorHandler(logger.Log),
	});

	app.Get("/health", func(c *fiber.Ctx) error {
		// return c.SendString("The server is running.")
		return core.NewInternalError("Test");
	});

	queries := sqlcDb.New(db);

	// repository 
	userRepo := repository.NewUserRepository(queries);

	// services
	userService := service.NewUserService(userRepo);

	// handler
	userHandler := handler.NewUserHandler(userService);

	// register routes
	routes.RegisterUserRoutes(app, userHandler);


	return &App{Fiber: app};
}
