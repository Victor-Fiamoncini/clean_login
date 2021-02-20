package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
)

// NewCompress func
func NewCompress() fiber.Handler {
	return compress.New()
}
