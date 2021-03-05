package controller_test

import (
	"testing"

	auth_usecase "github.com/Victor-Fiamoncini/auth_clean_architecture/src/domain/usecases/auth_usecase"
	"github.com/Victor-Fiamoncini/auth_clean_architecture/src/presentation/contracts"
	"github.com/Victor-Fiamoncini/auth_clean_architecture/src/presentation/contracts/payloads"
	controller "github.com/Victor-Fiamoncini/auth_clean_architecture/src/presentation/controllers"
	"github.com/Victor-Fiamoncini/auth_clean_architecture/src/presentation/http"
	"github.com/Victor-Fiamoncini/auth_clean_architecture/src/shared/validators"
	auth_usecase_mocks "github.com/Victor-Fiamoncini/auth_clean_architecture/src/test/domain/usecases/auth_usecase/mocks"
	validators_mocks "github.com/Victor-Fiamoncini/auth_clean_architecture/src/test/shared/validators/mocks"
	"github.com/stretchr/testify/assert"
)

func makeSut() (contracts.IController, auth_usecase.IAuthUseCase, validators.IEmailValidator) {
	authUseCaseSpy := auth_usecase_mocks.NewAuthUseCaseSpy()
	emailValidatorSpy := validators_mocks.NewEmailValidatorSpy()

	authUseCaseSpy.SetAccessToken("valid_token")
	emailValidatorSpy.SetIsEmailValid(true)

	return controller.NewLoginController(authUseCaseSpy, emailValidatorSpy), authUseCaseSpy, emailValidatorSpy
}

func TestShouldReturn400IfNoEmailIsProvided(t *testing.T) {
	sut, _, _ := makeSut()

	loginPayload := payloads.NewLoginPayload()
	loginPayload.Password = "any_password"

	httpRequest := http.NewRequest()
	httpRequest.SetBody(loginPayload)

	httpResponse := sut.Handle(httpRequest)

	assert.Equal(t, 400, httpResponse.GetStatusCode())
	assert.Equal(t, "Missing param: email", httpResponse.GetErrorObject().Error())
}

func TestShouldReturn400IfNoPasswordIsProvided(t *testing.T) {
	sut, _, _ := makeSut()

	loginPayload := payloads.NewLoginPayload()
	loginPayload.Email = "any_email@mail.com"

	httpRequest := http.NewRequest()
	httpRequest.SetBody(loginPayload)

	httpResponse := sut.Handle(httpRequest)

	assert.Equal(t, 400, httpResponse.GetStatusCode())
	assert.Equal(t, "Missing param: password", httpResponse.GetErrorObject().Error())
}

func TestShouldReturn500IfNoHTTPRequestIsProvided(t *testing.T) {
	sut, _, _ := makeSut()

	httpResponse := sut.Handle(nil)

	assert.Equal(t, 500, httpResponse.GetStatusCode())
	assert.Equal(t, "Internal Error", httpResponse.GetErrorObject().Error())
}

func TestShouldReturn500IfHTTPRequestHasNoBody(t *testing.T) {
	sut, _, _ := makeSut()

	httpRequest := http.NewRequest()

	httpResponse := sut.Handle(httpRequest)

	assert.Equal(t, 500, httpResponse.GetStatusCode())
	assert.Equal(t, "Internal Error", httpResponse.GetErrorObject().Error())
}

func TestShouldCallAuthUseCaseWithCorrectParams(t *testing.T) {
	sut, authUseCaseSpy, _ := makeSut()

	loginPayload := payloads.NewLoginPayload()
	loginPayload.Email = "any_email@mail.com"
	loginPayload.Password = "any_password"

	httpRequest := http.NewRequest()
	httpRequest.SetBody(loginPayload)

	sut.Handle(httpRequest)

	assert.Equal(t, "any_email@mail.com", authUseCaseSpy.GetEmail())
	assert.Equal(t, "any_password", authUseCaseSpy.GetPassword())
}

func TestShouldReturn401WhenInvalidCredentialsAreProvided(t *testing.T) {
	sut, authUseCaseSpy, _ := makeSut()

	authUseCaseSpy.SetAccessToken("")

	loginPayload := payloads.NewLoginPayload()
	loginPayload.Email = "invalid_email@mail.com"
	loginPayload.Password = "invalid_password"

	httpRequest := http.NewRequest()
	httpRequest.SetBody(loginPayload)

	httpResponse := sut.Handle(httpRequest)

	assert.Equal(t, 401, httpResponse.GetStatusCode())
	assert.Equal(t, "Unauthorized", httpResponse.GetErrorObject().Error())
}

func TestShouldReturn200WhenValidCredentailsAreProvided(t *testing.T) {
	sut, authUseCaseSpy, _ := makeSut()

	loginPayload := payloads.NewLoginPayload()
	loginPayload.Email = "valid_email@mail.com"
	loginPayload.Password = "valid_password"

	httpRequest := http.NewRequest()
	httpRequest.SetBody(loginPayload)

	httpResponse := sut.Handle(httpRequest)

	assert.Equal(t, 200, httpResponse.GetStatusCode())
	assert.Equal(t, httpResponse.GetBody()["AccessToken"], authUseCaseSpy.GetAccessToken())
}

func TestShouldReturn400IfAnInvalidEmailIsProvided(t *testing.T) {
	sut, _, emailValidatorSpy := makeSut()

	emailValidatorSpy.SetIsEmailValid(false)

	loginPayload := payloads.NewLoginPayload()
	loginPayload.Email = "invalid_email@mail.com"
	loginPayload.Password = "any_password"

	httpRequest := http.NewRequest()
	httpRequest.SetBody(loginPayload)

	httpResponse := sut.Handle(httpRequest)

	assert.Equal(t, 400, httpResponse.GetStatusCode())
	assert.Equal(t, "Invalid param: email", httpResponse.GetErrorObject().Error())
}
