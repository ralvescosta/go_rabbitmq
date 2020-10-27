package main

import (
	"encoding/json"
	"log"

	entities "pub/src/business/entities"
	messageService "pub/src/frameworks/message_service"
)

// AppModule ...
type AppModule struct {
	messageService messageService.IMessageService
}

// InitServer ...
func (m *AppModule) InitServer() {
	_msgService := messageService.MessageService()

	err := _msgService.Connect()
	if err != nil {
		log.Fatal(err)
	}
	m.messageService = _msgService

	m.registerTimeinterval()

	_msgService.Defer()
}

// RegisterAMQPSubscriptions ...
func (m *AppModule) registerTimeinterval() {

	entity := entities.AmqpDataReceivedEntity{
		FirstName: "Jon",
		LastName:  "Due",
	}

	amqpData, err := json.Marshal(entity)
	if err != nil {
		log.Println(err)
	}

	log.Println("Pub data to AMQP Broker...")
	m.messageService.Pub("", "hello", false, false, amqpData)

}
