package protocol

import "github.com/bloom-chat/internal/util"

type Envelop interface {
	String() string
}

type Message struct {
	RoomId     util.UUID `json:"roomId"`
	Msg        string
	SenderId   util.UUID
	SenderName string
}
