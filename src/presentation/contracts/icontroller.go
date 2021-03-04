package contracts

// IController interface
type IController interface {
	Handle(httpRequest IRequest) IResponse
}
