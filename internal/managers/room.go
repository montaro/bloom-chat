package managers

import (
	"github.com/bloom-chat/internal/util"
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

func (room *Room) JoinClient(client *Client) {
	mutex.Lock()
	room.clients[client.Id] = client
	defer mutex.Unlock()
}
