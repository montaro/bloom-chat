package managers

import (
	"sync"

	"github.com/bloom-chat/internal/models"
)

var roomOnce sync.Once

type RoomManager struct {
	rooms         map[int64]*models.Room
	clientsManger *RoomManager
}

var roomManager *RoomManager

func NewRoomManager() *RoomManager {
	roomOnce.Do(func() {
		rooms := make(map[int64]*models.Room)
		roomManager = &RoomManager{rooms: rooms}
	})
	return roomManager
}

//func (manager *RoomManager) createRoom(topic string, client *Client) *models.Room {
//	room := &models.Room{
//		Topic:      topic,
//		Clients:    make(map[util.UUID]chan string),
//		MessagesCh: make(chan string),
//		Owner: &models.User{
//			ClientID: client.Id,
//			Name:     client.Name,
//			Status:   models.Online,
//			Client:   models.Web,
//		},
//	}
//	room = room.SaveRoom()
//	mutex.Lock()
//	manager.rooms[room.Id] = room
//	mutex.Unlock()
//	go room.Broadcast()
//	return room
//}
//
//func (manager *RoomManager) getRoom(roomId int64) (*models.Room, error) {
//	mutex.Lock()
//	room, ok := manager.rooms[roomId]
//	mutex.Unlock()
//	if !ok {
//		return nil, errors.New("room not found")
//	} else {
//		return room, nil
//	}
//}
//
//func (manager *RoomManager) listRoomsIDs() ([]int64, error) {
//	roomsIDs := make([]int64, len(manager.rooms))
//	i := 0
//	for k, _ := range manager.rooms {
//		roomsIDs[i] = k
//		i++
//	}
//	return roomsIDs, nil
//}
