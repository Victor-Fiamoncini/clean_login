package helpers

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
}

// NewHTTPResponse func
func NewHTTPResponse() *HTTPResponse {
	return &HTTPResponse{}
}

// Success HTTPResponse method
func (hres *HTTPResponse) Success() *HTTPResponse {
	hres.StatusCode = 200

	return hres
}

// BadRequest HTTPResponse method
func (hres *HTTPResponse) BadRequest(param string) *HTTPResponse {
	hres.StatusCode = 400

	newMissingParamError := NewMissingParamError(param)

	hres.ErrorObject = newMissingParamError.Error
	hres.ErrorName = newMissingParamError.Name

	return hres
}

// Unauthorized HTTPResponse method
func (hres *HTTPResponse) Unauthorized() *HTTPResponse {
	hres.StatusCode = 401

	newUnauthorizedError := NewUnauthorizedError()

	hres.ErrorObject = newUnauthorizedError.Error
	hres.ErrorName = newUnauthorizedError.Name

	return hres
}

// ServerError HTTPResponse method
func (hres *HTTPResponse) ServerError() *HTTPResponse {
	hres.StatusCode = 500

	return hres
}
