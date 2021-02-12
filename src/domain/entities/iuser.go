package entities

// IUser interface
type IUser interface {
	GetPassword() string
	SetPassword(password string)
}
