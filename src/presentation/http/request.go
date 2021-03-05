package http

import (
	"github.com/Victor-Fiamoncini/auth_clean_architecture/src/presentation/contracts"
)

// Request struct
type Request struct {
	Body interface{}
}

// NewRequest func
func NewRequest() contracts.IRequest {
	return &Request{}
}

// GetBody Request method
func (r *Request) GetBody() interface{} {
	return r.Body
}

// SetBody Request method
func (r *Request) SetBody(Body interface{}) {
	r.Body = Body
}
