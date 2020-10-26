package main

import (
	"log"

	messageService "sub/src/frameworks/message_service"
	interfaces "sub/src/interfaces"
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

	m.registerAMQPSubscriptions()

	_msgService.Defer()
}

// RegisterAMQPSubscriptions ...
func (m *AppModule) registerAMQPSubscriptions() {
	_controller := interfaces.SubscriptController()

	m.messageService.Sub(_controller)
}
