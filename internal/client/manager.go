package client

import (
	"errors"
	"sync"

	"github.com/gorilla/websocket"

	"github.com/bloom-chat/internal/util"
)

var mutex = &sync.Mutex{}

type Manager struct {
	clients map[util.UUID]*Client
}

func NewManager() *Manager {
	clients := make(map[util.UUID]*Client)
	return &Manager{clients: clients}
}

func (manager *Manager) GetClient(id util.UUID) (*Client, error) {
	client, ok := manager.clients[id]
	if ok {
		return client, nil
	} else {
		return nil, errors.New("client not found")
	}
}

func (manager *Manager) AddClient(conn *websocket.Conn) *Client {
	client := &Client{
		Conn:               conn,
		Id:                 util.GenerateID(),
		IncomingMessagesCh: make(chan string),
		RoomsChs:           make(map[util.UUID] chan<- string),
		CloseCh:            make(chan bool),
	}
	mutex.Lock()
	manager.clients[client.Id] = client
	mutex.Unlock()
	return client
}

func (manager *Manager) RemoveClient(clientId util.UUID) {
	mutex.Lock()
	delete(manager.clients, clientId)
	mutex.Unlock()
}

func (manager *Manager) JoinRoom(clientId util.UUID, roomId util.UUID, roomCh chan<- string) {
	mutex.Lock()
	client, _ := manager.GetClient(clientId)
	client.RoomsChs[roomId]=roomCh
	mutex.Unlock()
}



//TODO Remove
func (manager *Manager) SendToAllClients(msg string) {
	for _, client := range manager.clients {
		client.IncomingMessagesCh <- msg
	}
}
