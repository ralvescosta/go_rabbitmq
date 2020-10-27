package application

import (
	"testing"

	entities "pub/src/business/entities"

	"github.com/stretchr/testify/assert"
)

type mocksStruct struct {
	receivedEntity *entities.AmqpDataReceivedEntity
}

func makeMocks() *mocksStruct {

	receivedEntity := &entities.AmqpDataReceivedEntity{}

	return &mocksStruct{
		receivedEntity: receivedEntity,
	}
}

func TestOutput(t *testing.T) {

	mocks := makeMocks()

	sut := OutputAmqpDataUsecase()

	result := sut.Output(mocks.receivedEntity)

	assert.Equal(t, true, result)
}
