package managers

import (
	"errors"
	"github.com/bloom-chat/internal/util"
	"sync"
)

var sessionOnce sync.Once

type SessionManager struct {
	sessions map[util.UUID]*Session
}

var sessionManager *SessionManager

func NewSessionManager() *SessionManager {
	sessionOnce.Do(func() {
		sessions := make(map[util.UUID]*Session)
		sessionManager = &SessionManager{sessions: sessions}
	})
	return sessionManager
}

func (manager *SessionManager) GetSession(id util.UUID) (*Session, error) {
	session, ok := manager.sessions[id]
	if ok {
		return session, nil
	} else {
		return nil, errors.New("session not found")
	}
}

func (manager *SessionManager) NewSession(displayName string) *Session {
	session := &Session{
		Id:       util.GenerateID(),
		DisplayName: displayName,
	}
	mutex.Lock()
	defer mutex.Unlock()
	manager.sessions[session.Id] = session
	return session
}

func (manager *SessionManager) AddClientToSession(id util.UUID, client *Client) error {
	session, err := manager.GetSession(id)
	if err != nil {
		return err
	} else {
		session.Clients = append(session.Clients, client)
	}
	return nil
}

func (manager *SessionManager) RemoveSession(sessionId util.UUID) {
	mutex.Lock()
	defer mutex.Unlock()
	delete(manager.sessions, sessionId)
}
