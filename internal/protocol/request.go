package protocol

import "github.com/bloom-chat/internal/util"

type RequestOperation string

const (
	REQUEST_MSG   RequestOperation = "REQ_MSG"
	SET_NAME      RequestOperation = "SET_NAME"
	SET_ROOM_NAME RequestOperation = "SET_ROOM_NAME"
)

type Request struct {
	Op   RequestOperation       `json:"op"`
	Data map[string]interface{} `json:"data"`
}

type RequestMessageData struct {
	To      string
	Message string
}

type SetNameData struct {
	Name string
}

type SetRoomNameData struct {
	RoomId util.UUID
	Name   string
}
