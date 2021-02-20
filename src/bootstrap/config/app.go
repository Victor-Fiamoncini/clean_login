package config

import (
	"github.com/gofiber/fiber/v2"
)

// NewApp func
func NewApp() *fiber.App {
	app := fiber.New()

	SetupApp(app)

	return app
}
