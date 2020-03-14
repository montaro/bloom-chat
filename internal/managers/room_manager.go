package managers

import (
	"errors"
	"sync"

	"github.com/bloom-chat/internal/models"
	"github.com/bloom-chat/internal/util"
)

var roomOnce sync.Once

type RoomManager struct {
	rooms         map[int]*models.Room
	clientsManger *RoomManager
}

var roomManager *RoomManager

func NewRoomManager() *RoomManager {
	roomOnce.Do(func() {
		rooms := make(map[int]*models.Room)
		roomManager = &RoomManager{rooms: rooms}
	})
	return roomManager
}

func (manager *RoomManager) createRoom(topic string) *models.Room {
	room := &models.Room{
		Topic:      topic,
		Clients:    make(map[util.UUID]chan string),
		MessagesCh: make(chan string),
	}
	mutex.Lock()
	manager.rooms[room.Id] = room
	mutex.Unlock()
	go room.Broadcast()
	return room
}

func (manager *RoomManager) getRoom(roomId int) (*models.Room, error) {
	mutex.Lock()
	room, ok := manager.rooms[roomId]
	mutex.Unlock()
	if !ok {
		return nil, errors.New("room not found")
	} else {
		return room, nil
	}
}
