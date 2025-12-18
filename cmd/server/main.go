package main

import (
	"github.com/Shoyeb45/simple-go-dob-api/config"
	"github.com/Shoyeb45/simple-go-dob-api/internal/app"
	"github.com/Shoyeb45/simple-go-dob-api/internal/database"
	"github.com/Shoyeb45/simple-go-dob-api/internal/logger"
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

	// initialise the database
	err := database.Connect();

	if err != nil {
		logger.Log.Fatal("Database connection failed", zap.Error(err))
	}
	defer database.Close();

	// instatiate the fiber app
	app := app.New(database.DB);

	logger.Log.Info("Starting server on port " + config.Cfg.PORT)

	// start the app
	err = app.Fiber.Listen(":" + config.Cfg.PORT)

	if err != nil {
		logger.Log.Fatal("Error occurred while starting server on port 8080.", zap.Error(err))
		panic(err)
	}
}



