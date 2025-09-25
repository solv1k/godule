package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func NewApiLimiter(limit int, duration time.Duration) fiber.Handler {
	return limiter.New(limiter.Config{
		Max:        limit,
		Expiration: duration,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},
	})
}
