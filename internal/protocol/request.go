package protocol

import (
	"github.com/bloom-chat/internal/util"
	"github.com/hackebrot/go-repr/repr"
)

type RequestOperation string

const (
	Initialize   RequestOperation = "INITIALIZE"
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

type InitializeRequest struct {
	ProtocolVersion float64
}

func (initializeRequest *InitializeRequest) String() string {
	return repr.Repr(initializeRequest)
}

type CreateRoomRequest struct {
	Topic string
}

func (createRoomRequest *CreateRoomRequest) String() string {
	return repr.Repr(createRoomRequest)
}

type SendMessageRequest struct {
	RoomId  int64
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
	RoomId int64
	Topic  string
}

func (setRoomTopicRequest *SetRoomTopicRequest) String() string {
	return repr.Repr(setRoomTopicRequest)
}

type JoinRoomRequest struct {
	RoomId int64
}

func (joinRoomRequest *JoinRoomRequest) String() string {
	return repr.Repr(joinRoomRequest)
}

type LeaveRoomRequest struct {
	roomId int64
}

func (leaveRoomRequest *LeaveRoomRequest) String() string {
	return repr.Repr(leaveRoomRequest)
}

type BeginTypingRequest struct {
	roomId int64
}

func (beginTypingRequest *BeginTypingRequest) String() string {
	return repr.Repr(beginTypingRequest)
}

type StopTypingRequest struct {
	roomId int64
}

func (stopTypingRequest *StopTypingRequest) String() string {
	return repr.Repr(stopTypingRequest)
}
