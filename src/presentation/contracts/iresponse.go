package contracts

import (
	shared_custom_errors "github.com/Victor-Fiamoncini/auth_clean_architecture/src/shared/errors"

	"github.com/Victor-Fiamoncini/auth_clean_architecture/src/shared/types"
)

// IResponse interface
type IResponse interface {
	GetStatusCode() int
	SetStatusCode(statusCode int)
	GetErrorObject() error
	SetErrorObject(errorObject error)
	GetErrorName() string
	SetErrorName(errorName string)
	GetBody() types.Map
	SetBody(body types.Map)

	Success(data types.Map) IResponse
	BadRequest(err shared_custom_errors.IDefaultError) IResponse
	Unauthorized() IResponse
	ServerError() IResponse
}
