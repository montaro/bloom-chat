package managers

import (
	"errors"
	"sync"

	"github.com/bloom-chat/internal/models"
)

var roomOnce sync.Once

type RoomManager struct {
	rooms         map[string]*models.Room
	clientsManger *RoomManager
}

var roomManager *RoomManager

func NewRoomManager() *RoomManager {
	roomOnce.Do(func() {
		rooms := make(map[string]*models.Room)
		roomManager = &RoomManager{rooms: rooms}
	})
	return roomManager
}

func (manager *RoomManager) createRoom(handle string, topic *models.Topic) *models.Room {
	room := &models.Room{
		Handle:     handle,
		Topic:      topic,
		Clients:    make([]*chan string, 5),
		MessagesCh: make(chan string),
		//Owner: &models.User{
		//	ClientID: client.Id,
		//	Name:     client.Name,
		//	Status:   models.Online,
		//	Client:   models.Web,
		//},
	}
	room = room.SaveRoom()
	mutex.Lock()
	manager.rooms[room.Handle] = room
	mutex.Unlock()
	go room.Broadcast()
	return room
}

func (manager *RoomManager) getRoom(handle string) (*models.Room, error) {
	mutex.Lock()
	room, ok := manager.rooms[handle]
	mutex.Unlock()
	if !ok {
		return nil, errors.New("room not found")
	} else {
		return room, nil
	}
}

//func (manager *RoomManager) listRoomsIDs() ([]int64, error) {
//	roomsIDs := make([]int64, len(manager.rooms))
//	i := 0
//	for k, _ := range manager.rooms {
//		roomsIDs[i] = k
//		i++
//	}
//	return roomsIDs, nil
//}
