package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"sync/atomic"

	"github.com/gorilla/websocket"

	clientz "github.com/bloom-chat/internal/client"
	roomz "github.com/bloom-chat/internal/room"
)

var address = flag.String("address", "localhost:8080", "HTTP service address")

var upgrader = websocket.Upgrader{}
var clientsCount uint64

var clientsManager *clientz.Manager
var roomsManager *roomz.Manager

//TODO Remove
var HolyRoom *roomz.Room

func chat(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade to websocket failed", err)
	}
	atomic.AddUint64(&clientsCount, 1)
	client := clientsManager.AddClient(conn)
	conn.SetCloseHandler(func(code int, text string) error {
		log.Printf("Client: %s disconnected!\n", client.Id)
		clientsManager.RemoveClient(client.Id)
		client.CloseCh <- true
		return nil
	})
	//TODO remove adding to the Holy Room
	err = HolyRoom.JoinClient(client.Id)
	clientsManager.JoinRoom(client.Id, HolyRoom.Id, HolyRoom.MessagesCh)
	if err != nil {
		//TODO will be removed with Holy Room removal
		log.Printf("client: %s failed to join Holy room, closing connection...", client.Id)
		_ = conn.Close()
	}
	client.Start()
}

func main() {
	flag.Parse()
	fmt.Println("Server blooming...: ", *address)
	log.SetFlags(0)
	http.HandleFunc("/chat", chat)
	clientsManager = clientz.NewManager()
	roomsManager = roomz.NewManager()
	HolyRoom = roomsManager.CreateRoom(clientsManager, "HolyRoom")
	go HolyRoom.Broadcast()
	log.Fatal(http.ListenAndServe(*address, nil))
}
