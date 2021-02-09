package routers

import (
	"reflect"
	"regexp"

	"github.com/Victor-Fiamoncini/auth_clean_architecture/src/presentation/helpers"
	"github.com/Victor-Fiamoncini/auth_clean_architecture/src/presentation/usecases"
)

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// ILoginRouter interface
type ILoginRouter interface {
	Route(hr *helpers.HTTPRequest) *helpers.HTTPResponse
}

// LoginRouter struct
type LoginRouter struct {
	AuthUseCase *usecases.AuthUseCase
}

// NewLoginRouter func
func NewLoginRouter(authUseCase *usecases.AuthUseCase) *LoginRouter {
	return &LoginRouter{
		AuthUseCase: authUseCase,
	}
}

// Route LoginRouter method
func (lr *LoginRouter) Route(hr *helpers.HTTPRequest) *helpers.HTTPResponse {
	httpResponse := helpers.NewHTTPResponse()

	if hr == nil || reflect.ValueOf(hr.Body).IsZero() || lr.AuthUseCase == nil {
		return httpResponse.ServerError()
	}

	email := hr.Body.Email
	password := hr.Body.Password

	if email == "" {
		return httpResponse.BadRequest("email")
	}

	if !emailRegex.MatchString(email) {
		return httpResponse.BadRequest("email")
	}

	if password == "" {
		return httpResponse.BadRequest("password")
	}

	accessToken := lr.AuthUseCase.Auth(email, password)

	if accessToken == "" {
		return httpResponse.Unauthorized()
	}

	responseBody := make(map[string]interface{})
	responseBody["AccessToken"] = accessToken

	return httpResponse.Success(responseBody)
}
