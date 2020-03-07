package protocol

import (
	"github.com/hackebrot/go-repr/repr"

	"github.com/bloom-chat/internal/util"
)

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

//func (request *Request) decode() (RequestData, error) {
//	switch request.Op {
//	case CreateRoom:
//		requestMessageData := &RequestMessageData{}
//		err := mapstructure.Decode(request.Data, requestMessageData)
//		if err != nil {
//			return nil, err
//		} else {
//			return requestMessageData, nil
//		}
//	default:
//		return nil, nil
//	}
//}

type CreateRoomData struct {
	Topic string
}

func (createRoomData *CreateRoomData) String() string {
	return repr.Repr(createRoomData)
}

type RequestMessageData struct {
	To      util.UUID
	Message string
}

func (requestMessageData *RequestMessageData) String() string {
	return repr.Repr(requestMessageData)
}

type SetName struct {
	Name string
}

type SetRoomName struct {
	RoomId util.UUID
	Name   string
}
