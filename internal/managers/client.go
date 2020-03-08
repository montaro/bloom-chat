package managers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
	"github.com/mitchellh/mapstructure"

	"github.com/bloom-chat/internal/protocol"
	"github.com/bloom-chat/internal/util"
)

type Client struct {
	Conn               *websocket.Conn
	Id                 util.UUID
	name               string
	IncomingMessagesCh chan string
	CloseCh            chan bool
}

var msgType = 1

func (client *Client) Start() {
	go client.Read()
	go client.Write()
}

func (client *Client) Read() {
	defer client.Conn.Close()
	welcomeMsg := fmt.Sprintf("Client connected: %s", client.Id)
	log.Println(welcomeMsg)
	if err := client.Conn.WriteMessage(msgType, []byte(welcomeMsg)); err != nil {
		log.Println("Write welcome message error: ", err)
	}
	for {
		_, msg, err := client.Conn.ReadMessage()
		message := string(msg)
		if err != nil {
			log.Println("Read message error: ", err)
			break
		}
		log.Printf("Received message: %s from Client: %s\n", message, client.Id)
		go client.Process(msg)
	}
}

func (client *Client) Write() {
	defer client.Conn.Close()
	for {
		select {
		case msg := <-client.IncomingMessagesCh:
			if err := client.Conn.WriteMessage(msgType, []byte(msg)); err != nil {
				log.Println("Write message error: ", err)
				break
			}
		case <-client.CloseCh:
			break
		}
	}
}

func (client *Client) Process(message []byte) {
	request, err := client.ParseRequest(message)
	if err != nil {
		var rId util.UUID
		if request != nil {
			rId = request.RequestId
		}
		client.returnError(rId, err)
	} else {
		switch request.Op {
		//Send message to room
		case protocol.SendMessage:
			requestMessageData := &protocol.SendMessageRequest{}
			err := mapstructure.Decode(request.Data, requestMessageData)
			if err != nil {
				client.returnParseDataError(request.RequestId, err)
			} else {
				client.handleSendMessage(request.RequestId, requestMessageData)
			}
		//Create room
		case protocol.CreateRoom:
			createRoomData := &protocol.CreateRoomRequest{}
			err := mapstructure.Decode(request.Data, createRoomData)
			if err != nil {
				client.returnParseDataError(request.RequestId, err)
			} else {
				client.handleCreateRoom(request.RequestId, createRoomData)
			}
		//Set user name
		case protocol.SetUserName:
			setUserNameData := &protocol.SetUserNameRequest{}
			err := mapstructure.Decode(request.Data, setUserNameData)
			if err != nil {
				client.returnParseDataError(request.RequestId, err)
			} else {
				client.handleSetUserName(request.RequestId, setUserNameData)
			}
		//Set room topic
		case protocol.SetRoomTopic:
			setRoomTopicData := &protocol.SetRoomTopicRequest{}
			err := mapstructure.Decode(request.Data, setRoomTopicData)
			if err != nil {
				client.returnParseDataError(request.RequestId, err)
			} else {
				client.handleSetRoomTopic(request.RequestId, setRoomTopicData)
			}
		//Join room
		case protocol.JoinRoom:
			joinRoomRequest := &protocol.JoinRoomRequest{}
			err := mapstructure.Decode(request.Data, joinRoomRequest)
			if err != nil {
				client.returnParseDataError(request.RequestId, err)
			} else {
				client.handleJoinRoom(request.RequestId, joinRoomRequest)
			}
		default:
			client.IncomingMessagesCh <- "UNKNOWN CMD: " + string(request.Op)
		}
	}
}

func (client *Client) ParseRequest(message []byte) (*protocol.Request, error) {
	request := protocol.Request{}
	err := json.Unmarshal(message, &request)
	if err != nil {
		log.Printf("error parsing a message: %s from client: %s\n%s\n", string(message), client.Id, err)
		return nil, err
	}
	if request.Op == "" {
		return nil, errors.New("op field is required")
	}
	if request.RequestId == "" {
		return nil, errors.New("request_id field is required")
	}
	return &request, nil
}
