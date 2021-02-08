package routers_test

import (
	"testing"

	"github.com/Victor-Fiamoncini/auth_clean_architecture/src/presentation/helpers"
	"github.com/Victor-Fiamoncini/auth_clean_architecture/src/presentation/mocks"
	"github.com/Victor-Fiamoncini/auth_clean_architecture/src/presentation/routers"
	"github.com/Victor-Fiamoncini/auth_clean_architecture/src/presentation/usecases"
	"github.com/stretchr/testify/assert"
)

func makeSut() (*routers.LoginRouter, *usecases.AuthUseCase) {
	// Convert a spy class to production class
	authUseCaseSpy := (*usecases.AuthUseCase)(mocks.NewAuthUseCaseSpy())

	return routers.NewLoginRouter(authUseCaseSpy), authUseCaseSpy
}

func TestShouldReturn400IfNoEmailIsProvided(t *testing.T) {
	sut, _ := makeSut()

	httpRequest := &helpers.HTTPRequest{
		Body: struct {
			Email    string
			Password string
		}{
			Password: "any_password",
		},
	}

	httpResponse := sut.Route(httpRequest)

	assert.Equal(t, 400, httpResponse.StatusCode)
	assert.Equal(t, "Missing param: email", httpResponse.ErrorObject.Error())
}

func TestShouldReturn400IfNoPasswordIsProvided(t *testing.T) {
	sut, _ := makeSut()

	httpRequest := &helpers.HTTPRequest{
		Body: struct {
			Email    string
			Password string
		}{
			Email: "any_email@mail.com",
		},
	}

	httpResponse := sut.Route(httpRequest)

	assert.Equal(t, 400, httpResponse.StatusCode)
	assert.Equal(t, "Missing param: password", httpResponse.ErrorObject.Error())

}

func TestShouldReturn500IfNoHTTPRequestIsProvided(t *testing.T) {
	sut, _ := makeSut()

	httpResponse := sut.Route(nil)

	assert.Equal(t, 500, httpResponse.StatusCode)
}

func TestShouldReturn500IfHTTPRequestHasNoBody(t *testing.T) {
	sut, _ := makeSut()

	httpRequest := &helpers.HTTPRequest{}

	httpResponse := sut.Route(httpRequest)

	assert.Equal(t, 500, httpResponse.StatusCode)
}

func TestShouldCallAuthUseCaseWithCorrectParams(t *testing.T) {
	sut, authUseCaseSpy := makeSut()

	httpRequest := &helpers.HTTPRequest{
		Body: struct {
			Email    string
			Password string
		}{
			Email:    "any_email@mail.com",
			Password: "my_pass",
		},
	}

	sut.Route(httpRequest)

	assert.Equal(t, httpRequest.Body.Email, authUseCaseSpy.Email)
	assert.Equal(t, httpRequest.Body.Password, authUseCaseSpy.Password)
}
