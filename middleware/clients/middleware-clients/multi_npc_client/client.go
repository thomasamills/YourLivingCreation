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
	"time"
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

	connectionResponse, err := client.Connect(context.Background(), &humanize_protobuf.MiddleWareConnectReq{
		SpeakerName: speakerName,
		NpcInformation: []*humanize_protobuf.NpcRequestInformation{
			{
				Name:                         "Bob",
				PresetEmotionalStateId:       "DEFAULT_EMOTIONAL_STATE",
				DefaultPromptId:              "FAR_RIGHT_PUB_MAN_PROMPT",
				GenConfigId:                  "DEFAULT_CONFIG",
				PersonalityIds:               []string{"DEFAULT_PERSONALITY"},
				PromptSetId:                  []string{"FAR_RIGHT_PUB_MAN_PROMPT_SET"},
				NeedsPromptSegmentSetIds:     []string{"BASIC_HUMAN_NEEDS_PROMPT_SEGMENT_SET"},
				ReligionSegmentSetIds:        []string{"CHRISTIANITY_PROMPT_SEGMENT_SET_ID"},
				IdeologySegmentSetIds:        []string{"CAPITALIST_IDEOLOGY_PROMPT_SEGMENT_SET"},
				PersonalityTypeSegmentSetIds: []string{"ISTJ_PERSONALITY_TYPE_PROMPT_SEGMENT_SET"},
				EmotionalPrimerIds:           []string{"DEFAULT_EMOTIONAL_PRIMER_SET_ID"},
				ActuationRuleSetIds:          []string{"TBD"},
			},
			{
				Name:                         "Alex Jones",
				PresetEmotionalStateId:       "DEFAULT_EMOTIONAL_STATE",
				DefaultPromptId:              "YOUNG_FAR_LEFT_SOBER_MAN_PROMPT",
				GenConfigId:                  "DEFAULT_CONFIG",
				PersonalityIds:               []string{"DEFAULT_PERSONALITY"},
				PromptSetId:                  []string{"YOUNG_FAR_LEFT_SOBER_MAN_PROMPT_SET"},
				NeedsPromptSegmentSetIds:     []string{"BASIC_HUMAN_NEEDS_PROMPT_SEGMENT_SET"},
				IdeologySegmentSetIds:        []string{"SOCIALIST_IDEOLOGY_PROMPT_SEGMENT_SET"},
				PersonalityTypeSegmentSetIds: []string{"ISTJ_PERSONALITY_TYPE_PROMPT_SEGMENT_SET"},
				EmotionalPrimerIds:           []string{"DEFAULT_EMOTIONAL_PRIMER_SET_ID"},
				ActuationRuleSetIds:          []string{"TBD"},
			},
		},
		StartAsyncGameLoop:                      true,
		WaitForCommitMessageBeforeUpdatingState: false,
	})

	if err != nil {
		logrus.Fatal(err)
	}

	npc1Name := "Bob"
	npc2Name := "Alan"

	npc1Id := connectionResponse.NpcNameToId[npc1Name]
	npc2Id := connectionResponse.NpcNameToId[npc2Name]

	sessionId := connectionResponse.SessionId

	// listen to async responses
	go func() {
		for {
			if client != nil {
				response, err := client.GetConversationInformation(context.Background(),
					&humanize_protobuf.GetConversationInformationRequest{
						SessionId: sessionId,
					})
				if err != nil {
				} else {
					for _, value := range response.MessageResponseData {
						if len(value.Response) == 0 && len(value.Message) > 0 {
							// Autonomous no target
							fmt.Println(fmt.Sprintf("%s has auto interjected with:"+
								" \n%s: %s", value.AskerName, value.AskerName, value.Message))
						}
						if len(value.Response) > 0 && len(value.Message) > 0 {
							fmt.Println(fmt.Sprintf("%s has asked %s:", value.AskerName, value.ResponderName))
							fmt.Println(fmt.Sprintf("%s: %s", value.AskerName, value.Message))
							fmt.Println(fmt.Sprintf("%s: %s", value.ResponderName, value.Response))
						}
					}

					if err != nil {
						logrus.Fatal(err)
					}
					time.Sleep(time.Millisecond * 500)
				}
			}
		}
	}()

	for {
		fmt.Println("Options: type 'l' to speak to the loyalist or 'e' to speak to the elder or 'p' to speak to the pagan")
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		inputText := input.Text()

		npcId := ""
		npcName := ""

		if inputText == "b" {
			npcId = npc1Id
			npcName = npc1Name
		} else if inputText == "a" {
			npcId = npc2Id
			npcName = npc2Name
		}

		if len(npcId) > 0 {
			fmt.Println(fmt.Sprintf("You are now speaking to %s, please type a message", npcName))
			input := bufio.NewScanner(os.Stdin)
			input.Scan()
			inputText := input.Text()
			fmt.Println(fmt.Sprintf("%s: %s", speakerName, inputText))

			response, err := client.SendText(context.Background(), &humanize_protobuf.SendTextRequest{
				SessionId:        sessionId,
				DirectedEntityId: npcId,
				Message:          inputText,
				IsAction:         false,
			})

			if err != nil {
				logrus.Fatal(response)
			}

			if err != nil {
				logrus.Fatal(err)
			}

		}

	}
}
