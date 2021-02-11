package routers

import (
	"reflect"

	"github.com/Victor-Fiamoncini/auth_clean_architecture/src/domain/usecases"
	global_custom_errors "github.com/Victor-Fiamoncini/auth_clean_architecture/src/errors"
	"github.com/Victor-Fiamoncini/auth_clean_architecture/src/presentation/helpers"
	"github.com/Victor-Fiamoncini/auth_clean_architecture/src/types"
	"github.com/Victor-Fiamoncini/auth_clean_architecture/src/validators"
)

// ILoginRouter interface
type ILoginRouter interface {
	Route(httpRequest *helpers.HTTPRequest) helpers.IHTTPResponse
}

// LoginRouter struct
type LoginRouter struct {
	AuthUseCase    usecases.IAuthUseCase
	EmailValidator validators.IEmailValidator
}

// NewLoginRouter func
func NewLoginRouter(authUseCase usecases.IAuthUseCase, emailValidator validators.IEmailValidator) ILoginRouter {
	return &LoginRouter{
		AuthUseCase:    authUseCase,
		EmailValidator: emailValidator,
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
		return httpResponse.BadRequest(global_custom_errors.NewMissingParamError("email"))
	}

	if !lr.EmailValidator.Run(email) {
		return httpResponse.BadRequest(global_custom_errors.NewInvalidParamError("email"))
	}

	if password == "" {
		return httpResponse.BadRequest(global_custom_errors.NewMissingParamError("password"))
	}

	accessToken := lr.AuthUseCase.Auth(email, password)

	if accessToken == "" {
		return httpResponse.Unauthorized()
	}

	responseBody := make(types.Map)
	responseBody["AccessToken"] = accessToken

	return httpResponse.Success(responseBody)
}
