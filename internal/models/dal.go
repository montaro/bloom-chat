package models

import (
	"log"
	"sync"
	"time"

	"github.com/astaxie/beego/orm"
)

var o orm.Ormer
var ormerOnce sync.Once

func NewOrmer() orm.Ormer {
	ormerOnce.Do(func() {
		o = orm.NewOrm()
	})
	return o
}

func (room *Room) SaveRoom() *Room {
	o = NewOrmer()
	room.CreatedAt = time.Now()
	room.UpdatedAt = time.Now()
	roomId, err := o.Insert(room)
	if err != nil {
		log.Panicf("Saving a room to DB failed with error: %v\n", err)
	}
	room.Id = roomId
	return room
}
