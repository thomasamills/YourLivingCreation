package conversation_code

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"testserver/conversation_code/rule_system"
	"testserver/db"
	"testserver/db/id_gen"
	humanize_protobuf "testserver/src/generated/humanize-protobuf"
)

type ConnectionManager interface {
	Connect(
		req *humanize_protobuf.MiddleWareConnectReq,
	) (*humanize_protobuf.MiddleWareConnectResp, error)
	SendMessage(req *humanize_protobuf.SendTextRequest,
	) (*humanize_protobuf.GetConversationInformationResponse, error)
	RetrieveInformation(
		request *humanize_protobuf.GetConversationInformationRequest,
	) (*humanize_protobuf.GetConversationInformationResponse, error)
	Commit(req *humanize_protobuf.CommitRequest) (*humanize_protobuf.CommitResponse, error)
}

type ConnectionManagerImpl struct {
	client                humanize_protobuf.HumanizeClient
	emotionalStateManager rule_system.EmotionalStateManager
	db                    db.HumanizeDB
	idGen                 id_gen.ULIDGenerator
	promptCreation        rule_system.PromptRuleSystemManager
	memoryManager         MemoryManager
	gameLoopRegistry      GameLoopManagerRegistry
	chatGptClient         ChatGptClient
}

func NewConnectionManager(
	db db.HumanizeDB,
	emotionalStateManager rule_system.EmotionalStateManager,
	idGen id_gen.ULIDGenerator,
	promptCreation rule_system.PromptRuleSystemManager,
	memManager MemoryManager,
	gameLoopRegistry GameLoopManagerRegistry,
	chatGptClient ChatGptClient,
) (ConnectionManager, error) {
	go func() {
		if err := gameLoopRegistry.PurgeGameLoops(); err != nil {
			logrus.Fatal(err)
		}
	}()
	client := &ConnectionManagerImpl{
		db:                    db,
		emotionalStateManager: emotionalStateManager,
		idGen:                 idGen,
		promptCreation:        promptCreation,
		memoryManager:         memManager,
		gameLoopRegistry:      gameLoopRegistry,
		chatGptClient:         chatGptClient,
	}
	return client, nil
}

func (cm *ConnectionManagerImpl) Connect(
	req *humanize_protobuf.MiddleWareConnectReq,
) (*humanize_protobuf.MiddleWareConnectResp, error) {
	// first make sure we have ids for both entities
	resp := &humanize_protobuf.MiddleWareConnectResp{
		NpcNameToId: make(map[string]string, 0),
	}
	failedMessage := "could not create new entity id"
	speakerNameId := cm.idGen.GetULIDfromtime()
	validationStatus, errMessage := cm.validateConnectResp(req)
	if validationStatus != humanize_protobuf.MiddleWareConnectResp_UNKNOWN {
		logrus.Error("Invalid connect request")
		resp.Status = validationStatus
		resp.ErrorMessage = errMessage
		return resp, errors.New("failed to connect")
	}
	convBetweenString := "Conversation attempting to start between"
	for i, npc := range req.NpcInformation {
		if i == 0 {
			convBetweenString += " " + npc.Name
		} else {
			convBetweenString += ", " + npc.Name
		}
	}
	logrus.Info(convBetweenString)

	sessionId := cm.idGen.GetULIDfromtime()
	resp.SpeakerNameId = speakerNameId
	resp.SessionId = sessionId
	resp.NpcEmotional_States = make(map[string]*humanize_protobuf.EmotionalState, 0)
	npcEntityIds := make([]string, 0)

	createEntity := func(
		npcInformation *humanize_protobuf.NpcRequestInformation,
	) (*humanize_protobuf.MiddleWareConnectResp, error) {
		id := cm.idGen.GetULIDfromtime()
		success := cm.db.CreateEntity(
			id, npcInformation.Name,
			sessionId, npcInformation.PersonalityId,
			npcInformation.GenConfigId, npcInformation.PromptSetId,
			npcInformation.PromptSegmentSetId, npcInformation.DefaultPromptId,
			npcInformation.ActuationRuleSetId, npcInformation.ActuationPromptSetId,
			npcInformation.ActuationPromptSegmentSetId,
		)

		if success != true {
			logrus.WithFields(logrus.Fields{
				"session_id":  sessionId,
				"entity_name": npcInformation.Name,
			}).Error("Failed creating a new entity")
			resp.Status = humanize_protobuf.MiddleWareConnectResp_FAILED
			resp.ErrorMessage = failedMessage
			return resp, errors.New(failedMessage)
		}

		resp.NpcNameToId[npcInformation.Name] = id

		logrus.WithFields(logrus.Fields{
			"session_id":      sessionId,
			"entity_name":     npcInformation.Name,
			"preset_state_id": npcInformation.PresetEmotionalStateId,
		}).Info("Copying emotional state for npc (entity 2)")
		emotionalState, err := cm.db.CopyEmotionalState(npcInformation.PresetEmotionalStateId, id)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"session_id":      sessionId,
				"entity_name":     npcInformation.Name,
				"preset_state_id": id,
			}).Error("Failed copying emotional state for npc (entity 2)")
			resp.Status = humanize_protobuf.MiddleWareConnectResp_FAILED
			resp.ErrorMessage = "failed to load initial preset emotional state"
			return resp, errors.New(resp.ErrorMessage)
		}
		resp.NpcEmotional_States[id] = emotionalState
		logrus.WithFields(logrus.Fields{
			"session_id":      sessionId,
			"entity_name":     npcInformation.Name,
			"preset_state_id": npcInformation.PresetEmotionalStateId,
		}).Info("Loaded entity")
		npcEntityIds = append(npcEntityIds, id)
		return nil, nil
	}

	for _, npcInfo := range req.NpcInformation {
		resp, err := createEntity(npcInfo)
		if err != nil {
			return resp, err
		}
	}

	session, err := cm.db.CreateSession(
		sessionId, req.SpeakerName,
		speakerNameId, npcEntityIds,
		req.StartAsyncGameLoop,
		req.WaitForCommitMessageBeforeUpdatingState,
		req.StartNarrative,
	)
	if err != nil {
		resp.Status = humanize_protobuf.MiddleWareConnectResp_FAILED
		errMessage := "failed to create session"
		logrus.Error(errMessage)
		resp.ErrorMessage = errMessage
		return resp, err
	}

	// Initialize game loop
	gameLoop := NewGameLoopManager(
		cm.db,
		cm.memoryManager,
		cm.emotionalStateManager,
		cm.promptCreation,
		session,
		cm.chatGptClient,
	)

	err = gameLoop.LoadNpcs()
	if err != nil {
		resp.Status = humanize_protobuf.MiddleWareConnectResp_FAILED
		errMessage := "failed to load the npc data"
		logrus.Error(errMessage)
		resp.ErrorMessage = errMessage
		return resp, err
	}

	if req.StartAsyncGameLoop {

		err = gameLoop.StartGameLoop(req.StartNarrative)
		if err != nil {
			resp.Status = humanize_protobuf.MiddleWareConnectResp_FAILED
			errMessage := "failed to start game loop"
			logrus.Error(errMessage)
			resp.ErrorMessage = errMessage
			return resp, err
		}
	}

	err = cm.gameLoopRegistry.AddGameLoop(sessionId, gameLoop)
	if err != nil {
		resp.Status = humanize_protobuf.MiddleWareConnectResp_FAILED
		errMessage := "failed to add game loop to the registry"
		logrus.Error(errMessage)
		resp.ErrorMessage = errMessage
		gameLoop.StopGameLoop()
		return resp, err
	}

	resp.Status = humanize_protobuf.MiddleWareConnectResp_CONNECTED
	return resp, nil
}

func (cm *ConnectionManagerImpl) SendMessage(
	req *humanize_protobuf.SendTextRequest,
) (*humanize_protobuf.GetConversationInformationResponse, error) {
	gameLoop, err := cm.gameLoopRegistry.GetGameLoop(req.SessionId)
	if err != nil {
		return &humanize_protobuf.GetConversationInformationResponse{
			Status:       humanize_protobuf.GetConversationInformationResponse_FAILED,
			ErrorMessage: "Could not connect to the game loop, could not retrieve game loop from registry",
		}, err
	}
	return gameLoop.SendMessage(req)
}

func (cm *ConnectionManagerImpl) RetrieveInformation(
	request *humanize_protobuf.GetConversationInformationRequest,
) (*humanize_protobuf.GetConversationInformationResponse, error) {
	gameLoop, err := cm.gameLoopRegistry.GetGameLoop(request.SessionId)
	if err != nil {
		return &humanize_protobuf.GetConversationInformationResponse{
			Status:       humanize_protobuf.GetConversationInformationResponse_FAILED,
			ErrorMessage: "Could not connect to the game loop, could not retrieve game loop from registry",
		}, err
	}
	return gameLoop.RetrieveInformation(request)
}

func (cm *ConnectionManagerImpl) validateConnectResp(
	req *humanize_protobuf.MiddleWareConnectReq,
) (humanize_protobuf.MiddleWareConnectResp_ConnectionStatus, string) {
	if len(req.SpeakerName) == 0 {
		return humanize_protobuf.MiddleWareConnectResp_FAILED, "No speaker name specified"
	}
	if len(req.NpcInformation) == 0 {
		return humanize_protobuf.MiddleWareConnectResp_FAILED, "No npc information"
	}
	// validate individual npc information
	for _, npc := range req.NpcInformation {
		if len(npc.ActuationRuleSetId) == 0 {
			return humanize_protobuf.MiddleWareConnectResp_FAILED,
				fmt.Sprintf("No actuation rule set id for %s", npc.Name)
		}
		if len(npc.PromptSetId) == 0 {
			return humanize_protobuf.MiddleWareConnectResp_FAILED,
				fmt.Sprintf("No prompt set id for %s", npc.Name)
		}
		if len(npc.PromptSegmentSetId) == 0 {
			return humanize_protobuf.MiddleWareConnectResp_FAILED,
				fmt.Sprintf("No prompt segment set id for %s", npc.Name)
		}
		if len(npc.DefaultPromptId) == 0 {
			return humanize_protobuf.MiddleWareConnectResp_FAILED,
				fmt.Sprintf("No default prompt id for %s", npc.Name)
		}
		if len(npc.PersonalityId) == 0 {
			return humanize_protobuf.MiddleWareConnectResp_FAILED,
				fmt.Sprintf("No personality id for %s", npc.Name)
		}
		if len(npc.PresetEmotionalStateId) == 0 {
			return humanize_protobuf.MiddleWareConnectResp_FAILED,
				fmt.Sprintf("No preset emotional state id for %s", npc.Name)
		}
		if len(npc.GenConfigId) == 0 {
			return humanize_protobuf.MiddleWareConnectResp_FAILED,
				fmt.Sprintf("No generation config for %s", npc.Name)
		}
	}
	return humanize_protobuf.MiddleWareConnectResp_UNKNOWN, ""
}

func (cm *ConnectionManagerImpl) Commit(req *humanize_protobuf.CommitRequest) (*humanize_protobuf.CommitResponse, error) {
	gameLoop, err := cm.gameLoopRegistry.GetGameLoop(req.SessionId)
	if err != nil {
		return &humanize_protobuf.CommitResponse{
			Outcome: humanize_protobuf.CommitResponse_COMMIT_OUTCOME_FAILED,
		}, err
	}
	messageResponseData, err := gameLoop.Commit(req.MessageId)
	if err != nil {
		return &humanize_protobuf.CommitResponse{
			Outcome: humanize_protobuf.CommitResponse_COMMIT_OUTCOME_FAILED,
		}, err
	}
	return &humanize_protobuf.CommitResponse{
		Outcome:             humanize_protobuf.CommitResponse_COMMIT_OUTCOME_SUCCESSFUL,
		MessageResponseData: messageResponseData,
	}, nil
}
