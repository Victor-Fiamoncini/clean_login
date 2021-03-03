package contracts

// IController interface
type IController interface {
	Handle(httpRequest IHTTPRequest) IHTTPResponse
}
