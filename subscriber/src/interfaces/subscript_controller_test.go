package interfaces

import (
	"testing"

	"github.com/stretchr/testify/assert"

	entities "sub/src/business/entities"
	stuntman "sub/src/interfaces/__test__"
)

type mocksStruct struct {
	handleInput    []byte
	receivedEntity *entities.AmqpDataReceivedEntity
	usecaseSpy     *stuntman.OutputAmqpDataUsecaseSpy
}

func makeMocks() *mocksStruct {
	handleInput := []byte{}

	receivedEntity := &entities.AmqpDataReceivedEntity{}

	usecaseSpy := &stuntman.OutputAmqpDataUsecaseSpy{}

	return &mocksStruct{
		handleInput:    handleInput,
		receivedEntity: receivedEntity,
		usecaseSpy:     usecaseSpy,
	}
}

func TestSubscriptController(t *testing.T) {
	mocks := makeMocks()

	mocks.usecaseSpy.On("Output", mocks.receivedEntity).Return(true)

	sut := SubscriptController(mocks.usecaseSpy)

	result := sut.Handle(mocks.handleInput)

	assert.Equal(t, true, result)
}
