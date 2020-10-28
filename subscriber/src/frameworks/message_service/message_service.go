package frameworks

import (
	"log"

	"github.com/streadway/amqp"

	interfaces "sub/src/interfaces"
)

// IMessageService ...
type IMessageService interface {
	Connect() error
	Sub(controller interfaces.IController) error
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

func (s *messageServicer) Sub(controller interfaces.IController) error {

	msgs, err := (*s.channel).Consume(
		"hello", // queue
		"",      // consumer
		true,    // auto-ack
		false,   // exclusive
		false,   // no-local
		false,   // no-wait
		nil,     // args
	)
	if err != nil {
		return err
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			controller.Handle(d.Body)
		}
	}()
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

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
