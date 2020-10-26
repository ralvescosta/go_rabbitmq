package interfaces

import "log"

type subscriptController struct{}

func (*subscriptController) Handle(data []byte) bool {
	log.Printf("Received a message: %s", data)
	return true
}

// SubscriptController ...
func SubscriptController() IController {
	return &subscriptController{}
}
