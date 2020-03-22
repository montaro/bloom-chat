package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3" // import your used driver
)
var o orm.Ormer

func init() {
	// register Model
	orm.RegisterModel(new(Room), new(Sender), new(ImageSize), new(Permission), new(Message))

	// set default database
	//TODO Handle error
	_ = orm.RegisterDataBase("default", "sqlite3", "vex.db", 30)

	// create table
	//TODO Handle error
	_ = orm.RunSyncdb("default", true, true)

	o = orm.NewOrm()
}
