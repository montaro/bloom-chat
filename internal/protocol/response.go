package protocol

import (
	"github.com/bloom-chat/internal/util"
	"github.com/hackebrot/go-repr/repr"
)

type ResponseOperation string

type Response struct {
	RequestId util.UUID `json:"request_id"`
	Data      Envelop   `json:"data"`
}

type ClientConnectedResponse struct {
	SessionID   util.UUID `json:"session_id"`
	DisplayName string    `json:"display_name"`
}

func (clientConnectedResponse *ClientConnectedResponse) String() string {
	return repr.Repr(clientConnectedResponse)
}

type ResponseError struct {
	Msg  string `json:"msg"`
	Code uint   `json:"code"`
}

func (responseError *ResponseError) Error() string {
	return repr.Repr(responseError)
}

func (responseError *ResponseError) String() string {
	return repr.Repr(responseError)
}

type Handshake struct {
	ProtocolVersion float64
}

func (handshake *Handshake) String() string {
	return repr.Repr(handshake)
}

type Ack struct {
	Done bool `json:"done"`
}

func (ack *Ack) String() string {
	return repr.Repr(ack)
}

type CreateRoomResponse struct {
	RoomId int64 `json:"roomId"`
}

func (createRoomResponse *CreateRoomResponse) String() string {
	return repr.Repr(createRoomResponse)
}

type Room struct {
	Id    int64  `json:"id"`
	Topic string `json:"topic"`
}

type ListRoomsResponse struct {
	Rooms []*Room `json:"rooms"`
}

func (listRoomsResponse *ListRoomsResponse) String() string {
	return repr.Repr(listRoomsResponse)
}
