package encrypter

// IEncrypter interface
type IEncrypter interface {
	GetPassword() string
	SetPassword(password string)
	GetHashedPassword() string
	SetHashedPassword(hashedPassword string)
}
