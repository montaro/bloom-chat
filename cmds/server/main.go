package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
	"sync/atomic"
)

var address = flag.String("address", "localhost:8080", "HTTP service address")

var upgrader = websocket.Upgrader{}
var clientsCount uint64
var roomChannel = make(chan string)
var msgType = 1
var mutex = &sync.Mutex{}

type Client struct {
	id          uint64
	messages    chan string
	closeSignal chan bool
}

var clients = make(map[uint64]Client)

func clientReader(conn *websocket.Conn, client *Client) {
	defer conn.Close()
	welcomeMsg := fmt.Sprintf("Client connected number: %d", client.id)
	log.Println(welcomeMsg)
	if err := conn.WriteMessage(msgType, []byte(welcomeMsg)); err != nil {
		log.Println("write message error: ", err)
	}
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("read message error: ", err)
			break
		}
		log.Printf("Received message: %s from Client: %d\n", string(msg), client.id)
		roomChannel <- string(msg)
	}
}

func clientWriter(conn *websocket.Conn, client *Client) {
	defer conn.Close()
	for {
		select {
		case msg := <-client.messages:
			if err := conn.WriteMessage(msgType, []byte(msg)); err != nil {
				log.Println("write message error: ", err)
				break
			}
		case <-client.closeSignal:
			break
		}
	}
}

func roomDispatch() {
	for {
		select {
		case msg := <-roomChannel:
			for _, client := range clients {
				client.messages <- msg
			}
		}
	}
}

func chat(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade to websocket failed", err)
	}
	atomic.AddUint64(&clientsCount, 1)
	client := Client{
		id:          clientsCount,
		messages:    make(chan string),
		closeSignal: make(chan bool),
	}
	mutex.Lock()
	clients[clientsCount] = client
	mutex.Unlock()
	conn.SetCloseHandler(func(code int, text string) error {
		log.Printf("Client: %d disconnected!\n", client.id)
		mutex.Lock()
		delete(clients, client.id)
		mutex.Unlock()
		client.closeSignal <- true
		return nil
	})
	go clientReader(conn, &client)
	go clientWriter(conn, &client)
}

func main() {
	flag.Parse()
	fmt.Println("Server starting...: ", *address)
	log.SetFlags(0)
	http.HandleFunc("/chat", chat)
	go roomDispatch()
	log.Fatal(http.ListenAndServe(*address, nil))
}
