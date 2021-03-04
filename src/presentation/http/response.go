package http

import (
	"github.com/Victor-Fiamoncini/auth_clean_architecture/src/presentation/contracts"
	custom_errors "github.com/Victor-Fiamoncini/auth_clean_architecture/src/presentation/errors"
	shared_custom_errors "github.com/Victor-Fiamoncini/auth_clean_architecture/src/shared/errors"
	"github.com/Victor-Fiamoncini/auth_clean_architecture/src/shared/types"
)

// Response struct
type Response struct {
	StatusCode  int
	ErrorObject error
	ErrorName   string
	Body        types.Map
}

// NewResponse func
func NewResponse() contracts.IResponse {
	return &Response{}
}

// GetStatusCode Response method
func (r *Response) GetStatusCode() int {
	return r.StatusCode
}

// SetStatusCode Response method
func (r *Response) SetStatusCode(statusCode int) {
	r.StatusCode = statusCode
}

// GetErrorObject Response method
func (r *Response) GetErrorObject() error {
	return r.ErrorObject
}

// SetErrorObject Response method
func (r *Response) SetErrorObject(errorObject error) {
	r.ErrorObject = errorObject
}

// GetErrorName Response method
func (r *Response) GetErrorName() string {
	return r.ErrorName
}

// SetErrorName Response method
func (r *Response) SetErrorName(errorName string) {
	r.ErrorName = errorName
}

// GetBody Response method
func (r *Response) GetBody() types.Map {
	return r.Body
}

// SetBody Response method
func (r *Response) SetBody(body types.Map) {
	r.Body = body
}

// Success Response method
func (r *Response) Success(data types.Map) contracts.IResponse {
	r.StatusCode = 200
	r.Body = data

	return r
}

// BadRequest Response method
func (r *Response) BadRequest(err shared_custom_errors.IDefaultError) contracts.IResponse {
	r.StatusCode = 400

	r.ErrorObject = err.GetError()
	r.ErrorName = err.GetName()

	return r
}

// Unauthorized Response method
func (r *Response) Unauthorized() contracts.IResponse {
	r.StatusCode = 401

	newUnauthorizedError := custom_errors.NewUnauthorizedError()

	r.ErrorObject = newUnauthorizedError.GetError()
	r.ErrorName = newUnauthorizedError.GetName()

	return r
}

// ServerError Response method
func (r *Response) ServerError() contracts.IResponse {
	r.StatusCode = 500

	newServerError := custom_errors.NewServerError()

	r.ErrorObject = newServerError.GetError()
	r.ErrorName = newServerError.GetName()

	return r
}
