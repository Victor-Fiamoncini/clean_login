package usecases

// IAuthUseCase interface
type IAuthUseCase interface {
	GetEmail() string
	SetEmail(email string)
	GetPassword() string
	SetPassword(password string)
	GetAccessToken() string
	SetAccessToken(accessToken string)

	Auth(email string, password string) string
}

// AuthUseCase struct
type AuthUseCase struct {
	Email       string
	Password    string
	AccessToken string
}

// NewAuthUseCase func
func NewAuthUseCase() IAuthUseCase {
	return &AuthUseCase{}
}

// GetEmail AuthUseCase method
func (auc *AuthUseCase) GetEmail() string {
	return auc.Email
}

// SetEmail AuthUseCase method
func (auc *AuthUseCase) SetEmail(email string) {
	auc.Email = email
}

// GetPassword AuthUseCase method
func (auc *AuthUseCase) GetPassword() string {
	return auc.Password
}

// SetPassword AuthUseCase method
func (auc *AuthUseCase) SetPassword(password string) {
	auc.Password = password
}

// GetAccessToken AuthUseCase method
func (auc *AuthUseCase) GetAccessToken() string {
	return auc.AccessToken
}

// SetAccessToken AuthUseCase method
func (auc *AuthUseCase) SetAccessToken(accessToken string) {
	auc.AccessToken = accessToken
}

// Auth AuthUseCase method
func (auc *AuthUseCase) Auth(email string, password string) string {
	auc.Email = email
	auc.Password = password

	return auc.AccessToken
}
