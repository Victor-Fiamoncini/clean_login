package errors

// IDefaultError interface
type IDefaultError interface {
	GetName() string
	SetName(name string)
	GetError() error
	SetError(err error)
}
