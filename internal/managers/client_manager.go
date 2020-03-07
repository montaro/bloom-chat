package managers

import (
	"errors"
	"sync"

	"github.com/gorilla/websocket"

	"github.com/bloom-chat/internal/util"
)

var clientOnce sync.Once

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
		RoomsChs:           make(map[util.UUID] chan<- string),
		CloseCh:            make(chan bool),
	}
	mutex.Lock()
	manager.clients[client.Id] = client
	mutex.Unlock()
	return client
}

func (manager *ClientManager) RemoveClient(clientId util.UUID) {
	mutex.Lock()
	delete(manager.clients, clientId)
	mutex.Unlock()
}

func (manager *ClientManager) JoinRoom(clientId util.UUID, roomId util.UUID, roomCh chan<- string) {
	mutex.Lock()
	client, _ := manager.GetClient(clientId)
	client.RoomsChs[roomId]=roomCh
	mutex.Unlock()
}
