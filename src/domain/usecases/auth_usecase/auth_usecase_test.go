package authusecase_test

import (
	"testing"

	usecases "github.com/Victor-Fiamoncini/auth_clean_architecture/src/domain/usecases/auth_usecase"
	encrypter "github.com/Victor-Fiamoncini/auth_clean_architecture/src/infra/libs/encrypter"
	encrypter_mocks "github.com/Victor-Fiamoncini/auth_clean_architecture/src/infra/libs/encrypter/mocks"
	luber "github.com/Victor-Fiamoncini/auth_clean_architecture/src/infra/repositories/load_user_by_email_repository"
	luber_mocks "github.com/Victor-Fiamoncini/auth_clean_architecture/src/infra/repositories/load_user_by_email_repository/mocks"
	"github.com/stretchr/testify/assert"
)

func makeSut() (usecases.IAuthUseCase, luber.ILoadUserByEmailRepository, encrypter.IEncrypter) {
	loadUserByEmailRepositorySpy := luber_mocks.NewLoadUserByEmailRepositorySpy()
	encrypterSpy := encrypter_mocks.NewEncrypterSpy()

	return usecases.NewAuthUseCase(loadUserByEmailRepositorySpy), loadUserByEmailRepositorySpy, encrypterSpy
}

func TestShouldCallLoadUserByEmailRepositoryWithCorrectEmail(t *testing.T) {
	sut, loadUserByEmailRepositorySpy, _ := makeSut()

	sut.Auth("any_email@mail.com", "any_password")

	assert.Equal(t, "any_email@mail.com", loadUserByEmailRepositorySpy.GetEmail())
}

func TestShouldCallEncrypterWithCorrectValues(t *testing.T) {
	sut, loadUserByEmailRepositorySpy, encrypterSpy := makeSut()

	sut.Auth("valid_email@mail.com", "any_password")

	assert.Equal(t, "any_password", encrypterSpy.GetPassword())
	assert.Equal(t, loadUserByEmailRepositorySpy.GetUser().GetPassword(), encrypterSpy.GetHashedPassword())
}
