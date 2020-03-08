package protocol

import (
	"github.com/bloom-chat/internal/util"
	"github.com/hackebrot/go-repr/repr"
)

type RequestOperation string

const (
	CreateRoom   RequestOperation = "CREATE_ROOM"
	SendMessage  RequestOperation = "SEND_MSG"
	SetUserName  RequestOperation = "SET_USER_NAME"
	SetRoomTopic RequestOperation = "SET_ROOM_TOPIC"
	JoinRoom     RequestOperation = "JOIN_ROOM"
)

type Request struct {
	RequestId util.UUID              `json:"request_id"`
	Op        RequestOperation       `json:"op"`
	Data      map[string]interface{} `json:"data"`
}

//func (request *Request) decode() (RequestData, error) {
//	switch request.Op {
//	case createRoom:
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

type CreateRoomRequest struct {
	Topic string
}

func (createRoomRequest *CreateRoomRequest) String() string {
	return repr.Repr(createRoomRequest)
}

type SendMessageRequest struct {
	RoomId  util.UUID
	Message string
}

func (sendMessageRequest *SendMessageRequest) String() string {
	return repr.Repr(sendMessageRequest)
}

type SetUserNameRequest struct {
	Name string
}

func (setUserNameRequest *SetUserNameRequest) String() string {
	return repr.Repr(setUserNameRequest)
}

type SetRoomTopicRequest struct {
	RoomId util.UUID
	Topic  string
}

func (setRoomTopicRequest *SetRoomTopicRequest) String() string {
	return repr.Repr(setRoomTopicRequest)
}

type JoinRoomRequest struct {
	RoomId util.UUID
}

func (joinRoomRequest *JoinRoomRequest) String() string {
	return repr.Repr(joinRoomRequest)
}

type LeaveRoomRequest struct {
	roomId util.UUID
}

func (leaveRoomRequest *LeaveRoomRequest) String() string {
	return repr.Repr(leaveRoomRequest)
}

type BeginTypingRequest struct {
	roomId util.UUID
}

func (beginTypingRequest *BeginTypingRequest) String() string {
	return repr.Repr(beginTypingRequest)
}

type StopTypingRequest struct {
	roomId util.UUID
}

func (stopTypingRequest *StopTypingRequest) String() string {
	return repr.Repr(stopTypingRequest)
}
