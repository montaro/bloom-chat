package managers

import (
	"sync"

	"github.com/bloom-chat/internal/util"
)

var roomOnce sync.Once

type RoomManager struct {
	rooms         map[util.UUID]*Room
	clientsManger *ClientManager
}

var roomManager *RoomManager

func NewRoomManager() *RoomManager {
	roomOnce.Do(func() {
		rooms := make(map[util.UUID]*Room)
		roomManager = &RoomManager{rooms: rooms}
	})
	return roomManager
}

func (manager *RoomManager) CreateRoom(topic string) *Room {
	room := &Room{
		id:         util.GenerateID(),
		topic:      topic,
		clients:    make(map[util.UUID]*Client),
		messagesCh: make(chan string),
	}
	mutex.Lock()
	manager.rooms[room.id] = room
	mutex.Unlock()
	return room
}
