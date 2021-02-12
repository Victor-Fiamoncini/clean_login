package entities

// IUser interface
type IUser interface {
	GetID() string
	SetID(id string)
	GetPassword() string
	SetPassword(password string)
}
