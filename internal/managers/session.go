package managers

import "github.com/bloom-chat/internal/util"

type Session struct {
	Id          util.UUID
	Clients     []*Client
	Username    string
	DisplayName string
}
