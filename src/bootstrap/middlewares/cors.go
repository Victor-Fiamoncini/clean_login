package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// NewCors func
func NewCors() fiber.Handler {
	return cors.New()
}
