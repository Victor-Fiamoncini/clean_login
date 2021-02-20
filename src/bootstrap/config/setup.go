package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// SetupApp func
func SetupApp(app *fiber.App) {
	app.Use(logger.New())
	app.Use(cors.New())
}
