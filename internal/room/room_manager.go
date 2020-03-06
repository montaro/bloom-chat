package room

import (
	"github.com/bloom-chat/internal/util"
)

type RoomManager struct {
	rooms         map[util.UUID]*Room
	clientsManger *ClientManager
}

func NewRoomManager() *RoomManager {
	rooms := make(map[util.UUID]*Room)
	return &RoomManager{rooms: rooms}
}

func (manager *RoomManager) CreateRoom(clientsManager *ClientManager, topic string) *Room {
	room := &Room{
		Id:             util.GenerateID(),
		topic:          topic,
		clientsManager: clientsManager,
		clients:        make(map[util.UUID]*Client),
		MessagesCh:     make(chan string),
	}
	mutex.Lock()
	manager.rooms[room.Id] = room
	mutex.Unlock()
	return room
}
