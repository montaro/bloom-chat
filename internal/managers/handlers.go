package managers

import (
	"fmt"
	"github.com/bloom-chat/internal/protocol"
	"github.com/mitchellh/mapstructure"
)

func (client *Client) handleRequestMessage(data map[string]interface{}) {
	requestMessageData := protocol.RequestMessageData{}
	err := mapstructure.Decode(data, &requestMessageData)
	if err != nil {
		client.returnParseError()
	} else {
		r := requestMessageData.To
		room, ok := client.RoomsChs[r]
		if !ok {
			client.IncomingMessagesCh <- fmt.Sprintf("unknown or closed room: %s", r)
		} else {
			room <- requestMessageData.Message
		}
		//TODO Remove
		client.IncomingMessagesCh <- string(requestMessageData.To)
		client.IncomingMessagesCh <- requestMessageData.Message
	}
}

func (client *Client) handleCreateRoom(data map[string]interface{}) {
	createRoomData := protocol.CreateRoomData{}
	err := mapstructure.Decode(data, &createRoomData)
	if err != nil {
		client.returnParseError()
	} else {
		room := roomManager.CreateRoom(createRoomData.Topic)
		client.IncomingMessagesCh <- string(room.id)
		client.IncomingMessagesCh <- room.topic
	}
}

func (client *Client) returnParseError() {
	client.IncomingMessagesCh <- "failed to parse data"
}
