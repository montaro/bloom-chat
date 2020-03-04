package client

import (
	"encoding/json"
	"fmt"
	"github.com/bloom-chat/internal/protocol"
	"github.com/gorilla/websocket"
	"log"
)

type Id string

type Client struct {
	Conn        *websocket.Conn
	Id          Id
	Messages    chan string
	CloseSignal chan bool
}

var msgType = 1

func (client *Client) Start() {
	go client.Read()
	go client.Write()
}

func (client *Client) Read() {
	defer client.Conn.Close()
	welcomeMsg := fmt.Sprintf("Client connected number: %d", client.Id)
	log.Println(welcomeMsg)
	if err := client.Conn.WriteMessage(msgType, []byte(welcomeMsg)); err != nil {
		log.Println("Write message error: ", err)
	}
	for {
		_, msg, err := client.Conn.ReadMessage()
		if err != nil {
			log.Println("Read message error: ", err)
			break
		}
		log.Printf("Received message: %s from Client: %d\n", string(msg), client.Id)
		request := protocol.Request{}
		err = json.Unmarshal(msg, &request)
		if err != nil {
			log.Printf("Unmarshal Failed for message: %s\n", msg)
		} else {
			//dispatch message to rooms manager
			roomChannel <- string(msg)
		}
	}
}

func (client *Client) Write() {
	defer client.Conn.Close()
	for {
		select {
		case msg := <-client.Messages:
			if err := client.Conn.WriteMessage(msgType, []byte(msg)); err != nil {
				log.Println("Write message error: ", err)
				break
			}
		case <-client.CloseSignal:
			break
		}
	}
}

func (client *Client) Process() (protocol.Request, error) {
	return protocol.Request{}, nil
}
