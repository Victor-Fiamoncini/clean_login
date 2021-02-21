package routes

import "github.com/gofiber/fiber/v2"

// LoginRoutes func
func LoginRoutes(router fiber.Router) {
	router.Post("/api/login", func(c *fiber.Ctx) error {
	})
}
