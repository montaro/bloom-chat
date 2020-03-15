package models

import (
	"sync"
	"time"

	"github.com/bloom-chat/internal/util"
)

var mutex = &sync.Mutex{}

// Model Struct
type Model struct {
	Id        int64 `orm:"auto"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `orm:"null"`
}

type Room struct {
	Model
	Topic      string                    `orm:"size(100)"`
	Clients    map[util.UUID]chan string `orm:"-"`
	MessagesCh chan string               `orm:"-"`
}

func (room *Room) Broadcast() {
	for {
		select {
		case msg := <-room.MessagesCh:
			for _, clientCh := range room.Clients {
				clientCh <- msg
			}
		}
	}
}

func (room *Room) JoinClient(clientId util.UUID, clientCh chan string) {
	mutex.Lock()
	room.Clients[clientId] = clientCh
	defer mutex.Unlock()
}

type Message struct {
	Model
	Content  string `orm:"size(10000)"`
	Room     *Room  `orm:"rel(fk)"`
	SenderId util.UUID
}
