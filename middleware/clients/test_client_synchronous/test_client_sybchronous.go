package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
	humanize_protobuf "testserver/src/generated/humanize-protobuf"
)

const (
	MiddleWarePort = "MIDDLEWARE_PORT"
)

func main() {
	speakerName := "Tom"
	// Connect to grpc
	address := "localhost:" + os.Getenv(MiddleWarePort)
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logrus.Fatal(err)
	}
	client := humanize_protobuf.NewConnectClient(conn)

	npc1Name := "Sophia"

	connectionResponse, err := client.Connect(context.Background(), &humanize_protobuf.MiddleWareConnectReq{
		SpeakerName: speakerName,
		NpcInformation: []*humanize_protobuf.NpcRequestInformation{
			{
				Name:                   npc1Name,
				PresetEmotionalStateId: "DEFAULT_EMOTIONAL_STATE",
				DefaultPromptId:        "SOFIA_PROMPT",
				PersonalityId:          "DEFAULT_PERSONALITY",
				PromptSetIds:           "SOFIA_PROMPT_SET",
				PromptSegmentSetId:     "SOFIA_PROMPT_SEGMENT_SET",
				GenConfigId:            "DEFAULT_CONFIG",
				ActuationRuleSetId:     "SOFIA_FAKE",
			},
		},
		StartAsyncGameLoop:                      false,
		WaitForCommitMessageBeforeUpdatingState: true,
	})

	if err != nil {
		logrus.Fatal(err)
	}

	npc1Id := connectionResponse.NpcNameToId[npc1Name]
	sessionId := connectionResponse.SessionId

	for {

		fmt.Println(fmt.Sprintf("You are now speaking to %s, please type a message", "Sophia"))
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		inputText := input.Text()
		fmt.Println(fmt.Sprintf("%s: %s", speakerName, inputText))

		response, err := client.SendText(context.Background(), &humanize_protobuf.SendTextRequest{
			SessionId:        sessionId,
			DirectedEntityId: npc1Id,
			Message:          inputText,
			IsAction:         false,
		})

		for _, value := range response.MessageResponseData {
			fmt.Println(fmt.Sprintf("%s:%s", "Sophia", value.Response))
			_, err := client.Commit(context.Background(), &humanize_protobuf.CommitRequest{
				SessionId: sessionId,
				MessageId: value.CommitToken,
			})
			if err != nil {
				logrus.Fatal(err)
			}
			// Commit message
		}
		if err != nil {
			logrus.Fatal(err)
		}

	}

}
