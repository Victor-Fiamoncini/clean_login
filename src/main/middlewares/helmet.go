package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/helmet/v2"
)

// NewHelmet func
func NewHelmet() fiber.Handler {
	return helmet.New()
}
