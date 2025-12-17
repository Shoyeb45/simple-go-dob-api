package main

import (

	"github.com/Shoyeb45/simple-go-dob-api/config"
	"github.com/Shoyeb45/simple-go-dob-api/internal/logger"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func main() {
	// Initiliase environment variables
	configErr := config.LoadEnvironmentVariables();

	if configErr != nil {
		panic(configErr.Error());
	}
	// initialise the logger
	logger.Init(config.Cfg.APP_ENV);


	app := fiber.New()

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("The server is running.")
	})

	logger.Log.Info("Starting server on port " + config.Cfg.PORT)

	// start the app
	err := app.Listen(":" + config.Cfg.PORT)

	if err != nil {
		logger.Log.Fatal("Error occurred while starting server on port 8080.", zap.Error(err))
		panic(err)
	}
}
