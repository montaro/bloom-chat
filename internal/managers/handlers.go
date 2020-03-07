package managers

import (
	"fmt"
	"log"

	"github.com/bloom-chat/internal/protocol"
)

func (client *Client) handleRequestMessage(requestMessageData *protocol.RequestMessageData) {
	r := requestMessageData.To
	room, ok := client.RoomsChs[r]
	if !ok {
		client.IncomingMessagesCh <- fmt.Sprintf("unknown or closed room: %s", r)
	} else {
		room <- requestMessageData.Message
	}
	log.Printf("Request Message cmd received:\n%s", requestMessageData.String())
	//TODO Remove, return response instead
	client.IncomingMessagesCh <- string(requestMessageData.To)
	client.IncomingMessagesCh <- requestMessageData.Message

}

func (client *Client) handleCreateRoom(createRoomData *protocol.CreateRoomData) {
	room := roomManager.CreateRoom(createRoomData.Topic)
	log.Printf("Create Room cmd received:\n%s", createRoomData.String())
	//TODO Remove, return response instead
	client.IncomingMessagesCh <- string(room.id)
	client.IncomingMessagesCh <- room.topic
}

func (client *Client) returnParseError() {
	client.IncomingMessagesCh <- "failed to parse data"
}
