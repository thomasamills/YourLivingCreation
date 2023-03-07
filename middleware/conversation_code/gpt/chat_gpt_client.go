package gpt

import (
	"errors"
	"fmt"
	gpt35 "github.com/AlmazDelDiablo/gpt3-5-turbo-go"
	"log"
	humanize_protobuf "testserver/src/generated/humanize-protobuf"
)

const (
	OpenAIKey = "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"
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
	logger      log.Logger
	gpt35Client *gpt35.Client
}

func (c *ChatGptClientImpl) Init() error {
	key := OpenAIKey
	if len(key) <= 0 {
		return errors.New("no open ai key")
	}
	c.gpt35Client = gpt35.NewClient(key)

	return nil
}

func (c *ChatGptClientImpl) SendPrompt(
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
	req := &gpt35.Request{
		Model: gpt35.ModelGpt35Turbo,
		Messages: []*gpt35.Message{
			{
				Role:    gpt35.RoleUser,
				Content: prompt,
			},
		},
	}
	resp, err := c.gpt35Client.GetChat(req)
	if err != nil {
		return nil, nil
	}
	return &humanize_protobuf.HumanizeResponse{
		Message:        message,
		Response:       resp.Choices[0].Message.Content,
		NpcTextEmotion: "neutral",
	}, nil
}

func NewChatGptClient() (ChatGptClient, error) {
	client := &ChatGptClientImpl{}
	err := client.Init()
	if err != nil {
		return nil, err
	}
	return client, nil
}
