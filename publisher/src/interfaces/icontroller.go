package interfaces

// IController ...
type IController interface {
	Handle(data []byte) bool
}
