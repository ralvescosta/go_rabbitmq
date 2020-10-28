package frameworks

import (
	"github.com/streadway/amqp"
)

// IMessageService ...
type IMessageService interface {
	Connect() error
	Pub(exchange string, routingKey string, mandatory bool, immediate bool, data []byte) error
	Defer()
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

func (s *messageServicer) Pub(exchange string, routingKey string, mandatory bool, immediate bool, data []byte) error {

	err := s.channel.Publish(
		exchange,   // exchange = ""
		routingKey, // routing key = "queue name"
		mandatory,  // mandatory = false
		immediate,  // immediate = false
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        data,
		})

	if err != nil {
		return err
	}

	return nil
}

func (s *messageServicer) Defer() {
	defer s.channel.Close()
	defer s.conn.Close()
}

// MessageService ..
func MessageService() IMessageService {
	return &messageServicer{}
}
