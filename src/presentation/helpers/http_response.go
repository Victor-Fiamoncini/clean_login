package helpers

import (
	custom_errors "github.com/Victor-Fiamoncini/auth_clean_architecture/src/presentation/errors"
)

// IHTTPResponse interface
type IHTTPResponse interface {
	Success() *HTTPResponse
	BadRequest(param string) *HTTPResponse
	Unauthorized() *HTTPResponse
	ServerError() *HTTPResponse
}

// HTTPResponse struct
type HTTPResponse struct {
	StatusCode  int
	ErrorObject error
	ErrorName   string
	Body        map[string]interface{}
}

// NewHTTPResponse func
func NewHTTPResponse() *HTTPResponse {
	return &HTTPResponse{}
}

// Success HTTPResponse method
func (hres *HTTPResponse) Success(data map[string]interface{}) *HTTPResponse {
	hres.StatusCode = 200
	hres.Body = data

	return hres
}

// BadRequest HTTPResponse method
func (hres *HTTPResponse) BadRequest(err custom_errors.IDefaultError) *HTTPResponse {
	hres.StatusCode = 400

	newMissingParamError := err(param)

	hres.ErrorObject = error.Error()
	hres.ErrorName = errors

	return hres
}

// Unauthorized HTTPResponse method
func (hres *HTTPResponse) Unauthorized() *HTTPResponse {
	hres.StatusCode = 401

	newUnauthorizedError := custom_errors.NewUnauthorizedError()

	hres.ErrorObject = newUnauthorizedError.Error
	hres.ErrorName = newUnauthorizedError.Name

	return hres
}

// ServerError HTTPResponse method
func (hres *HTTPResponse) ServerError() *HTTPResponse {
	hres.StatusCode = 500

	newServerError := custom_errors.NewServerError()

	hres.ErrorObject = newServerError.Error
	hres.ErrorName = newServerError.Name

	return hres
}
