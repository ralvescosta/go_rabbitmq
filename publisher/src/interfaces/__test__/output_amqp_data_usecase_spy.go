package interfaces

import (
	entities "pub/src/business/entities"

	"github.com/stretchr/testify/mock"
)

// OutputAmqpDataUsecaseSpy ...
type OutputAmqpDataUsecaseSpy struct {
	mock.Mock
}

// Output ...
func (u *OutputAmqpDataUsecaseSpy) Output(data *entities.AmqpDataReceivedEntity) bool {
	args := u.Called()

	return args.Bool(0)
}
