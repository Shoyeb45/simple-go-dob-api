package app

import (
	"github.com/Shoyeb45/simple-go-dob-api/internal/core"
	"github.com/Shoyeb45/simple-go-dob-api/internal/handler"
	"github.com/Shoyeb45/simple-go-dob-api/internal/logger"
	"github.com/Shoyeb45/simple-go-dob-api/internal/middlewares"
	"github.com/Shoyeb45/simple-go-dob-api/internal/repository"
	"github.com/Shoyeb45/simple-go-dob-api/internal/routes"
	"github.com/Shoyeb45/simple-go-dob-api/internal/service"

	sqlcDb "github.com/Shoyeb45/simple-go-dob-api/db/sqlc"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type App struct {
	Fiber *fiber.App
}

func New(db *pgxpool.Pool) *App {
	app := fiber.New(fiber.Config{
		ErrorHandler: core.ErrorHandler(logger.Log),
	});

	// middleware to add requestId.
	app.Use(requestid.New(requestid.Config{
		Header: "X-Request-Id",
	}));

	app.Use(middlewares.RequestDurationLogger(logger.Log));

	// Register health and root routes
	routes.RegisterHealthRoutes(app);

	queries := sqlcDb.New(db);

	// repository
	userRepo := repository.NewUserRepository(queries);

	// services
	userService := service.NewUserService(userRepo);

	// handler
	userHandler := handler.NewUserHandler(userService);

	// register routes
	routes.RegisterUserRoutes(app, userHandler);

	// Any route which is not found
	app.Use(func(c *fiber.Ctx) error {
		return core.NewNotFoundError("No route matched with given endpoint.");
	})

	return &App{Fiber: app};
}
