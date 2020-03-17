package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3" // import your used driver
)

func init() {
	// register Model
	orm.RegisterModel(new(Room), new(Message))

	// set default database
	//TODO Handle error
	_ = orm.RegisterDataBase("default", "sqlite3", "vex.db", 30)

	// create table
	//TODO Handle error
	_ = orm.RunSyncdb("default", false, true)

	//o := orm.NewOrm()
	//
	//// insert
	//room1 := Room{
	//	Topic: "ShittyChat",
	//}
	//
	//message1 := Message{
	//	Content: "A msg",
	//	Room:    &room1,
	//}
	//room1Id, err := o.Insert(&room1)
	//messageId1, err := o.Insert(&message1)
	//
	//fmt.Println("msg 1 ID", messageId1)
	//fmt.Println("room 1 ID", room1Id)
	//
	//// update
	//message1.Content = "Hello World!"
	//num, err := o.Update(&message1)
	//
	//// read one
	//message1Updated := Message{Id: message1.Id}
	//err = o.Read(&message1Updated)
	//
	//// delete
	////num, err = o.Delete(&message1Updated)
	//fmt.Println(num)
	//fmt.Println(err)
}
