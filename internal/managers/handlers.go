package managers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/bloom-chat/internal/util"
	"log"

	"github.com/bloom-chat/internal/protocol"
)

func (client *Client) handleSendMessage(requestId util.UUID, requestMessageData *protocol.SendMessageRequest) {
	room, err := roomManager.getRoom(requestMessageData.RoomId)
	if err != nil {
		client.returnError(requestId, err)
	} else {
		//make sure client is a member in the room
		_, ok := room.clients[client.Id]
		if ok {
			message := protocol.Message{
				RoomId:     room.id,
				Msg:        requestMessageData.Message,
				SenderId:   client.Id,
				SenderName: client.name,
			}
			msg, _ := json.Marshal(message)
			streamMsg := string(msg)
			room.messagesCh <- streamMsg
			log.Printf("Send Message cmd received:\n%s", requestMessageData.String())
			client.returnAck(requestId)
		} else {
			client.returnForbiddenError(requestId	)
		}
	}
}

func (client *Client) handleCreateRoom(requestId util.UUID, createRoomData *protocol.CreateRoomRequest) {
	room := roomManager.createRoom(createRoomData.Topic)
	log.Printf("Create Room cmd received:\n%s", createRoomData.String())
	createRoomResponse := protocol.CreateRoomResponse{
		RoomId: room.id,
	}
	bits, err := json.Marshal(createRoomResponse)
	if err != nil {
		client.returnSystemError(requestId, err)
	} else {
		client.IncomingMessagesCh <- string(bits)
	}
}

func (client *Client) handleSetUserName(requestId util.UUID, setUserNameData *protocol.SetUserNameRequest) {
	client.name = setUserNameData.Name
	log.Printf("Set User Topic cmd received:\n%s", setUserNameData.String())
	client.returnAck(requestId)
}

func (client *Client) handleSetRoomTopic(requestId util.UUID, setRoomTopicData *protocol.SetRoomTopicRequest) {
	room, err := roomManager.getRoom(setRoomTopicData.RoomId)
	if err != nil {
		client.returnError(requestId, err)
	} else {
		room.topic = setRoomTopicData.Topic
		log.Printf("Set Room Topic cmd received:\n%s", setRoomTopicData.String())
		client.returnAck(requestId)
	}
}

func (client *Client) handleJoinRoom(requestId util.UUID, joinRoomRequest *protocol.JoinRoomRequest) {
	room, err := roomManager.getRoom(joinRoomRequest.RoomId)
	if err != nil {
		client.returnError(requestId, err)
	} else {
		room.JoinClient(client)
		log.Printf("Join Room cmd received:\n%s", joinRoomRequest.String())
		client.returnAck(requestId)
	}
}

func (client *Client) returnAck(requestId util.UUID) {
	ack := &protocol.Ack{Done: true}
	response := protocol.Response{
		RequestId: requestId,
		Data:      ack,
	}
	msg, _ := json.Marshal(response)
	streamMsg := string(msg)
	client.IncomingMessagesCh <- streamMsg
}

func (client *Client) returnError(requestId util.UUID, err error) {
	errorResponse := &protocol.ErrorResponse{Error: err.Error()}
	response := protocol.Response{
		RequestId: requestId,
		Data:      errorResponse,
	}
	msg, _ := json.Marshal(response)
	streamMsg := string(msg)
	client.IncomingMessagesCh <- streamMsg
}

func (client *Client) returnParseDataError(requestId util.UUID, err error) {
	client.returnError(requestId, errors.New(fmt.Sprintf("failed to parse data with error:\n%s",
		err.Error())))
}

func (client *Client) returnSystemError(requestId util.UUID, err error) {
	client.returnError(requestId, errors.New(fmt.Sprintf("internal system error:\n%s",
		err.Error())))
}

func (client *Client) returnForbiddenError(requestId util.UUID) {
	client.returnError(requestId, errors.New("forbidden action"))
}
