package repositories

// LoadUserByEmailRepositorySpy struct
type LoadUserByEmailRepositorySpy struct{}

// NewLoadUserByEmailRepositorySpy func
func NewLoadUserByEmailRepositorySpy() ILoadUserByEmailRepository {
	return &LoadUserByEmailRepositorySpy{}
}

// Load LoadUserByEmailRepositorySpy method
func (lubers *LoadUserByEmailRepositorySpy) Load(email string) {

}
