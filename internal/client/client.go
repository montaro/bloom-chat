package client

import (
	"fmt"
	"github.com/bloom-chat/internal/util"
	"log"

	"github.com/gorilla/websocket"

	"github.com/bloom-chat/internal/protocol"
)

type Client struct {
	Conn               *websocket.Conn
	Id                 util.UUID
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
		//TODO send to the mean room, now sends to the Holy Room
		for _, roomCh := range client.RoomsChs {
			roomCh <- message

		}
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

func (client *Client) Process() (protocol.Request, error) {
	return protocol.Request{}, nil
}
