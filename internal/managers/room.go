package managers

import (
	"github.com/bloom-chat/internal/util"
	"log"
	"sync"
)

var mutex = &sync.Mutex{}

type Room struct {
	id         util.UUID
	topic      string
	clients    map[util.UUID]*Client
	messagesCh chan string
}

func (room *Room) Broadcast() {
	for {
		select {
		case msg := <-room.messagesCh:
			for _, client := range room.clients {
				client.IncomingMessagesCh <- msg
			}
		}
	}
}

func (room *Room) JoinClient(clientId util.UUID) error {
	client, err := clientManager.GetClient(clientId)
	if err != nil {
		log.Printf("Room: %s-%s, failed to join client: %s on error: %v",
			room.id, room.topic, clientId, err)
		return err
	}
	mutex.Lock()
	room.clients[clientId] = client
	defer mutex.Unlock()
	return nil
}
