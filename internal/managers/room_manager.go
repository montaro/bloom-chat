package managers

import (
	"errors"
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

func (manager *RoomManager) createRoom(topic string) *Room {
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

func (manager *RoomManager) getRoom(roomId util.UUID) (*Room, error) {
	mutex.Lock()
	room, ok := manager.rooms[roomId]
	mutex.Unlock()
	if !ok {
		return nil, errors.New("room not found")
	} else {
		return room, nil
	}
}
