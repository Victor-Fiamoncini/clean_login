package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/helmet/v2"
)

// SetupApp func
func SetupApp(app *fiber.App) {
	app.Use(helmet.New())
	app.Use(cors.New())
	app.Use(compress.New())
	app.Use(logger.New())
}
