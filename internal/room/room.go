package room

import (
	"github.com/bloom-chat/internal/util"
	"log"
	"sync"
)

var mutex = &sync.Mutex{}

type Room struct {
	Id             util.UUID
	topic          string
	clientsManager *ClientManager
	clients        map[util.UUID]*Client
	MessagesCh     chan string
}

func (room *Room) Broadcast() {
	for {
		select {
		case msg := <-room.MessagesCh:
			for _, client := range room.clients {
				client.IncomingMessagesCh <- msg
			}
		}
	}
}

func (room *Room) JoinClient(clientId util.UUID) error {
	client, err := room.clientsManager.GetClient(clientId)
	if err != nil {
		log.Printf("Room: %s-%s, failed to join client: %s on error: %v",
			room.Id, room.topic, clientId, err)
		return err
	}
	mutex.Lock()
	room.clients[clientId] = client
	defer mutex.Unlock()
	return nil
}
