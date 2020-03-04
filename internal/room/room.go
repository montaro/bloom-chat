package room

import "github.com/bloom-chat/internal/client"

type Room struct {
	id uint64
	clients []client.Id
	messages chan string
}

func (room *Room) broadcast() {
	for {
		select {
		case msg := <-room.messages:
			for _, client := range room.clients {
				client.Messages <- msg
			}
		}
	}
}