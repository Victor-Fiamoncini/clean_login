package updateaccesstokenrepository

import (
	"github.com/Victor-Fiamoncini/auth_clean_architecture/src/domain/entities"
	shared_custom_errors "github.com/Victor-Fiamoncini/auth_clean_architecture/src/shared/errors"
	"github.com/go-pg/pg/v10"
)

// UpdateAccessTokenRepository struct
type UpdateAccessTokenRepository struct {
	UserID      string
	AccessToken string
	Database    *pg.DB
}

// NewUpdateAccessTokenRepository func
func NewUpdateAccessTokenRepository(db *pg.DB) IUpdateAccessTokenRepository {
	return &UpdateAccessTokenRepository{
		Database: db,
	}
}

// GetUserID UpdateAccessTokenRepository method
func (uatr *UpdateAccessTokenRepository) GetUserID() string {
	return uatr.UserID
}

// SetUserID UpdateAccessTokenRepository method
func (uatr *UpdateAccessTokenRepository) SetUserID(userID string) {
	uatr.UserID = userID
}

// GetAccessToken UpdateAccessTokenRepository method
func (uatr *UpdateAccessTokenRepository) GetAccessToken() string {
	return uatr.AccessToken
}

// SetAccessToken UpdateAccessTokenRepository method
func (uatr *UpdateAccessTokenRepository) SetAccessToken(accessToken string) {
	uatr.AccessToken = accessToken
}

// Update UpdateAccessTokenRepository method
func (uatr *UpdateAccessTokenRepository) Update() shared_custom_errors.IDefaultError {
	if uatr.AccessToken == "" {
		return shared_custom_errors.NewMissingParamError("AccessToken")
	}

	if uatr.UserID == "" {
		return shared_custom_errors.NewMissingParamError("UserID")
	}

	user := entities.NewUser()

	err := uatr.Database.Model(user).Where("id = ?", uatr.UserID).First()

	if err != nil {
		return shared_custom_errors.NewDefaultError("UpdateAccessTokenRepository.Update()")
	}

	_, err = uatr.Database.Model(user).Set("access_token = ?", uatr.AccessToken).Where("id = ?", uatr.UserID).Update()

	if err != nil {
		return shared_custom_errors.NewDefaultError("UpdateAccessTokenRepository.Update()")
	}

	return nil
}
