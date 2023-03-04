package game_loop

import humanize_protobuf "testserver/src/generated/humanize-protobuf"

type ChatGptClient interface {
	SendPrompt(
		responderName,
		askerName,
		prompt string,
		requestType humanize_protobuf.HumanizeRequest_RequestType,
	) (*humanize_protobuf.HumanizeResponse, error)
}

type ChatGptClientImpl struct {
}

func (c ChatGptClientImpl) SendPrompt(
	responderName,
	askerName,
	prompt string,
	requestType humanize_protobuf.HumanizeRequest_RequestType,
) (*humanize_protobuf.HumanizeResponse, error) {
	switch requestType {
	case humanize_protobuf.HumanizeRequest_REQUEST_TYPE_AUTONOMOUS:
		prompt = prompt
		break
	case humanize_protobuf.HumanizeRequest_REQUEST_TYPE_SEND_MESSAGE_TO_NPC:
		break
	}
	return nil, nil
}

func NewChatGptClient() ChatGptClient {
	return &ChatGptClientImpl{}
}
