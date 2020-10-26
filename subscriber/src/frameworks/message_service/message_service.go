package frameworks

import (
	"github.com/streadway/amqp"
)

// IMessageService ...
type IMessageService interface {
	Connect() error
	Pub() error
}

type messageServicer struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

// Connect ...
func (s *messageServicer) Connect() error {
	conn, err := amqp.Dial("amqp://rabbitmq:123456@localhost:5672/")
	if err != nil {
		return err
	}

	ch, err := conn.Channel()
	if err != nil {
		return err
	}

	_, err = ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		return err
	}

	s.conn = conn
	s.channel = ch

	return nil
}

func (s *messageServicer) Pub() error {
	return nil
}

// MessageService ..
func MessageService() IMessageService {
	return &messageServicer{}
}
