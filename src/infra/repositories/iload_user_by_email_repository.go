package repositories

// ILoadUserByEmailRepository interface
type ILoadUserByEmailRepository interface {
	Load(email string)
}
