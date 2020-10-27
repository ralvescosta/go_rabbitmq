package interfaces

import (
	"encoding/json"
	"log"

	usecase "pub/src/application/usecases"
	entities "pub/src/business/entities"
)

type subscriptController struct {
	usecase usecase.IOutputAmqpDataUsecase
}

func (s *subscriptController) Handle(data []byte) bool {
	receivedData := &entities.AmqpDataReceivedEntity{}

	err := json.Unmarshal(data, receivedData)

	if err != nil {
		log.Println("Error on Unmarshal")
		return true
	}

	result := s.usecase.Output(receivedData)

	return result
}

// SubscriptController ...
func SubscriptController(usecase usecase.IOutputAmqpDataUsecase) IController {
	return &subscriptController{usecase: usecase}
}
