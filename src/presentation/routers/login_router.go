package routers

import (
	"reflect"

	"github.com/Victor-Fiamoncini/auth_clean_architecture/src/presentation/helpers"
)

// LoginRouter struct
type LoginRouter struct{}

// NewLoginRouter func
func NewLoginRouter() *LoginRouter {
	return &LoginRouter{}
}

// Route LoginRouter method
func (lr *LoginRouter) Route(hr *helpers.HTTPRequest) *helpers.HTTPResponse {
	if hr == nil || reflect.ValueOf(hr.Body).IsZero() {
		return helpers.NewHTTPResponse().ServerError()
	}

	if hr.Body.Email == "" {
		return helpers.NewHTTPResponse().BadRequest("email")
	}

	if hr.Body.Password == "" {
		return helpers.NewHTTPResponse().BadRequest("password")

	}

	return helpers.NewHTTPResponse().Success()

}
