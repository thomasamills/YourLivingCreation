package main

// #include <stdio.h>
// #include <stdlib.h>
import "C"
import (
	"context"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
	humanize_protobuf "testserver/src/generated/humanize-protobuf"
)

const (
	MiddleWarePort  = "MIDDLEWARE_PORT"
	HumanizeAddress = "HUMANIZE_ADDRESS"
)

func CreateClient() humanize_protobuf.ConnectClient {
	address := "localhost:" + os.Getenv(MiddleWarePort)
	log.Println("address: " + address)
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logrus.Fatal(err)
	}
	client := humanize_protobuf.NewConnectClient(conn)
	return client
}

type StartSessionResponse struct {
	SessionId string
	SophiaId  string
}

type SendMessageRequest struct {
	MessageInput string
	SessionId    string
	SophiaId     string
	ExtraData    string
}

type SendMessageResponse struct {
	Response    string
	Confidence  float32
	CommitToken string
}

type SendCommitRequest struct {
	SessionId   string
	CommitToken string
}

type SendCommitResponse struct {
	EmotionalState EmotionalState
}

type EmotionalState struct {
	SadHappy           int
	AgitatedPatient    int
	BurnedOutAmbitious int
	HateLove           int
}

//export startSession
func startSession(speakerName *C.char) *C.char {
	client := CreateClient()
	connectionResponse, err := client.Connect(context.Background(), &humanize_protobuf.MiddleWareConnectReq{
		SpeakerName: C.GoString(speakerName),
		NpcInformation: []*humanize_protobuf.NpcRequestInformation{
			{
				Name:                   "Bob",
				PresetEmotionalStateId: "DEFAULT_EMOTIONAL_STATE",
				DefaultPromptId:        "FAR_RIGHT_PUB_MAN_PROMPT",
				PersonalityId:          "DEFAULT_PERSONALITY",
				PromptSetId:            "FAR_RIGHT_PUB_MAN_PROMPT_SET",
				PromptSegmentSetId:     "BASIC_HUMAN_NEEDS_PROMPT_SEGMENT_SET",
				GenConfigId:            "DEFAULT_CONFIG",
				ActuationRuleSetId:     "",
			},
		},
		StartAsyncGameLoop:                      false,
		WaitForCommitMessageBeforeUpdatingState: true,
	})
	if err != nil {
		log.Println("Error connecting to the grpc server " + err.Error())
		return C.CString("")
	}
	sessionResponse := &StartSessionResponse{
		SessionId: connectionResponse.SessionId,
		SophiaId:  connectionResponse.NpcNameToId["Bob"],
	}
	marshalledString, err := json.Marshal(sessionResponse)
	return C.CString(string(marshalledString))
}

//export sendMessage
func sendMessage(sendMessageInfo *C.char) *C.char {
	client := CreateClient()
	var sendMessageRequest SendMessageRequest
	err := json.Unmarshal([]byte(C.GoString(sendMessageInfo)), &sendMessageRequest)
	response, err := client.SendText(context.Background(), &humanize_protobuf.SendTextRequest{
		SessionId:        sendMessageRequest.SessionId,
		Message:          sendMessageRequest.MessageInput,
		IsAction:         false,
		DirectedEntityId: sendMessageRequest.SophiaId,
		AdditionalJson:   sendMessageRequest.ExtraData,
	})
	if err != nil {
		return C.CString("")
	}
	responseText := ""
	commitToken := ""
	if len(response.MessageResponseData) >= 0 {
		for _, message := range response.MessageResponseData {
			responseText = message.Response
			commitToken = message.CommitToken
			break
		}
	}
	responseData := &SendMessageResponse{
		Response:    responseText,
		Confidence:  1,
		CommitToken: commitToken,
	}
	marshalledString, err := json.Marshal(responseData)
	if err != nil {
		return C.CString("")
	}
	return C.CString(string(marshalledString))

}

//export commit
func commit(commitToken *C.char) *C.char {
	client := CreateClient()
	var sendCommitRequest SendCommitRequest
	err := json.Unmarshal([]byte(C.GoString(commitToken)), &sendCommitRequest)
	response, err := client.Commit(context.Background(), &humanize_protobuf.CommitRequest{
		SessionId:  sendCommitRequest.SessionId,
		CommitType: humanize_protobuf.CommitRequest_COMMIT_TYPE_ACCEPT,
		MessageId:  sendCommitRequest.CommitToken,
	})
	if err != nil {
		return C.CString("")
	}
	if response.Outcome != humanize_protobuf.CommitResponse_COMMIT_OUTCOME_SUCCESSFUL {
		return C.CString("")
	}
	if response.MessageResponseData == nil {
		return C.CString("")
	}
	// Extract the emotional state
	responseData := &SendCommitResponse{
		EmotionalState: EmotionalState{
			SadHappy: int(response.MessageResponseData.
				ResponderEmotionalStateUpdate.
				EmotionalBounds["sad_happy"].CurrentPercentage),
			AgitatedPatient: int(response.MessageResponseData.
				ResponderEmotionalStateUpdate.
				EmotionalBounds["agitated_patient"].CurrentPercentage),
			BurnedOutAmbitious: int(response.MessageResponseData.
				ResponderEmotionalStateUpdate.
				EmotionalBounds["burned_out_ambitious"].CurrentPercentage),
			HateLove: int(response.MessageResponseData.
				ResponderEmotionalStateUpdate.
				EmotionalBounds["hate_love"].CurrentPercentage),
		},
	}
	marshalledString, err := json.Marshal(responseData)
	if err != nil {
		return C.CString("")
	}
	return C.CString(string(marshalledString))
}

func main() {
	startSession("Tom")
	print(sendMessage("Hello, who are you?"))
}
