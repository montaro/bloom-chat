package managers

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"log"

	"github.com/gorilla/websocket"

	"github.com/bloom-chat/internal/protocol"
	"github.com/bloom-chat/internal/util"
)

type Client struct {
	Conn               *websocket.Conn
	Id                 util.UUID
	name               string
	IncomingMessagesCh chan string
	RoomsChs           map[util.UUID]chan<- string
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
		response := protocol.Response{
			Op:   protocol.ERROR,
			Data: map[string]string{"value": err.Error()},
		}
		msg, _ := json.Marshal(response)
		client.IncomingMessagesCh <- string(msg)
	} else {
		switch request.Op {
		case protocol.RequestMessage:
			requestMessageData := &protocol.RequestMessageData{}
			err := mapstructure.Decode(request.Data, requestMessageData)
			if err != nil {
				client.returnParseError()
			} else {
				client.handleRequestMessage(requestMessageData)
			}
		case protocol.CreateRoom:
			createRoomData := &protocol.CreateRoomData{}
			err := mapstructure.Decode(request.Data, createRoomData)
			if err != nil {
				client.returnParseError()
			} else {
				client.handleCreateRoom(createRoomData)
			}
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
	return &request, nil
}
