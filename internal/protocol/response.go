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
	UserID util.UUID `json:"client_id"`
}

func (clientConnectedResponse *ClientConnectedResponse) String() string {
	return repr.Repr(clientConnectedResponse)
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func (errorResponse *ErrorResponse) String() string {
	return repr.Repr(errorResponse)
}

type Handshake struct {
	ProtocolVersion	float64
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

type ListRoomsResponse struct {
	Rooms map[int64]string `json:"rooms"`
}

func (listRoomsResponse *ListRoomsResponse) String() string {
	return repr.Repr(listRoomsResponse)
}
