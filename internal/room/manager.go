package room

import (
	"github.com/bloom-chat/internal/client"
	"github.com/bloom-chat/internal/util"
)

type Manager struct {
	rooms         map[util.UUID]*Room
	clientsManger *client.Manager
}

func NewManager() *Manager {
	rooms := make(map[util.UUID]*Room)
	return &Manager{rooms: rooms}
}

func (manager *Manager) CreateRoom(clientsManager *client.Manager, topic string) *Room {
	room := &Room{
		Id:             util.GenerateID(),
		topic:          topic,
		clientsManager: clientsManager,
		clients:        make(map[util.UUID]*client.Client),
		MessagesCh:     make(chan string),
	}
	mutex.Lock()
	manager.rooms[room.Id] = room
	mutex.Unlock()
	return room
}
