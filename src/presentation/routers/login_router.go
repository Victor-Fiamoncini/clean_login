package routers

import (
	"reflect"

	"github.com/Victor-Fiamoncini/auth_clean_architecture/src/presentation/helpers"
	"github.com/Victor-Fiamoncini/auth_clean_architecture/src/presentation/usecases"
)

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

	if hr == nil || reflect.ValueOf(hr.Body).IsZero() {
		return httpResponse.ServerError()
	}

	email := hr.Body.Email
	password := hr.Body.Password

	if email == "" {
		return httpResponse.BadRequest("email")
	}

	if password == "" {
		return httpResponse.BadRequest("password")
	}

	lr.AuthUseCase.Auth(email, password)

	return httpResponse.Unauthorized()
}
