package routers_test

import (
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
	if hr.Body.Email == "" {
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
			Email:    "",
			Password: "any_password",
		},
	}

	httpResponse := sut.Route(httpRequest)

	assert.Equal(t, 400, httpResponse.statusCode)
}
