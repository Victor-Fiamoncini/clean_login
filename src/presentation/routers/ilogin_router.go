package routers

import "github.com/Victor-Fiamoncini/auth_clean_architecture/src/presentation/helpers"

// ILoginRouter interface
type ILoginRouter interface {
	Route(httpRequest *helpers.HTTPRequest) helpers.IHTTPResponse
}
