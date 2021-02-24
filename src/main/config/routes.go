package config

import (
	"github.com/Victor-Fiamoncini/auth_clean_architecture/src/main/routes"
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes func
func SetupRoutes(app *fiber.App) {
	router := app.Group("/api")

	routes.LoginRoutes(router)
}
