package routers

import (
	"reflect"

	custom_errors "github.com/Victor-Fiamoncini/auth_clean_architecture/src/presentation/errors"
	"github.com/Victor-Fiamoncini/auth_clean_architecture/src/presentation/helpers"
	"github.com/Victor-Fiamoncini/auth_clean_architecture/src/presentation/usecases"
	"github.com/Victor-Fiamoncini/auth_clean_architecture/src/types"
)

// var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// ILoginRouter interface
type ILoginRouter interface {
	Route(httpRequest *helpers.HTTPRequest) helpers.IHTTPResponse
}

// LoginRouter struct
type LoginRouter struct {
	AuthUseCase usecases.IAuthUseCase
}

// NewLoginRouter func
func NewLoginRouter(authUseCase usecases.IAuthUseCase) ILoginRouter {
	return &LoginRouter{
		AuthUseCase: authUseCase,
	}
}

// Route LoginRouter method
func (lr *LoginRouter) Route(httpRequest *helpers.HTTPRequest) helpers.IHTTPResponse {
	httpResponse := helpers.NewHTTPResponse()

	if httpRequest == nil || reflect.ValueOf(httpRequest.Body).IsZero() || lr.AuthUseCase == nil {
		return httpResponse.ServerError()
	}

	email := httpRequest.Body.Email
	password := httpRequest.Body.Password

	if email == "" {
		return httpResponse.BadRequest(custom_errors.NewMissingParamError("email"))
	}

	// if !emailRegex.MatchString(email) {
	// 	return httpResponse.BadRequest("email")
	// }

	if password == "" {
		return httpResponse.BadRequest(custom_errors.NewMissingParamError("password"))
	}

	accessToken := lr.AuthUseCase.Auth(email, password)

	if accessToken == "" {
		return httpResponse.Unauthorized()
	}

	responseBody := make(types.Map)
	responseBody["AccessToken"] = accessToken

	return httpResponse.Success(responseBody)
}
