package database

import (
	"context"
	"time"

	"github.com/Shoyeb45/simple-go-dob-api/config"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/Shoyeb45/simple-go-dob-api/internal/logger"
)

var DB *pgxpool.Pool;

func Connect() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second);
	defer cancel();

	pool, err := pgxpool.New(ctx, config.Cfg.DB_URL);
	if err != nil {
		return err;
	}
	err = pool.Ping(ctx);
	if err != nil {
		return err;
	}

	DB = pool;
	
	logger.Log.Info("Connected to PostgreSQL database.");
	
	return nil;
}

func Close() {
	if DB != nil {
		DB.Close()
	}
}