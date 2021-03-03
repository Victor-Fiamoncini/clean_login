package contracts

// IHTTPRequest struct
type IHTTPRequest interface {
	GetBody() interface{}
	SetBody(body interface{})
}
