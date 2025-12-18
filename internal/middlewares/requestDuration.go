package middlewares

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func RequestDurationLogger(log *zap.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		err := c.Next()

		duration := time.Since(start)
		requestId := c.GetRespHeader("X-Request-Id");

		log.Info("http request",
			zap.String("method", c.Method()),
			zap.String("path", c.OriginalURL()),
			zap.Int("status", c.Response().StatusCode()),
			zap.Duration("duration", duration),
			zap.Any("requestId", requestId),
		);
		return err;
	}
}
