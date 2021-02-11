package loaduserbyemailrepository

// ILoadUserByEmailRepository interface
type ILoadUserByEmailRepository interface {
	GetEmail() string
	SetEmail(email string)

	Load(email string)
}
