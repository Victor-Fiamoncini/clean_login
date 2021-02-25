package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// NewLogger func
func NewLogger() fiber.Handler {
	return logger.New()
}
