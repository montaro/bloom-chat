package protocol

import (
	"github.com/bloom-chat/internal/models"
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
	ListRooms    RequestOperation = "LIST_ROOMS"
)

type Request struct {
	RequestId util.UUID              `json:"request_id"`
	Op        RequestOperation       `json:"op"`
	Data      map[string]interface{} `json:"data"`
}

type InitializeRequest struct {
	SessionID   util.UUID `mapstructure:"session_id"`
	DisplayName string    `mapstructure:"display_name"`
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

type ListRoomsRequest struct {
}

func (listRoomsRequest *ListRoomsRequest) String() string {
	return repr.Repr(listRoomsRequest)
}

type SendMessageRequest struct {
	RoomId  int64          `mapstructure:"room_id"`
	Message models.Message `mapstructure:"message"`
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
	Handle string `mapstructure:"handle"`
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
