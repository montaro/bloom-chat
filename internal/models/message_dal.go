package models

import (
	"log"
	"time"
)

func (message *Message) SaveMessage() *Message {
	message.CreatedAt = time.Now()
	message.UpdatedAt = time.Now()
	messageId, err := o.Insert(message)
	if err != nil {
		log.Panicf("Saving a message to DB failed with error: %v\n", err)
	}
	message.Id = messageId
	return message
}
