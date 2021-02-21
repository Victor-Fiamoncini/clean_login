package config

import (
	"github.com/gofiber/fiber/v2"
)

// NewApp func
func NewApp() *fiber.App {
	LoadEnv()

	app := fiber.New()

	SetupApp(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})

	return app
}
