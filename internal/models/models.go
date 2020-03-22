package models

import (
	"github.com/hackebrot/go-repr/repr"
	"sync"
	"time"

	"github.com/bloom-chat/internal/util"
)

var mutex = &sync.Mutex{}

// Model Struct
type Model struct {
	Id        int64      `orm:"auto" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `orm:"null" json:"deleted_at,omitempty"`
}

type Room struct {
	Model
	Topic      string                    `orm:"size(100)"`
	Clients    map[util.UUID]chan string `orm:"-" json:"-"`
	MessagesCh chan string               `orm:"-" json:"-"`
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

type MessageKind string

const (
	MessageText          MessageKind = "MSG_TXT"
	MessagePhoto         MessageKind = "MSG_PHOTO"
	MessagePhotoAnimated MessageKind = "MSG_PHOTO_ANIM"
)

type MessageStatus string

const (
	Seen      MessageStatus = "SEEN"
	Delivered MessageStatus = "DELIVERED"
)

type ClientDevice string

const (
	ClientWeb    ClientDevice = "WEB"
	ClientMobile ClientDevice = "Mobile"
)

type PermissionValue string

const (
	PermissionCanDelete PermissionValue = "CAN_DELETE"
	PermissionCanEdit   PermissionValue = "CAN_EDIT"
)

type Permission struct {
	Model
	Value   string
	Message []*Message `orm:"rel(m2m)"`
}

type ClientStatus string

const (
	Online  ClientStatus = "ONLINE"
	Offline ClientStatus = "OFFLINE"
	Away    ClientStatus = "AWAY"
	Busy    ClientStatus = "Busy"
)

type Sender struct {
	Model
	ClientID util.UUID    `json:"id"`
	Name     string       `json:"name"`
	Status   ClientStatus `json:"status"`
	Client   ClientDevice `json:"client"`
}

type ImageSize struct {
	Model
	URl     string   `json:"url"`
	Width   string   `json:"width"`
	Height  string   `json:"height"`
	Message *Message `orm:"rel(fk)"`
}

type Message struct {
	Model
	Room             *Room         `orm:"rel(fk)"`
	Kind             MessageKind   `json:"kind"`
	Content          string        `json:"content"`
	FormattedContent string        `json:"formatted_content" mapstructure:"formatted_content"`
	//TODO implement SeenBy
	//Status           MessageStatus `json:"status"`
	Sender           *Sender       `orm:"rel(fk)" json:"sender"`
	ReplyTo          *Message      `orm:"null;rel(fk)" json:"reply_to,omitempty"`
	Sizes            []*ImageSize  `orm:"reverse(many)" json:"sizes,omitempty"`
	//Permissions      []*Permission `orm:"reverse(many)";json:"permissions"`
}

func (message *Message) String() string {
	return repr.Repr(message)
}
