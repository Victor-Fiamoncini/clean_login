package http

import "github.com/Victor-Fiamoncini/auth_clean_architecture/src/presentation/contracts"

// HTTPRequest struct
type HTTPRequest struct {
	Body interface{}
}

// NewHTTPRequest func
func NewHTTPRequest() contracts.IHTTPRequest {
	return &HTTPRequest{}
}

// GetBody HTTPRequest method
func (hr *HTTPRequest) GetBody() interface{} {
	return hr.Body
}

// SetBody HTTPRequest method
func (hr *HTTPRequest) SetBody(Body interface{}) {
	hr.Body = Body
}
