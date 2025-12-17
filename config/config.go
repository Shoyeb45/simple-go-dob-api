package config

import (
	"errors"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	APP_ENV string
	PORT    string
	DB_URL  string
}

var Cfg *Config

func LoadEnvironmentVariables() error {
	err := godotenv.Load(".env")

	if err != nil {
		return errors.New("error loading .env file")
	}

	Cfg = &Config{
		APP_ENV: getEnv("APP_ENV", "development"),
		PORT:    getEnv("PORT", "8080"),
		DB_URL:  getEnv("DB_URL", ""),
	}

	return validate()
}

func validate() error {
	port, err := strconv.Atoi(Cfg.PORT)

	if err != nil {
		return errors.New("PORT must be a number, PORT: " + Cfg.PORT)
	}

	if port < 0 || port > 65535 {
		return errors.New("PORT out of valid range, expected between min : 1, max: 65535, got " + Cfg.PORT)
	}

	if Cfg.DB_URL == "" {
		return errors.New("DB_URL cannot be empty");
	}

	if Cfg.APP_ENV != "production" && Cfg.APP_ENV != "development" {
		return errors.New("Unknown APP_ENV received, only production or development is valid, APP_ENV: " + Cfg.APP_ENV);
	}
	return nil;
}

func getEnv(key string, defaultVal string) string {
	val, ok := os.LookupEnv(key)

	if ok {
		return val
	}
	return defaultVal
}
