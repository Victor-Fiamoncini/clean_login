package routes

import (
	"os"

	auth_usecase "github.com/Victor-Fiamoncini/auth_clean_architecture/src/domain/usecases/auth_usecase"
	"github.com/Victor-Fiamoncini/auth_clean_architecture/src/infra/database"
	luber "github.com/Victor-Fiamoncini/auth_clean_architecture/src/infra/repositories/load_user_by_email_repository"
	uatr "github.com/Victor-Fiamoncini/auth_clean_architecture/src/infra/repositories/update_access_token_repository"
	"github.com/Victor-Fiamoncini/auth_clean_architecture/src/presentation/helpers"
	"github.com/Victor-Fiamoncini/auth_clean_architecture/src/presentation/routers"
	encrypter "github.com/Victor-Fiamoncini/auth_clean_architecture/src/shared/helpers/encrypter"
	token_generator "github.com/Victor-Fiamoncini/auth_clean_architecture/src/shared/helpers/token_generator"
	"github.com/Victor-Fiamoncini/auth_clean_architecture/src/shared/validators"
	"github.com/gofiber/fiber/v2"
)

// LoginRoutes func
func LoginRoutes(router fiber.Router) {
	tokenSecret := os.Getenv("TOKEN_SECRET")

	userModel := database.GetCollection("users")

	router.Post("/login", func(c *fiber.Ctx) error {
		httpRequest := helpers.NewHTTPRequest()

		err := c.BodyParser(httpRequest.Body)

		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"status":  "error",
				"message": "Invalid params are provided",
			})
		}

		loadUserByEmailRepository := luber.NewLoadUserByEmailRepository(userModel)
		encrypter := encrypter.NewEncrypter()
		tokenGenerator := token_generator.NewTokenGenerator(tokenSecret)
		updateAccessTokenRepository := uatr.NewUpdateAccessTokenRepository(userModel)

		authUseCase := auth_usecase.NewAuthUseCase(
			loadUserByEmailRepository,
			encrypter,
			tokenGenerator,
			updateAccessTokenRepository,
		)

		emailValidator := validators.NewEmailValidator()

		loginRouter := routers.NewLoginRouter(authUseCase, emailValidator)

		httpResponse := loginRouter.Route(httpRequest)

		if httpResponse.GetErrorObject() != nil {
			return c.Status(httpResponse.GetStatusCode()).JSON(fiber.Map{
				"status":  "error",
				"message": httpResponse.GetErrorObject().Error(),
			})
		}

		return c.Status(200).JSON(httpResponse.GetBody())
	})
}
