package managers

import (
	//"errors"

	"github.com/bloom-chat/internal/models"
)

type MessageManager struct {
}

var messageManager *MessageManager

func (messageManager *MessageManager) createMessage(message *models.Message, room *models.Room,
	client *Client) *models.Message {
	message = &models.Message{
		Room:             room,
		Kind:             message.Kind,
		Content:          message.Content,
		FormattedContent: message.FormattedContent,
		Sender: &models.User{
			ClientID: client.Id,
			Name:     client.Name,
			Status:   models.Online,
			Client:   models.ClientWeb,
		},
	}
	message = message.SaveMessage()
	return message
}
