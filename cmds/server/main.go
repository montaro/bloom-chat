package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var address = flag.String("address", "localhost:8080", "HTTP service address")
var upgrader = websocket.Upgrader{}

func echoHandler(conn *websocket.Conn) {
	defer conn.Close()
	for {
		msgT, msg, err := conn.ReadMessage()
		if err!=nil {
			log.Println("read message error: ", err)
			continue
		}
		log.Println("Received message: ", string(msg))
		if err = conn.WriteMessage(msgT, msg); err != nil {
			log.Println("write message error: ", err)
			continue
		}
	}
}

func echo(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade to websocket failed",  err)
	}
	go echoHandler(conn)
}

func main() {
	flag.Parse()
	fmt.Println("Server starting...: ", *address)
	log.SetFlags(0)
	http.HandleFunc("/echo", echo)
	log.Fatal(http.ListenAndServe(*address, nil))
}