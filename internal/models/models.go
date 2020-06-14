package models

import (
	"sync"
	"time"
)

var mutex = &sync.Mutex{}

// Base Model Struct
type Model struct {
	Id        int64     `orm:"auto" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `orm:"null" json:"deleted_at,omitempty"`
}

type UserVisual struct {
	Model
	DisplayName string `json:"display_name"`
	Color       string `json:"color"`
	Photo       string `json:"photo"`
}

//User in a room
type User struct {
	Model
	//SessionID *managers.Session    `orm:"rel(fk)" json:"session_id"`
	Room   *Room       `orm:"rel(fk)" json:"room_id"`
	Visual *UserVisual `orm:"rel(fk)" json:"visual"`
	//UserAgent UserAgent   `json:"user_agent"`
}

type Topic struct {
	Model
	Text  string `json:"text"`
	SetBy *User  `orm:"rel(fk)" json:"set_by"`
}

type Room struct {
	Model
	Handle     string         `json:"handle"`
	//Owner      *User          `orm:"rel(fk)" json:"sender"`
	Topic      *Topic         `orm:"rel(fk)" json:"topic"`
	Clients    []*chan string `orm:"-" json:"-"`
	MessagesCh chan string    `orm:"-" json:"-"`
}

func (room *Room) Broadcast() {
	for {
		select {
		case msg := <-room.MessagesCh:
			for _, client := range room.Clients {
				*client <- msg
			}
		}
	}
}

func (room *Room) JoinClient(clientCh *chan string) {
	mutex.Lock()
	room.Clients = append(room.Clients, clientCh)
	defer mutex.Unlock()
}

type MessageKind string

const (
	MessageText          MessageKind = "MSG_TXT"
	MessagePhoto         MessageKind = "MSG_PHOTO"
	MessagePhotoAnimated MessageKind = "MSG_PHOTO_ANIM"
	MessageVideo         MessageKind = "MSG_VIDEO"
)

type ImageSize struct {
	Model
	URl     string   `json:"url"`
	Width   string   `json:"width"`
	Height  string   `json:"height"`
	Message *Message `orm:"rel(fk)"`
}

type Message struct {
	Model
	Room             *Room        `orm:"rel(fk)"`
	Kind             MessageKind  `json:"kind"`
	Content          string       `json:"content"`
	FormattedContent string       `json:"formatted_content"`
	Sender           *User        `orm:"rel(fk)" json:"sender"`
	Sizes            []*ImageSize `orm:"reverse(many)" json:"sizes,omitempty"`
}
