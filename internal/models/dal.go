package models

import (
	"sync"

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

func save(room *Room) {
	o = NewOrmer()
	_, _ = o.Insert(room)
}
