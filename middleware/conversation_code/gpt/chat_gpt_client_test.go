package gpt

import (
	"fmt"
	"testing"
	humanize_protobuf "testserver/src/generated/humanize-protobuf"
)

func TestChatGptClientImpl_Init(t *testing.T) {
	gptClient, err := NewChatGptClient()
	if err != nil {
		t.Fail()
		return
	}
	prompt, err := gptClient.SendPrompt(
		"Bob",
		"Tom",
		`You are Bob you are having a back and forth conversation. Bob is a lawyer.$

Tom previously said in the conversation: Hello can you tell me what we are doing here today?$
Bob previously responded in the conversation with: Hello Tom$

You are both situated next to a well full of dirty water.$

Bob has a need to get some water asap.$

Bob is a far right christian, he always quotes the bible when answering a question.$

Bob is feeling Very Happy.$`,
		humanize_protobuf.HumanizeRequest_REQUEST_TYPE_SEND_MESSAGE_TO_NPC, "How do you feel now, Bob?",
	)
	if err != nil {
		t.Fail()
		return
	}
	fmt.Println(prompt.Response)
}
