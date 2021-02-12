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
	"github.com/stretchr/testify/assert"
)

func makeSut() (usecases.IAuthUseCase, luber.ILoadUserByEmailRepository, encrypter.IEncrypter, token_generator.ITokenGenerator) {
	loadUserByEmailRepositorySpy := luber_mocks.NewLoadUserByEmailRepositorySpy()
	encrypterSpy := encrypter_mocks.NewEncrypterSpy()
	tokenGeneratorSpy := token_generator_mocks.NewTokenGeneratorSpy()

	encrypterSpy.SetIsValid(true)
	tokenGeneratorSpy.SetAccessToken("any_token")

	user := entities.NewUser()
	user.SetID("any_id")
	user.SetPassword("hashed_password")

	loadUserByEmailRepositorySpy.SetUser(user)

	return usecases.NewAuthUseCase(loadUserByEmailRepositorySpy, encrypterSpy, tokenGeneratorSpy), loadUserByEmailRepositorySpy, encrypterSpy, tokenGeneratorSpy
}

func TestShouldCallLoadUserByEmailRepositoryWithCorrectEmail(t *testing.T) {
	sut, loadUserByEmailRepositorySpy, _, _ := makeSut()

	sut.Auth("any_email@mail.com", "any_password")

	assert.Equal(t, "any_email@mail.com", loadUserByEmailRepositorySpy.GetEmail())
}

func TestShouldCallEncrypterWithCorrectValues(t *testing.T) {
	sut, loadUserByEmailRepositorySpy, encrypterSpy, _ := makeSut()

	sut.Auth("valid_email@mail.com", "any_password")

	assert.Equal(t, "any_password", encrypterSpy.GetPassword())
	assert.Equal(t, loadUserByEmailRepositorySpy.GetUser().GetPassword(), encrypterSpy.GetHashedPassword())
}

func TestShouldReturnAnEmptyTokenIfAnInvalidPasswordIsProvided(t *testing.T) {
	sut, _, encrypterSpy, _ := makeSut()

	encrypterSpy.SetIsValid(false)

	accessToken := sut.Auth("valid_email@mail.com", "invalid_password")

	assert.Empty(t, "", accessToken)
}

func TestShouldCallTokenGeneratorWithCorrectUserID(t *testing.T) {
	sut, loadUserByEmailRepositorySpy, _, tokenGeneratorSpy := makeSut()

	sut.Auth("valid_email@mail.com", "valid_password")

	assert.Equal(t, loadUserByEmailRepositorySpy.GetUser().GetID(), tokenGeneratorSpy.GetUserID())
}
