package routers_test

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type httpRequest struct {
	Body struct {
		Email    string
		Password string
	}
}

type httpResponse struct {
	statusCode int
}

type LoginRouter struct{}

func NewLoginRouter() *LoginRouter {
	return &LoginRouter{}
}

func (lr *LoginRouter) Route(hr *httpRequest) *httpResponse {
	if hr == nil || reflect.ValueOf(hr.Body).IsZero() {
		return &httpResponse{
			statusCode: 500,
		}
	}

	if hr.Body.Email == "" || hr.Body.Password == "" {
		return &httpResponse{
			statusCode: 400,
		}
	}

	return &httpResponse{
		statusCode: 200,
	}
}

func TestShouldReturn400IfNoEmailIsProvided(t *testing.T) {
	sut := NewLoginRouter()

	httpRequest := &httpRequest{
		Body: struct {
			Email    string
			Password string
		}{
			Password: "any_password",
		},
	}

	httpResponse := sut.Route(httpRequest)

	assert.Equal(t, 400, httpResponse.statusCode)
}

func TestShouldReturn400IfNoPasswordIsProvided(t *testing.T) {
	sut := NewLoginRouter()

	httpRequest := &httpRequest{
		Body: struct {
			Email    string
			Password string
		}{
			Email: "any_email@mail.com",
		},
	}

	httpResponse := sut.Route(httpRequest)

	assert.Equal(t, 400, httpResponse.statusCode)
}

func TestShouldReturn500IfNoHTTPRequestIsProvided(t *testing.T) {
	sut := NewLoginRouter()

	httpResponse := sut.Route(nil)

	assert.Equal(t, 500, httpResponse.statusCode)
}

func TestShouldReturn500IfHTTPRequestHasNoBody(t *testing.T) {
	sut := NewLoginRouter()

	httpRequest := &httpRequest{}

	httpResponse := sut.Route(httpRequest)

	assert.Equal(t, 500, httpResponse.statusCode)
}
