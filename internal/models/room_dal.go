package models

import (
	"log"
	"time"
)

func (room *Room) SaveRoom() *Room {
	room.CreatedAt = time.Now()
	room.UpdatedAt = time.Now()
	roomId, err := o.Insert(room)
	if err != nil {
		log.Panicf("Saving a room to DB failed with error: %v\n", err)
	}
	room.Id = roomId
	return room
}
