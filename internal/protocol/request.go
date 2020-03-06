package protocol

import "github.com/bloom-chat/internal/util"

type RequestOperation string

const (
	CreateRoom     RequestOperation = "CREATE_ROOM"
	RequestMessage RequestOperation = "REQ_MSG"
	SET_NAME       RequestOperation = "SET_NAME"
	SET_ROOM_NAME  RequestOperation = "SET_ROOM_NAME"
)

type Request struct {
	Op   RequestOperation       `json:"op"`
	Data map[string]interface{} `json:"data"`
}

type CreateRoomData struct {
	Name   string
}

type RequestMessageData struct {
	To      util.UUID
	Message string
}

type SetName struct {
	Name string
}

type SetRoomName struct {
	RoomId util.UUID
	Name   string
}
