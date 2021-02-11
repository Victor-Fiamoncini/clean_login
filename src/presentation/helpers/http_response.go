package helpers

import (
	global_custom_errors "github.com/Victor-Fiamoncini/auth_clean_architecture/src/errors"
	custom_errors "github.com/Victor-Fiamoncini/auth_clean_architecture/src/presentation/errors"
	"github.com/Victor-Fiamoncini/auth_clean_architecture/src/types"
)

// HTTPResponse struct
type HTTPResponse struct {
	StatusCode  int
	ErrorObject error
	ErrorName   string
	Body        types.Map
}

// NewHTTPResponse func
func NewHTTPResponse() IHTTPResponse {
	return &HTTPResponse{}
}

// GetStatusCode HTTPResponse method
func (hr *HTTPResponse) GetStatusCode() int {
	return hr.StatusCode
}

// SetStatusCode HTTPResponse method
func (hr *HTTPResponse) SetStatusCode(statusCode int) {
	hr.StatusCode = statusCode
}

// GetErrorObject HTTPResponse method
func (hr *HTTPResponse) GetErrorObject() error {
	return hr.ErrorObject
}

// SetErrorObject HTTPResponse method
func (hr *HTTPResponse) SetErrorObject(errorObject error) {
	hr.ErrorObject = errorObject
}

// GetErrorName HTTPResponse method
func (hr *HTTPResponse) GetErrorName() string {
	return hr.ErrorName
}

// SetErrorName HTTPResponse method
func (hr *HTTPResponse) SetErrorName(errorName string) {
	hr.ErrorName = errorName
}

// GetBody HTTPResponse method
func (hr *HTTPResponse) GetBody() types.Map {
	return hr.Body
}

// SetBody HTTPResponse method
func (hr *HTTPResponse) SetBody(body types.Map) {
	hr.Body = body
}

// Success HTTPResponse method
func (hr *HTTPResponse) Success(data types.Map) *HTTPResponse {
	hr.StatusCode = 200
	hr.Body = data

	return hr
}

// BadRequest HTTPResponse method
func (hr *HTTPResponse) BadRequest(err global_custom_errors.IDefaultError) *HTTPResponse {
	hr.StatusCode = 400

	hr.ErrorObject = err.GetError()
	hr.ErrorName = err.GetName()

	return hr
}

// Unauthorized HTTPResponse method
func (hr *HTTPResponse) Unauthorized() *HTTPResponse {
	hr.StatusCode = 401

	newUnauthorizedError := custom_errors.NewUnauthorizedError()

	hr.ErrorObject = newUnauthorizedError.GetError()
	hr.ErrorName = newUnauthorizedError.GetName()

	return hr
}

// ServerError HTTPResponse method
func (hr *HTTPResponse) ServerError() *HTTPResponse {
	hr.StatusCode = 500

	newServerError := custom_errors.NewServerError()

	hr.ErrorObject = newServerError.GetError()
	hr.ErrorName = newServerError.GetName()

	return hr
}
