package client

import "github.com/bloom-chat/internal/protocol"
import "github.com/mitchellh/mapstructure"

func (client *Client) handleRequestMessage(data map[string]interface{}) {
	requestMessageData := protocol.RequestMessageData{}
	err := mapstructure.Decode(data, &requestMessageData)
	if err != nil {
		client.returnParseError()
	} else {
		client.IncomingMessagesCh <- requestMessageData.To
		client.IncomingMessagesCh <- requestMessageData.Message
	}
}

func (client *Client) returnParseError() {
	client.IncomingMessagesCh <- "failed to parse data"
}