package http

import (
	"github.com/Victor-Fiamoncini/auth_clean_architecture/src/presentation/contracts"
	"github.com/Victor-Fiamoncini/auth_clean_architecture/src/shared/types"
)

// Request struct
type Request struct {
	Body types.StringMap
}

// NewRequest func
func NewRequest() contracts.IRequest {
	return &Request{}
}

// GetBody Request method
func (r *Request) GetBody() types.StringMap {
	return r.Body
}

// SetBody Request method
func (r *Request) SetBody(Body types.StringMap) {
	r.Body = Body
}
