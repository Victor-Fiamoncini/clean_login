package updateaccesstokenrepository

import (
	"context"

	shared_custom_errors "github.com/Victor-Fiamoncini/auth_clean_architecture/src/shared/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// UpdateAccessTokenRepository struct
type UpdateAccessTokenRepository struct {
	UserID      string
	AccessToken string
	UserModel   *mongo.Collection
}

// NewUpdateAccessTokenRepository func
func NewUpdateAccessTokenRepository(userModel *mongo.Collection) IUpdateAccessTokenRepository {
	return &UpdateAccessTokenRepository{
		UserModel: userModel,
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
	ctx := context.Background()

	defer ctx.Done()

	userObjectID, err := primitive.ObjectIDFromHex(uatr.UserID)

	if err != nil {
		return shared_custom_errors.NewDefaultError("U    pdateAccessTokenRepository.Update()")
	}

	_, err = uatr.UserModel.UpdateOne(
		ctx,
		bson.M{
			"_id": bson.M{
				"$eq": userObjectID,
			},
		},
		bson.M{
			"$set": bson.M{
				"access_token": uatr.AccessToken,
			},
		},
	)

	if err != nil {
		return shared_custom_errors.NewDefaultError("UpdateAccessTokenRepository.Update()")
	}

	return nil
}
