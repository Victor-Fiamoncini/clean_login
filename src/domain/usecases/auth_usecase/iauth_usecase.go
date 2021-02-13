package authusecase

// IAuthUseCase interface
type IAuthUseCase interface {
	GetEmail() string
	SetEmail(email string)
	GetPassword() string
	SetPassword(password string)
	GetAccessToken() string
	SetAccessToken(accessToken string)

	Auth() string
}
