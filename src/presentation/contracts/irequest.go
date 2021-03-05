package contracts

// IRequest struct
type IRequest interface {
	GetBody() interface{}
	SetBody(body interface{})
}
