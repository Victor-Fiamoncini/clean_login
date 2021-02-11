package authusecase_test

import (
	"testing"

	usecases "github.com/Victor-Fiamoncini/auth_clean_architecture/src/domain/usecases/auth_usecase"
	luber "github.com/Victor-Fiamoncini/auth_clean_architecture/src/infra/repositories/load_user_by_email_repository"
	"github.com/stretchr/testify/assert"
)

func makeSut() (usecases.IAuthUseCase, luber.ILoadUserByEmailRepository) {
	loadUserByEmailRepositorySpy := luber.NewLoadUserByEmailRepositorySpy()

	return usecases.NewAuthUseCase(loadUserByEmailRepositorySpy), loadUserByEmailRepositorySpy
}

func TestShouldCallLoadUserByEmailRepositoryWithCorrectEmail(t *testing.T) {
	sut, loadUserByEmailRepositorySpy := makeSut()

	sut.Auth("any_email@mail.com", "any_password")

	assert.Equal(t, "any_email@mail.com", loadUserByEmailRepositorySpy.GetEmail())
}
