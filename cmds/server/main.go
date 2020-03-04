package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"sync/atomic"

	"github.com/gorilla/websocket"

	"github.com/bloom-chat/internal/client"
)

var address = flag.String("address", "localhost:8080", "HTTP service address")

var upgrader = websocket.Upgrader{}
var clientsCount uint64
var roomChannel = make(chan string)


var clientManager *client.Manager

func roomDispatch() {
	for {
		select {
		case msg := <-roomChannel:
			for _, client := range *clientManager.GetAllClients() {
				client.Messages <- msg
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
	client := clientManager.AddClient(conn)
	conn.SetCloseHandler(func(code int, text string) error {
		log.Printf("Client: %d disconnected!\n", client.Id)
		clientManager.RemoveClient(client.Id)
		client.CloseSignal <- true
		return nil
	})
	client.Start()
}

func main() {
	flag.Parse()
	fmt.Println("Server starting...: ", *address)
	log.SetFlags(0)
	http.HandleFunc("/chat", chat)
	clientManager = client.NewManager()
	go roomDispatch()
	log.Fatal(http.ListenAndServe(*address, nil))
}
