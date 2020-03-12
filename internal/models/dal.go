package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/bloom-chat/internal/managers"
)

type RoomDAL struct {
	o orm.Ormer
}

func (roomDAL *RoomDAL) save(room *managers.Room) {
	o := orm.NewOrm()
	_, _ = o.Insert(room)
}
