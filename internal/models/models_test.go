package models

import (
	"github.com/bloom-chat/internal/util"
	"testing"
)

func TestBroadcast(t *testing.T) {
	clients := make(map[util.UUID]chan string)

	clientId1 := util.GenerateID()
	clientCh1 := make(chan string)
	clients[clientId1] = clientCh1
	messagesCh1 := make(chan string)
	testMsg1 := "test-1"
	testMsg2 := "test-2"
	room1 := Room{
		Topic:      "test1",
		Clients:    clients,
		MessagesCh: messagesCh1,
	}
	go room1.Broadcast()
	room1.MessagesCh <- testMsg1
	msg := <-clientCh1
	if msg != testMsg1 {
		t.Errorf("Unexpected message, wanted: %s got:%s", testMsg1, msg)
	}
	room1.MessagesCh <- testMsg2
	msg = <-clientCh1
	if msg != testMsg2 {
		t.Errorf("Unexpected message, wanted: %s got:%s", testMsg2, msg)
	}
}

func TestJoinClient(t *testing.T) {
	clients := make(map[util.UUID]chan string)

	clientId1 := util.GenerateID()
	clientId2 := util.GenerateID()
	clientCh1 := make(chan string)
	clientCh2 := make(chan string)
	messagesCh1 := make(chan string)
	room1 := Room{
		Topic:      "test1",
		Clients:    clients,
		MessagesCh: messagesCh1,
	}
	room1.JoinClient(clientId1, clientCh1)
	room1.JoinClient(clientId2, clientCh2)
	if room1.Clients[clientId1] != clientCh1 {
		t.Errorf("Client: %s with channel: %v didn't join the room", clientId1, clientCh1)
	}
	if room1.Clients[clientId2] != clientCh2 {
		t.Errorf("Client: %s with channel: %v didn't join the room", clientId2, clientCh2)
	}
}