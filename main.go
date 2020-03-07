package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"

	"github.com/bloom-chat/internal/managers"
)

var address = flag.String("address", "0.0.0.0:8080", "HTTP service address")

var upgrader = websocket.Upgrader{}

var clientsManager *managers.ClientManager

func chat(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade to websocket failed", err)
	}
	client := clientsManager.AddClient(conn)
	conn.SetCloseHandler(func(code int, text string) error {
		log.Printf("Client: %s disconnected!\n", client.Id)
		clientsManager.RemoveClient(client.Id)
		client.CloseCh <- true
		return nil
	})
	client.Start()
}

func main() {
	flag.Parse()
	fmt.Println("Server blooming...: ", *address)
	log.SetFlags(0)
	http.HandleFunc("/_ah/health", healthCheckHandler)
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.HandleFunc("/ws", chat)
	//initialize managers
	clientsManager = managers.NewClientManager()
	managers.NewRoomManager()
	log.Fatal(http.ListenAndServe(*address, nil))
}

// healthCheckHandler is used by App Engine Flex to check instance health.
func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprint(w, "ok")
}
