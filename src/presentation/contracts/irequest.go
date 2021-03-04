package contracts

import "github.com/Victor-Fiamoncini/auth_clean_architecture/src/shared/types"

// IRequest struct
type IRequest interface {
	GetBody() types.StringMap
	SetBody(body types.StringMap)
}
