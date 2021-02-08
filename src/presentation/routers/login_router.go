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
	if hr == nil || reflect.ValueOf(hr.Body).IsZero() {
		return helpers.NewHTTPResponse().ServerError()
	}

	email := hr.Body.Email
	password := hr.Body.Password

	if email == "" {
		return helpers.NewHTTPResponse().BadRequest("email")
	}

	if password == "" {
		return helpers.NewHTTPResponse().BadRequest("password")
	}

	lr.AuthUseCase.Auth(email, password)

	return helpers.NewHTTPResponse().Success()
}
