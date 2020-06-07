package managers

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gorilla/websocket"
	"github.com/mitchellh/mapstructure"
	log "github.com/sirupsen/logrus"

	"github.com/bloom-chat/internal/protocol"
	"github.com/bloom-chat/internal/util"
)

type IP string
type UserAgent string

const (
	Web    UserAgent = "WEB"
	Mobile UserAgent = "Mobile"
	CLI    UserAgent = "CLI"
)

type Client struct {
	Conn               *websocket.Conn
	Id                 util.UUID
	UserAgent          UserAgent
	IP                 IP
	Name               string
	IncomingMessagesCh chan string
	CloseCh            chan bool
	Initialized        bool
}

var msgType = 1

func (client *Client) Start() {
	go client.Read()
	go client.Write()
}

func (client *Client) Read() {
	defer client.Conn.Close()
	welcomeMsg := fmt.Sprintf("Client connected: %s", client.Id)
	log.Info(welcomeMsg)
	for {
		_, msg, err := client.Conn.ReadMessage()
		message := string(msg)
		if err != nil {
			log.Info("Read message error: ", err)
			break
		}
		log.Infof("Received message: %s from Client: %s\n", message, client.Id)
		//TODO Should it run in a different Goroutine?
		client.Process(msg)
	}
}

func (client *Client) Write() {
	defer client.Conn.Close()
	for {
		select {
		case msg := <-client.IncomingMessagesCh:
			if err := client.Conn.WriteMessage(msgType, []byte(msg)); err != nil {
				log.Info("Write message error: ", err)
				break
			}
		case <-client.CloseCh:
			break
		}
	}
}

func (client *Client) Process(message []byte) {
	request, err := client.parseRequest(message)
	if err != nil {
		var rId util.UUID
		if request != nil {
			rId = request.RequestId
		}
		client.returnError(rId, err)
	} else {
		switch request.Op {
		//Initialize
		case protocol.Initialize:
			initializeRequest := &protocol.InitializeRequest{}
			err := mapstructure.Decode(request.Data, initializeRequest)
			if err != nil {
				client.returnParseDataError(request.RequestId, err)
			} else {
				client.handleInitialize(request.RequestId, initializeRequest)
				//TODO replace with FSM
				client.Initialized = true
			}
		////Send message to room
		//case protocol.SendMessage:
		//	client.assertInitialized(request.RequestId)
		//	requestMessageData := &protocol.SendMessageRequest{}
		//	err := mapstructure.Decode(request.Data, requestMessageData)
		//	if err != nil {
		//		client.returnParseDataError(request.RequestId, err)
		//	} else {
		//		client.handleSendMessage(request.RequestId, requestMessageData)
		//	}
		////Create room
		//case protocol.CreateRoom:
		//	client.assertInitialized(request.RequestId)
		//	createRoomData := &protocol.CreateRoomRequest{}
		//	err := mapstructure.Decode(request.Data, createRoomData)
		//	if err != nil {
		//		client.returnParseDataError(request.RequestId, err)
		//	} else {
		//		client.handleCreateRoom(request.RequestId, createRoomData)
		//	}
		////Set user name
		//case protocol.SetUserName:
		//	client.assertInitialized(request.RequestId)
		//	client.assertInitialized(request.RequestId)
		//	setUserNameData := &protocol.SetUserNameRequest{}
		//	err := mapstructure.Decode(request.Data, setUserNameData)
		//	if err != nil {
		//		client.returnParseDataError(request.RequestId, err)
		//	} else {
		//		client.handleSetUserName(request.RequestId, setUserNameData)
		//	}
		////Set room topic
		//case protocol.SetRoomTopic:
		//	client.assertInitialized(request.RequestId)
		//	client.assertInitialized(request.RequestId)
		//	setRoomTopicData := &protocol.SetRoomTopicRequest{}
		//	err := mapstructure.Decode(request.Data, setRoomTopicData)
		//	if err != nil {
		//		client.returnParseDataError(request.RequestId, err)
		//	} else {
		//		client.handleSetRoomTopic(request.RequestId, setRoomTopicData)
		//	}
		////Join room
		//case protocol.JoinRoom:
		//	client.assertInitialized(request.RequestId)
		//	client.assertInitialized(request.RequestId)
		//	joinRoomRequest := &protocol.JoinRoomRequest{}
		//	err := mapstructure.Decode(request.Data, joinRoomRequest)
		//	if err != nil {
		//		client.returnParseDataError(request.RequestId, err)
		//	} else {
		//		client.handleJoinRoom(request.RequestId, joinRoomRequest)
		//	}
		////List rooms
		//case protocol.ListRooms:
		//	client.assertInitialized(request.RequestId)
		//	client.assertInitialized(request.RequestId)
		//	listRoomsRequest := &protocol.ListRoomsRequest{}
		//	err := mapstructure.Decode(request.Data, listRoomsRequest)
		//	if err != nil {
		//		client.returnParseDataError(request.RequestId, err)
		//	} else {
		//		client.handleListRooms(request.RequestId, listRoomsRequest)
		//	}
		default:
			client.returnUnexpectedCMDError(request.RequestId, string(request.Op))
		}
	}
}

func (client *Client) parseRequest(message []byte) (*protocol.Request, error) {
	request := protocol.Request{}
	err := json.Unmarshal(message, &request)
	if err != nil {
		log.Infof("error parsing a message: %s from client: %s\n%s\n", string(message), client.Id, err)
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

func (client *Client) returnHandshake(requestId util.UUID) {
	handshake := &protocol.Handshake{ProtocolVersion: protocol.ProtocolVersion}
	response := protocol.Response{
		RequestId: requestId,
		Data:      handshake,
	}
	msg, _ := json.Marshal(response)
	streamMsg := string(msg)
	client.IncomingMessagesCh <- streamMsg
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
	client.returnError(requestId, errors.New(fmt.Sprintf("failed to parse data with error: %s",
		err.Error())))
}

func (client *Client) returnSystemError(requestId util.UUID, err error) {
	client.returnError(requestId, errors.New(fmt.Sprintf("internal system error: %s",
		err.Error())))
}

func (client *Client) returnHandshakeError(requestId util.UUID, err error) {
	client.returnError(requestId, errors.New(fmt.Sprintf("handshake error: %s",
		err.Error())))
}

func (client *Client) returnForbiddenError(requestId util.UUID) {
	client.returnError(requestId, errors.New("forbidden action"))
}

func (client *Client) returnUnexpectedCMDError(requestId util.UUID, op string) {
	client.returnError(requestId, errors.New(fmt.Sprintf("unexpected op: %s", op)))
}

func (client *Client) assertInitialized(requestId util.UUID) {
	if !client.Initialized {
		client.returnHandshakeError(requestId, errors.New("expecting initialize message"))
	}
}
