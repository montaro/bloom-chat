package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3" // import your used driver
)

//func init() {
//	db, err := gorm.Open("sqlite3", "test.db")
//	if err != nil {
//		panic("failed to connect database")
//	}
//	defer db.Close()
//
//	// Migrate the schema
//	db.AutoMigrate(&Message{})
//	db.AutoMigrate(&Room{})
//
//	// Create
//	db.Create(&Room{Topic: "ShittyChat"})
//
//	// Read
//	var room Room
//	db.First(&room, 1) // find product with id 1
//	db.First(&room, "topic = ?", "ShittyChat") // find product with code l1212
//
//	// Update - update product's price to 2000
//	db.Model(&room).Update("Topic", "Cool Stuff")
//	//db.Model(&Message{}).AddForeignKey("RoomID", "rooms(id)", "RESTRICT", "RESTRICT")
//
//	fmt.Println(repr.Repr(&room))
//
//	db.Create(&Message{
//		Content: "Hello World!",
//		Room:    room,
//	})
//
//	db.Create(&Message{
//		Content: "Hey World!",
//		Room:    room,
//	})
//
//	var message Message
//
//	db.First(&message, "Content = ?", "Hey World!")
//
//	fmt.Println(repr.Repr(&message.Room))
//	// Delete - delete product
//	//db.Delete(&room)
//	//db.Delete(&message)
//}

func init() {
	// register Model
	orm.RegisterModel(new(Room), new(Message))

	// set default database
	//TODO Handle error
	_ = orm.RegisterDataBase("default", "sqlite3", "vex.db", 30)

	// create table
	//TODO Handle error
	_ = orm.RunSyncdb("default", true, true)

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
