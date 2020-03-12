package models

import (
	_ "github.com/mattn/go-sqlite3" // import your used driver
	"time"
)

// Model Struct
type model struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

type Room struct {
	model
	Id    int    `orm:"auto"`
	Topic string `orm:"size(100)"`
}

type Message struct {
	model
	Id      int    `orm:"auto"`
	Content string `orm:"size(10000)"`
	Room    *Room  `orm:"rel(fk)"`
}
