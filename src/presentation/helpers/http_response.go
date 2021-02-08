package helpers

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

// ServerError HTTPResponse method
func (hres *HTTPResponse) ServerError() *HTTPResponse {
	hres.StatusCode = 500

	return hres
}
