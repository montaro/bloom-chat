package client

import (
	"errors"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"sync"
)

var mutex = &sync.Mutex{}

type Manager struct {
	clients map[Id]*Client
}

func NewManager() *Manager {
	clients := make(map[Id]*Client)
	return &Manager{clients:clients}
}

func (manager *Manager) GetClient(id Id) (*Client, error) {
	client, ok := manager.clients[id]
	if ok {
		return client, nil
	} else {
		return nil, errors.New("client not found")
	}
}

func (manager *Manager) AddClient(conn *websocket.Conn) *Client {
	client := &Client{
		Conn:        conn,
		Id:          manager.GenerateID(),
		Messages:    make(chan string),
		CloseSignal: make(chan bool),
	}
	mutex.Lock()
	manager.clients[client.Id] = client
	mutex.Unlock()
	return client
}

func (manager *Manager) RemoveClient(clientId Id) {
	mutex.Lock()
	delete(manager.clients, clientId)
	mutex.Unlock()
}

//TODO remove
func (manager *Manager) GetAllClients() *map[Id]*Client {
	return &manager.clients
}

func (manager *Manager) GenerateID() Id {
	return Id(uuid.New().String())
}
