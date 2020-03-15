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

type Ack struct {
	Done bool `json:"done"`
}

func (ack *Ack) String() string {
	return repr.Repr(ack)
}

type CreateRoomResponse struct {
	RoomId int `json:"roomId"`
}

func (createRoomResponse *CreateRoomResponse) String() string {
	return repr.Repr(createRoomResponse)
}

//type SendMessageResponse struct {
//	RoomId  util.UUID
//	Message string
//}
//
//func (sendMessageResponse *SendMessageResponse) String() string {
//	return repr.Repr(sendMessageResponse)
//}

//type SetUserNameResponse struct {
//	Topic string
//}
//
//func (setUserNameResponse *SetUserNameResponse) String() string {
//	return repr.Repr(setUserNameResponse)
//}

//type SetRoomTopicResponse struct {
//	RoomId util.UUID
//	Topic   string
//}
//
//func (setRoomTopicResponse *SetRoomTopicResponse) String() string {
//	return repr.Repr(setRoomTopicResponse)
//}

//type JoinRoomResponse struct {
//	roomId util.UUID
//}
//
//func (joinRoomResponse *JoinRoomResponse) String() string {
//	return repr.Repr(joinRoomResponse)
//}

//type LeaveRoomResponse struct {
//	roomId util.UUID
//}
//
//func (leaveRoomResponse *LeaveRoomResponse) String() string {
//	return repr.Repr(leaveRoomResponse)
//}

//type BeginTypingResponse struct {
//	roomId util.UUID
//}
//
//func (beginTypingResponse *BeginTypingResponse) String() string {
//	return repr.Repr(beginTypingResponse)
//}
//
//type StopTypingResponse struct {
//	roomId util.UUID
//}
//
//func (stopTypingResponse *StopTypingResponse) String() string {
//	return repr.Repr(stopTypingResponse)
//}
