package managers

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/bloom-chat/internal/protocol"
)

func (client *Client) handleRequestMessage(requestMessageData *protocol.SendMessageRequest) {
	r := requestMessageData.RoomId
	room, ok := client.RoomsChs[r]
	if !ok {
		client.IncomingMessagesCh <- fmt.Sprintf("unknown or closed room: %s", r)
	} else {
		room <- requestMessageData.Message
	}
	log.Printf("Request Message cmd received:\n%s", requestMessageData.String())
	//TODO Remove, return response instead
	client.IncomingMessagesCh <- string(requestMessageData.RoomId)
	client.IncomingMessagesCh <- requestMessageData.Message

}

func (client *Client) handleCreateRoom(createRoomData *protocol.CreateRoomRequest) {
	room := roomManager.createRoom(createRoomData.Topic)
	log.Printf("Create Room cmd received:\n%s", createRoomData.String())
	createRoomResponse := protocol.CreateRoomResponse{
		RoomId: room.id,
	}
	bits, err := json.Marshal(createRoomResponse)
	if err != nil {
		client.returnSystemError(err)
	} else {
		client.IncomingMessagesCh <- string(bits)
	}
}

func (client *Client) handleSetUserName(setUserNameData *protocol.SetUserNameRequest) {
	client.name = setUserNameData.Name
	log.Printf("Set User Topic cmd received:\n%s", setUserNameData.String())
	//TODO Remove
	client.IncomingMessagesCh <- client.name
}

func (client *Client) handleSetRoomTopic(setRoomTopicData *protocol.SetRoomTopicRequest) {
	room, err := roomManager.getRoom(setRoomTopicData.RoomId)
	if err != nil {
		client.returnClientError(err)
	} else {
		room.topic = setRoomTopicData.Topic
		log.Printf("Set Room Topic cmd received:\n%s", setRoomTopicData.String())
		//TODO Remove
		client.IncomingMessagesCh <- room.topic
	}
}

func (client *Client) returnClientError(err error) {
	errorResponse := &protocol.ErrorResponse{Error: err.Error()}
	response := protocol.Response{
		Data: errorResponse,
	}
	msg, _ := json.Marshal(response)
	streamMsg := string(msg)
	client.IncomingMessagesCh <- streamMsg
}

func (client *Client) returnParseDataError(err error) {
	client.IncomingMessagesCh <- fmt.Sprintf("failed to parse data with error:\n%s", err.Error())
}

func (client *Client) returnSystemError(err error) {
	//TODO enhance error handling
	client.IncomingMessagesCh <- err.Error()
}
