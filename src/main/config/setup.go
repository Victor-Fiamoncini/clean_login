package config

import (
	"github.com/Victor-Fiamoncini/auth_clean_architecture/src/main/middlewares"
	"github.com/gofiber/fiber/v2"
)

// SetupApp func
func SetupApp(app *fiber.App) {
	app.Use(middlewares.NewHelmet())
	app.Use(middlewares.NewCors())
	app.Use(middlewares.NewCompress())
	app.Use(middlewares.NewLogger())
}
