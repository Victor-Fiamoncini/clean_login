package controller_test

import (
	"testing"

	auth_usecase "github.com/Victor-Fiamoncini/auth_clean_architecture/src/domain/usecases/auth_usecase"
	auth_usecase_mocks "github.com/Victor-Fiamoncini/auth_clean_architecture/src/domain/usecases/auth_usecase/mocks"
	"github.com/Victor-Fiamoncini/auth_clean_architecture/src/presentation/contracts"
	controller "github.com/Victor-Fiamoncini/auth_clean_architecture/src/presentation/controllers"
	http "github.com/Victor-Fiamoncini/auth_clean_architecture/src/presentation/http"
	"github.com/Victor-Fiamoncini/auth_clean_architecture/src/shared/types"
	"github.com/Victor-Fiamoncini/auth_clean_architecture/src/shared/validators"
	validators_mocks "github.com/Victor-Fiamoncini/auth_clean_architecture/src/shared/validators/mocks"
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

	httpRequest := http.NewRequest()
	body := make(types.StringMap)
	body["password"] = "any_password"
	body["email"] = "hello"

	httpRequest.SetBody(body)

	httpResponse := sut.Handle(httpRequest)

	assert.Equal(t, 400, httpResponse.GetStatusCode())
	assert.Equal(t, "Missing param: email", httpResponse.GetErrorObject().Error())
}

// func TestShouldReturn400IfNoPasswordIsProvided(t *testing.T) {
// 	sut, _, _ := makeSut()

// 	httpRequest := &helpers.HTTPRequest{
// 		Body: struct {
// 			Email    string `json:"email"`
// 			Password string `json:"password"`
// 		}{
// 			Email: "any_email@mail.com",
// 		},
// 	}

// 	httpResponse := sut.Route(httpRequest)

// 	assert.Equal(t, 400, httpResponse.GetStatusCode())
// 	assert.Equal(t, "Missing param: password", httpResponse.GetErrorObject().Error())
// }

// func TestShouldReturn500IfNoHTTPRequestIsProvided(t *testing.T) {
// 	sut, _, _ := makeSut()

// 	httpResponse := sut.Route(nil)

// 	assert.Equal(t, 500, httpResponse.GetStatusCode())
// 	assert.Equal(t, "Internal Error", httpResponse.GetErrorObject().Error())
// }

// func TestShouldReturn500IfHTTPRequestHasNoBody(t *testing.T) {
// 	sut, _, _ := makeSut()

// 	httpRequest := &helpers.HTTPRequest{}

// 	httpResponse := sut.Route(httpRequest)

// 	assert.Equal(t, 500, httpResponse.GetStatusCode())
// 	assert.Equal(t, "Internal Error", httpResponse.GetErrorObject().Error())
// }

// func TestShouldCallAuthUseCaseWithCorrectParams(t *testing.T) {
// 	sut, authUseCaseSpy, _ := makeSut()

// 	httpRequest := &helpers.HTTPRequest{
// 		Body: struct {
// 			Email    string `json:"email"`
// 			Password string `json:"password"`
// 		}{
// 			Email:    "any_email@mail.com",
// 			Password: "any_password",
// 		},
// 	}

// 	sut.Route(httpRequest)

// 	assert.Equal(t, httpRequest.Body.Email, authUseCaseSpy.GetEmail())
// 	assert.Equal(t, httpRequest.Body.Password, authUseCaseSpy.GetPassword())
// }

// func TestShouldReturn401WhenInvalidCredentialsAreProvided(t *testing.T) {
// 	sut, authUseCaseSpy, _ := makeSut()

// 	authUseCaseSpy.SetAccessToken("")

// 	httpRequest := &helpers.HTTPRequest{
// 		Body: struct {
// 			Email    string `json:"email"`
// 			Password string `json:"password"`
// 		}{
// 			Email:    "invalid_email@mail.com",
// 			Password: "invalid_pass",
// 		},
// 	}

// 	httpResponse := sut.Route(httpRequest)

// 	assert.Equal(t, 401, httpResponse.GetStatusCode())
// 	assert.Equal(t, "Unauthorized", httpResponse.GetErrorObject().Error())
// }

// func TestShouldReturn500IfNoAuthUseCaseIsProvided(t *testing.T) {
// 	sut := routers.NewLoginRouter(nil, nil)

// 	httpRequest := &helpers.HTTPRequest{
// 		Body: struct {
// 			Email    string `json:"email"`
// 			Password string `json:"password"`
// 		}{
// 			Email:    "any_email@mail.com",
// 			Password: "any_pass",
// 		},
// 	}

// 	httpResponse := sut.Route(httpRequest)

// 	assert.Equal(t, 500, httpResponse.GetStatusCode())
// 	assert.Equal(t, "Internal Error", httpResponse.GetErrorObject().Error())
// }

// func TestShouldReturn200WhenValidCredentailsAreProvided(t *testing.T) {
// 	sut, authUseCaseSpy, _ := makeSut()

// 	httpRequest := &helpers.HTTPRequest{
// 		Body: struct {
// 			Email    string `json:"email"`
// 			Password string `json:"password"`
// 		}{
// 			Email:    "valid_email@mail.com",
// 			Password: "valid_pass",
// 		},
// 	}

// 	httpResponse := sut.Route(httpRequest)

// 	assert.Equal(t, 200, httpResponse.GetStatusCode())
// 	assert.Equal(t, httpResponse.GetBody()["AccessToken"], authUseCaseSpy.GetAccessToken())
// }

// func TestShouldReturn400IfAnInvalidEmailIsProvided(t *testing.T) {
// 	sut, _, emailValidatorSpy := makeSut()

// 	emailValidatorSpy.SetIsEmailValid(false)

// 	httpRequest := &helpers.HTTPRequest{
// 		Body: struct {
// 			Email    string `json:"email"`
// 			Password string `json:"password"`
// 		}{
// 			Email:    "invalid_email@mail.com",
// 			Password: "any_password",
// 		},
// 	}

// 	httpResponse := sut.Route(httpRequest)

// 	assert.Equal(t, 400, httpResponse.GetStatusCode())
// 	assert.Equal(t, "Invalid param: email", httpResponse.GetErrorObject().Error())
// }
