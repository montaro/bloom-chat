package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3" // import your used driver
)
var o orm.Ormer

func init() {
	// register Model
	orm.RegisterModel(
		new(Session),
		new(Client),
		new(UserVisual),
		new(User),
		new(Topic),
		new(Room),
		new(ImageSize),
		new(Message))

	// set default database
	//TODO Handle error
	_ = orm.RegisterDataBase("default", "sqlite3", "vex.db", 30)

	// create table
	//TODO Handle error
	_ = orm.RunSyncdb("default", false, true)

	o = orm.NewOrm()
}
