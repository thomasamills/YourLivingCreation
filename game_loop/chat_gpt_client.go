package game_loop

import (
	"fmt"
	humanize_protobuf "testserver/src/generated/humanize-protobuf"
)

type ChatGptClient interface {
	SendPrompt(
		responderName,
		askerName,
		prompt string,
		requestType humanize_protobuf.HumanizeRequest_RequestType,
		message string,
	) (*humanize_protobuf.HumanizeResponse, error)
}

type ChatGptClientImpl struct {
}

func (c ChatGptClientImpl) SendPrompt(
	responderName,
	askerName,
	prompt string,
	requestType humanize_protobuf.HumanizeRequest_RequestType,
	message string,
) (*humanize_protobuf.HumanizeResponse, error) {
	switch requestType {
	case humanize_protobuf.HumanizeRequest_REQUEST_TYPE_AUTONOMOUS:
		prompt = prompt + fmt.Sprintf("\n\n%s", responderName)
		break
	case humanize_protobuf.HumanizeRequest_REQUEST_TYPE_SEND_MESSAGE_TO_NPC:
		prompt = prompt + fmt.Sprintf("\n\n%s:%s\n%s:", askerName, message, responderName)
		break
	}
	return nil, nil
}

func NewChatGptClient() ChatGptClient {
	return &ChatGptClientImpl{}
}
