package managers

import (
	"errors"
	"sync"

	"github.com/gorilla/websocket"

	"github.com/bloom-chat/internal/util"
)

var clientOnce sync.Once
var mutex = &sync.Mutex{}

type ClientManager struct {
	clients map[util.UUID]*Client
}

var clientManager *ClientManager

func NewClientManager() *ClientManager {
	clientOnce.Do(func() {
		clients := make(map[util.UUID]*Client)
		clientManager = &ClientManager{clients: clients}
	})
	return clientManager
}

func (manager *ClientManager) GetClient(id util.UUID) (*Client, error) {
	client, ok := manager.clients[id]
	if ok {
		return client, nil
	} else {
		return nil, errors.New("client not found")
	}
}

func (manager *ClientManager) AddClient(conn *websocket.Conn) *Client {
	client := &Client{
		Conn:               conn,
		Id:                 util.GenerateID(),
		IncomingMessagesCh: make(chan string),
		CloseCh:            make(chan bool),
	}
	mutex.Lock()
	defer mutex.Unlock()
	manager.clients[client.Id] = client
	return client
}

func (manager *ClientManager) RemoveClient(clientId util.UUID) {
	mutex.Lock()
	defer mutex.Unlock()
	delete(manager.clients, clientId)
}
