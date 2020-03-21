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

//type Message struct {
//	Model
//	Content  string `orm:"size(10000)"`
//	Room     *Room  `orm:"rel(fk)"`
//	SenderId util.UUID
//}

type MessageKind string

const (
	MessageText          MessageKind = "MSG_TXT"
	MessagePhoto         MessageKind = "MSG_PHOTO"
	MessagePhotoAnimated MessageKind = "MSG_PHOTO_ANIM"
)

type ClientDevice string

const (
	ClientWeb    ClientDevice = "WEB"
	ClientMobile ClientDevice = "Mobile"
)

type ReplyTo struct {
	Id int64 `json:"id"`
}

type Permission string

const (
	PermissionCanDelete Permission = "CAN_DELETE"
)

type Sender struct {
	Id     util.UUID    `json:"id"`
	Name   string       `json:"name"`
	Status string       `json:"status"`
	Client ClientDevice `json:"client"`
}

type ImageSize struct {
	URl    string `json:"url"`
	Width  string `json:"width"`
	Height string `json:"height"`
}

type Message struct {
	Model
	Id               int64         `json:"id"`
	Room             *Room         `orm:"rel(fk)"`
	Kind             MessageKind   `json:"kind"`
	Content          string        `json:"content"`
	FormattedContent string        `json:"formatted_content"`
	Timestamp        int64         `json:"timestamp"`
	Status           string        `json:"status"`
	Sender           Sender        `json:"sender"`
	ReplyTo          ReplyTo       `json:"reply_to"`
	Permissions      []*Permission `json:"permissions"`
	Sizes            []*ImageSize  `json:"sizes"`
}
