package helpers

// HTTPRequest struct
type HTTPRequest struct {
	Body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
}

// NewHTTPRequest func
func NewHTTPRequest() *HTTPRequest {
	return &HTTPRequest{}
}
