package authusecase_test

import (
	"testing"

	"github.com/Victor-Fiamoncini/auth_clean_architecture/src/domain/entities"
	usecases "github.com/Victor-Fiamoncini/auth_clean_architecture/src/domain/usecases/auth_usecase"
	encrypter "github.com/Victor-Fiamoncini/auth_clean_architecture/src/infra/libs/encrypter"
	encrypter_mocks "github.com/Victor-Fiamoncini/auth_clean_architecture/src/infra/libs/encrypter/mocks"
	token_generator "github.com/Victor-Fiamoncini/auth_clean_architecture/src/infra/libs/token_generator"
	token_generator_mocks "github.com/Victor-Fiamoncini/auth_clean_architecture/src/infra/libs/token_generator/mocks"
	luber "github.com/Victor-Fiamoncini/auth_clean_architecture/src/infra/repositories/load_user_by_email_repository"
	luber_mocks "github.com/Victor-Fiamoncini/auth_clean_architecture/src/infra/repositories/load_user_by_email_repository/mocks"
	uatr "github.com/Victor-Fiamoncini/auth_clean_architecture/src/infra/repositories/update_access_token_repository"
	uatr_mocks "github.com/Victor-Fiamoncini/auth_clean_architecture/src/infra/repositories/update_access_token_repository/mocks"
	"github.com/stretchr/testify/assert"
)

func makeSut() (usecases.IAuthUseCase, luber.ILoadUserByEmailRepository, encrypter.IEncrypter, token_generator.ITokenGenerator, uatr.IUpdateAccessTokenRepository) {
	loadUserByEmailRepositorySpy := luber_mocks.NewLoadUserByEmailRepositorySpy()
	encrypterSpy := encrypter_mocks.NewEncrypterSpy()
	tokenGeneratorSpy := token_generator_mocks.NewTokenGeneratorSpy()
	updateAccessTokenRepositorySpy := uatr_mocks.NewUpdateAccessTokenRepositorySpy()

	encrypterSpy.SetIsValid(true)
	tokenGeneratorSpy.SetAccessToken("any_token")

	user := entities.NewUser()
	user.SetID("any_id")
	user.SetPassword("hashed_password")

	loadUserByEmailRepositorySpy.SetUser(user)

	authUseCase := usecases.NewAuthUseCase(loadUserByEmailRepositorySpy, encrypterSpy, tokenGeneratorSpy, updateAccessTokenRepositorySpy)

	return authUseCase, loadUserByEmailRepositorySpy, encrypterSpy, tokenGeneratorSpy, updateAccessTokenRepositorySpy
}

func TestShouldCallLoadUserByEmailRepositoryWithCorrectEmail(t *testing.T) {
	sut, loadUserByEmailRepositorySpy, _, _, _ := makeSut()

	sut.SetEmail("any_email@mail.com")
	sut.SetPassword("any_password")

	sut.Auth()

	assert.Equal(t, "any_email@mail.com", loadUserByEmailRepositorySpy.GetEmail())
}

func TestShouldCallEncrypterWithCorrectValues(t *testing.T) {
	sut, loadUserByEmailRepositorySpy, encrypterSpy, _, _ := makeSut()

	sut.SetEmail("valid_email@mail.com")
	sut.SetPassword("any_password")

	sut.Auth()

	assert.Equal(t, "any_password", encrypterSpy.GetPassword())
	assert.Equal(t, loadUserByEmailRepositorySpy.GetUser().GetPassword(), encrypterSpy.GetHashedPassword())
}

func TestShouldCallTokenGeneratorWithCorrectUserID(t *testing.T) {
	sut, loadUserByEmailRepositorySpy, _, tokenGeneratorSpy, _ := makeSut()

	sut.SetEmail("valid_email@mail.com")
	sut.SetPassword("valid_password")

	sut.Auth()

	assert.Equal(t, loadUserByEmailRepositorySpy.GetUser().GetID(), tokenGeneratorSpy.GetUserID())
}

func TestShouldReturnAnAccessTokenIfCorrectCredentialsAreProvided(t *testing.T) {
	sut, _, _, tokenGeneratorSpy, _ := makeSut()

	sut.SetEmail("valid_email@mail.com")
	sut.SetPassword("valid_password")

	accessToken := sut.Auth()

	assert.Equal(t, tokenGeneratorSpy.GetAccessToken(), accessToken)
	assert.NotNil(t, accessToken)
	assert.NotEmpty(t, accessToken)
}

// func TestShouldCallUpdateAccessTokenRepositoryWithCorrectValues(t *testing.T) {
// 	sut, loadUserByEmailRepositorySpy, _, tokenGeneratorSpy, updateAccessTokenRepositorySpy := makeSut()

// 	sut.SetEmail("valid_email@mail.com")
// 	sut.SetPassword("valid_password")

// 	accessToken := sut.Auth()

// 	assert.Equal(t, updateAccessTokenRepositorySpy.GetUserID(), loadUserByEmailRepositorySpy.GetUser().GetID())
// 	assert.Equal(t, updateAccessTokenRepositorySpy.GetAccessToken(), tokenGeneratorSpy.GetAccessToken())
// }
