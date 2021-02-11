package loaduserbyemailrepository

// LoadUserByEmailRepositorySpy struct
type LoadUserByEmailRepositorySpy struct {
	Email string
}

// NewLoadUserByEmailRepositorySpy func
func NewLoadUserByEmailRepositorySpy() ILoadUserByEmailRepository {
	return &LoadUserByEmailRepositorySpy{}
}

// GetEmail LoadUserByEmailRepositorySpy method
func (lubers *LoadUserByEmailRepositorySpy) GetEmail() string {
	return lubers.Email
}

// SetEmail LoadUserByEmailRepositorySpy method
func (lubers *LoadUserByEmailRepositorySpy) SetEmail(email string) {
	lubers.Email = email
}

// Load LoadUserByEmailRepositorySpy method
func (lubers *LoadUserByEmailRepositorySpy) Load(email string) {
	lubers.Email = email
}
