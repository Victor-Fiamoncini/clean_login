package payloads

// LoginPayload struct
type LoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// NewLoginPayload func
func NewLoginPayload() *LoginPayload {
	return &LoginPayload{}
}
