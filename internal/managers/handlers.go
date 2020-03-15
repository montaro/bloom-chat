package managers

import (
	"encoding/json"
	"log"

	"github.com/bloom-chat/internal/models"
	"github.com/bloom-chat/internal/protocol"
	"github.com/bloom-chat/internal/util"
)

func (client *Client) handleConnect(requestId util.UUID, initializeRequest *protocol.InitializeRequest) {
	ClientConnectedResponse := &protocol.ClientConnectedResponse{
		UserID: client.Id,
	}
	ClientConnectedResponseWrapper := protocol.Response{
		RequestId: requestId,
		Data:      ClientConnectedResponse,
	}
	response, _ := json.Marshal(ClientConnectedResponseWrapper)
	streamMsg := string(response)
	client.IncomingMessagesCh <- streamMsg
}

func (client *Client) handleSendMessage(requestId util.UUID, requestMessageData *protocol.SendMessageRequest) {
	room, err := roomManager.getRoom(requestMessageData.RoomId)
	if err != nil {
		client.returnError(requestId, err)
	} else {
		//make sure client is a member in the room
		_, ok := room.Clients[client.Id]
		if ok {
			message := models.Message{
				Room:     room,
				Content:  requestMessageData.Message,
				SenderId: client.Id,
			}
			msg, _ := json.Marshal(message.Content)
			streamMsg := string(msg)
			room.MessagesCh <- streamMsg
			log.Printf("Send Message cmd received:\n%s", requestMessageData.String())
			client.returnAck(requestId)
		} else {
			client.returnForbiddenError(requestId)
		}
	}
}

func (client *Client) handleCreateRoom(requestId util.UUID, createRoomData *protocol.CreateRoomRequest) {
	room := roomManager.createRoom(createRoomData.Topic)
	log.Printf("Create Room cmd received:\n%s", createRoomData.String())
	createRoomResponse := protocol.CreateRoomResponse{
		RoomId: room.Id,
	}
	room.JoinClient(client.Id, client.IncomingMessagesCh)
	bits, err := json.Marshal(createRoomResponse)
	if err != nil {
		client.returnSystemError(requestId, err)
	} else {
		client.IncomingMessagesCh <- string(bits)
	}
}

func (client *Client) handleSetUserName(requestId util.UUID, setUserNameData *protocol.SetUserNameRequest) {
	client.Name = setUserNameData.Name
	log.Printf("Set User Topic cmd received:\n%s", setUserNameData.String())
	client.returnAck(requestId)
}

func (client *Client) handleSetRoomTopic(requestId util.UUID, setRoomTopicData *protocol.SetRoomTopicRequest) {
	room, err := roomManager.getRoom(setRoomTopicData.RoomId)
	if err != nil {
		client.returnError(requestId, err)
	} else {
		room.Topic = setRoomTopicData.Topic
		log.Printf("Set Room Topic cmd received:\n%s", setRoomTopicData.String())
		client.returnAck(requestId)
	}
}

func (client *Client) handleJoinRoom(requestId util.UUID, joinRoomRequest *protocol.JoinRoomRequest) {
	room, err := roomManager.getRoom(joinRoomRequest.RoomId)
	if err != nil {
		client.returnError(requestId, err)
	} else {
		room.JoinClient(client.Id, client.IncomingMessagesCh)
		log.Printf("Join Room cmd received:\n%s", joinRoomRequest.String())
		client.returnAck(requestId)
	}
}
