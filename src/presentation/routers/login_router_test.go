package routers_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type HttpRequest struct {
	Body struct {
		Email    string
		Password string
	}
}

type HttpResponse struct {
	StatusCode  int
	ErrorObject error
	ErrorName   string
}

func NewHttpResponse() *HttpResponse {
	return &HttpResponse{}
}

func (hres *HttpResponse) Success() *HttpResponse {
	hres.StatusCode = 200

	return hres
}

func (hres *HttpResponse) BadRequest(param string) *HttpResponse {
	hres.StatusCode = 400

	newMissingParamError := NewMissingParamError(param)

	hres.ErrorObject = newMissingParamError.Error
	hres.ErrorName = newMissingParamError.Name

	return hres
}

func (hres *HttpResponse) ServerError() *HttpResponse {
	hres.StatusCode = 500

	return hres
}

type MissingParamError struct {
	Error error
	Name  string
}

func NewMissingParamError(param string) *MissingParamError {
	return &MissingParamError{
		Error: errors.New("Missing param: " + param),
		Name:  "MissingParamError",
	}
}

type LoginRouter struct{}

func NewLoginRouter() *LoginRouter {
	return &LoginRouter{}
}

func (lr *LoginRouter) Route(hr *HttpRequest) *HttpResponse {
	if hr == nil || reflect.ValueOf(hr.Body).IsZero() {
		return NewHttpResponse().ServerError()
	}

	if hr.Body.Email == "" {
		return NewHttpResponse().BadRequest("email")
	}

	if hr.Body.Password == "" {
		return NewHttpResponse().BadRequest("password")

	}

	return NewHttpResponse().Success()

}

func TestShouldReturn400IfNoEmailIsProvided(t *testing.T) {
	sut := NewLoginRouter()

	httpRequest := &HttpRequest{
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
	sut := NewLoginRouter()

	httpRequest := &HttpRequest{
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
	sut := NewLoginRouter()

	httpResponse := sut.Route(nil)

	assert.Equal(t, 500, httpResponse.StatusCode)
}

func TestShouldReturn500IfHTTPRequestHasNoBody(t *testing.T) {
	sut := NewLoginRouter()

	httpRequest := &HttpRequest{}

	httpResponse := sut.Route(httpRequest)

	assert.Equal(t, 500, httpResponse.StatusCode)
}
