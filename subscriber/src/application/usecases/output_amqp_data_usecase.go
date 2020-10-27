package application

import (
	"log"
	entities "sub/src/business/entities"
)

// IOutputAmqpDataUsecase ...
type IOutputAmqpDataUsecase interface {
	Output(data *entities.AmqpDataReceivedEntity) bool
}

type usecase struct{}

func (*usecase) Output(data *entities.AmqpDataReceivedEntity) bool {
	log.Println("Print on Usecase: ", data)

	return true
}

// OutputAmqpDataUsecase ...
func OutputAmqpDataUsecase() IOutputAmqpDataUsecase {
	return &usecase{}
}
