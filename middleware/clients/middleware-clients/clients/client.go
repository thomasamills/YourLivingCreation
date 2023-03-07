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
			{},
		},
		StartAsyncGameLoop:                      true,
		WaitForCommitMessageBeforeUpdatingState: false,
	})

	if err != nil {
		logrus.Fatal(err)
	}

	npc1Id := connectionResponse.NpcNameToId[npc1Name]
	npc2Id := connectionResponse.NpcNameToId[npc2Name]
	npc3Id := connectionResponse.NpcNameToId[npc3Name]

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

		if inputText == "l" {
			npcId = npc1Id
			npcName = npc1Name
		} else if inputText == "e" {
			npcId = npc2Id
			npcName = npc2Name
		} else if inputText == "p" {
			npcId = npc3Id
			npcName = npc3Name
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