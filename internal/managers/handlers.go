package managers

import (
	"encoding/json"
	"errors"
	"github.com/bloom-chat/internal/models"
	"github.com/bloom-chat/internal/protocol"
	"github.com/bloom-chat/internal/util"
	log "github.com/sirupsen/logrus"
)

func newSession(requestId util.UUID, client *Client, displayName string) *Session {
	session := sessionManager.NewSession(displayName)
	err := sessionManager.AddClientToSession(session.Id, client)
	if err != nil {
		client.returnSystemError(requestId, nil)
		log.Warnf("failed to associate a client to a session\n"+
			"session ID: %s, client ID: %s, request ID: %s", session.Id, client.Id, requestId)
	}
	return session
}

func (client *Client) handleInitialize(requestId util.UUID, initializeRequest *protocol.InitializeRequest) {
	var session *Session
	if initializeRequest.SessionID != "" {
		var err error
		session, err = sessionManager.GetSession(initializeRequest.SessionID)
		if err != nil {
			session = newSession(requestId, client, initializeRequest.DisplayName)
		}
		if client.Initialized {
			if session.Id != initializeRequest.SessionID {
				client.returnHandshakeError(requestId, errors.New("session already initialized"))
				return
			}
		}
	} else {
		session = newSession(requestId, client, initializeRequest.DisplayName)
	}
	if initializeRequest.DisplayName != "" {
		session.DisplayName = initializeRequest.DisplayName
	}
	clientConnectedResponse := &protocol.ClientConnectedResponse{
		SessionID:   session.Id,
		DisplayName: session.DisplayName,
	}
	clientConnectedResponseWrapper := protocol.Response{
		RequestId: requestId,
		Data:      clientConnectedResponse,
	}
	response, _ := json.Marshal(clientConnectedResponseWrapper)
	streamMsg := string(response)
	client.Initialized = true
	client.IncomingMessagesCh <- streamMsg
}

//func (client *Client) handleSendMessage(requestId util.UUID, requestMessageData *protocol.SendMessageRequest) {
//	room, err := roomManager.getRoom(requestMessageData.RoomId)
//	if err != nil {
//		client.returnError(requestId, err)
//	} else {
//		//make sure client is a member in the room
//		_, ok := room.Clients[client.Id]
//		if !ok {
//			client.returnForbiddenError(requestId)
//		} else {
//			message := messageManager.createMessage(
//				&requestMessageData.Message, room, client)
//			MessageResponseWrapper := &protocol.Response{
//				RequestId: requestId,
//				Data:      message,
//			}
//			msg, err := json.Marshal(MessageResponseWrapper)
//			if err != nil {
//				client.returnSystemError(requestId, err)
//			} else {
//				streamMsg := string(msg)
//				room.MessagesCh <- streamMsg
//				log.Printf("Send Message cmd received:\n%s", requestMessageData.String())
//				client.returnAck(requestId)
//			}
//		}
//	}
//}
//
//func (client *Client) handleCreateRoom(requestId util.UUID, createRoomData *protocol.CreateRoomRequest) {
//	room := roomManager.createRoom(createRoomData.Topic, client)
//	log.Printf("Create Room cmd received:\n%s", createRoomData.String())
//	createRoomResponse := protocol.CreateRoomResponse{
//		RoomId: room.Id,
//	}
//	room.JoinClient(client.Id, client.IncomingMessagesCh)
//	bits, err := json.Marshal(createRoomResponse)
//	if err != nil {
//		client.returnSystemError(requestId, err)
//	} else {
//		client.IncomingMessagesCh <- string(bits)
//	}
//}

//func (client *Client) handleSetUserName(requestId util.UUID, setUserNameData *protocol.SetUserNameRequest) {
//	client.Name = setUserNameData.Name
//	log.Printf("Set User Topic cmd received:\n%s", setUserNameData.String())
//	client.returnAck(requestId)
//}
//
//func (client *Client) handleSetRoomTopic(requestId util.UUID, setRoomTopicData *protocol.SetRoomTopicRequest) {
//	room, err := roomManager.getRoom(setRoomTopicData.RoomId)
//	if err != nil {
//		client.returnError(requestId, err)
//	} else {
//		room.Topic = setRoomTopicData.Topic
//		log.Printf("Set Room Topic cmd received:\n%s", setRoomTopicData.String())
//		client.returnAck(requestId)
//	}
//}
//

func (client *Client) handleJoinRoom(requestId util.UUID, joinRoomRequest *protocol.JoinRoomRequest) {
	//TODO check DB!!
	room, err := roomManager.getRoom(joinRoomRequest.Handle)
	if err != nil {
		//TODO Room not found, create a room
		topic := models.Topic{
			Text: util.GenerateDisplayName(),
		}
		room = roomManager.createRoom(joinRoomRequest.Handle, &topic)
	}
	//TODO Room already exists
	//room.JoinClient(&client.IncomingMessagesCh)
	log.Infof("Join Room cmd received:\n%s", joinRoomRequest.String())
	log.Infof("new room created: %d", room)
	//TODO Return
	client.returnAck(requestId)
}

//func (client *Client) handleListRooms(requestId util.UUID, listRoomRequest *protocol.ListRoomsRequest) {
//	var rooms []*protocol.Room
//	roomsIDs, _ := roomManager.listRoomsIDs()
//	for _, roomID := range roomsIDs {
//		var responseRoom *protocol.Room
//		room, err := roomManager.getRoom(roomID)
//		if err != nil {
//			client.returnSystemError(requestId, err)
//		}
//		responseRoom = &protocol.Room{
//			Id:    roomID,
//			Topic: room.Topic,
//		}
//		rooms = append(rooms, responseRoom)
//	}
//	listRoomsResponse := &protocol.ListRoomsResponse{Rooms: rooms}
//	listRoomsResponseWrapper := protocol.Response{
//		RequestId: requestId,
//		Data:      listRoomsResponse,
//	}
//	response, _ := json.Marshal(listRoomsResponseWrapper)
//	streamMsg := string(response)
//	client.IncomingMessagesCh <- streamMsg
//}
