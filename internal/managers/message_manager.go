package managers

type MessageManager struct {
}

var messageManager *MessageManager

//func (messageManager *MessageManager) createMessage(message *models.Message, room *models.Room,
//	client *Client) *models.Message {
//	message = &models.Message{
//		Room:             room,
//		Kind:             message.Kind,
//		Content:          message.Content,
//		FormattedContent: message.FormattedContent,
//		Sender: &models.User{
//			ClientID: client.Id,
//			Name:     client.Name,
//			Status:   models.Online,
//			Client:   models.Web,
//		},
//	}
//	message = message.SaveMessage()
//	return message
//}
