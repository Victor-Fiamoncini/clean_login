package routes

import (
	"github.com/Victor-Fiamoncini/auth_clean_architecture/src/bootstrap/factories"
	"github.com/Victor-Fiamoncini/auth_clean_architecture/src/presentation/contracts/payloads"
	"github.com/Victor-Fiamoncini/auth_clean_architecture/src/presentation/http"
	"github.com/gofiber/fiber/v2"
)

// LoginRoutes func
func LoginRoutes(router fiber.Router) {
	router.Post("/login", func(c *fiber.Ctx) error {
		httpRequest := http.NewRequest()
		httpRequest.SetBody(payloads.NewLoginPayload())

		err := c.BodyParser(httpRequest.GetBody())

		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"status":  "error",
				"message": "Invalid params are provided",
			})
		}

		httpResponse := factories.MakeLoginController().Handle(httpRequest)

		if httpResponse.GetErrorObject() != nil {
			return c.Status(httpResponse.GetStatusCode()).JSON(fiber.Map{
				"status":  "error",
				"message": httpResponse.GetErrorObject().Error(),
			})
		}

		return c.Status(200).JSON(httpResponse.GetBody())
	})
}
