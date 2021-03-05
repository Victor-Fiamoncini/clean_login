package factories

import (
	"os"

	auth_usecase "github.com/Victor-Fiamoncini/auth_clean_architecture/src/domain/usecases/auth_usecase"
	"github.com/Victor-Fiamoncini/auth_clean_architecture/src/infra/database"
	load_user_by_email_repository "github.com/Victor-Fiamoncini/auth_clean_architecture/src/infra/repositories/load_user_by_email_repository"
	update_access_token_repository "github.com/Victor-Fiamoncini/auth_clean_architecture/src/infra/repositories/update_access_token_repository"
	"github.com/Victor-Fiamoncini/auth_clean_architecture/src/presentation/contracts"
	controllers "github.com/Victor-Fiamoncini/auth_clean_architecture/src/presentation/controllers"
	encrypter "github.com/Victor-Fiamoncini/auth_clean_architecture/src/shared/helpers/encrypter"
	token_generator "github.com/Victor-Fiamoncini/auth_clean_architecture/src/shared/helpers/token_generator"
	"github.com/Victor-Fiamoncini/auth_clean_architecture/src/shared/validators"
)

// MakeLoginController func
func MakeLoginController() contracts.IController {
	tokenSecret := os.Getenv("TOKEN_SECRET")
	userModel := database.OpenConnection()

	loadUserByEmailRepository := load_user_by_email_repository.NewLoadUserByEmailRepository(userModel)
	encrypter := encrypter.NewEncrypter()
	tokenGenerator := token_generator.NewTokenGenerator(tokenSecret)
	updateAccessTokenRepository := update_access_token_repository.NewUpdateAccessTokenRepository(userModel)

	authUseCase := auth_usecase.NewAuthUseCase(
		loadUserByEmailRepository,
		encrypter,
		tokenGenerator,
		updateAccessTokenRepository,
	)

	emailValidator := validators.NewEmailValidator()

	return controllers.NewLoginController(authUseCase, emailValidator)
}
