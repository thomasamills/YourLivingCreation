package game_loop

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"math/rand"
	"strings"
	"sync"
	"testserver/db"
	"testserver/id_gen"
	"testserver/memory_manager"
	npcdata "testserver/npcs"
	"testserver/rule_system"
	humanize_protobuf "testserver/src/generated/humanize-protobuf"
	"time"
)

type GameLoopManager interface {
	StartGameLoop(startNarrative bool) error
	SendMessage(req *humanize_protobuf.SendTextRequest,
	) (*humanize_protobuf.GetConversationInformationResponse, error)
	RetrieveInformation(
		request *humanize_protobuf.GetConversationInformationRequest,
	) (*humanize_protobuf.GetConversationInformationResponse, error)
	StopGameLoop()
	LoadNpcs() error
	KeepAlive()
	IsAlive(duration time.Duration) bool
	Commit(commitToken string) (*humanize_protobuf.MessageResponseData, error)
}

type GameLoopManagerImpl struct {
	db                     db.HumanizeDB
	memoryManager          memory_manager.MemoryManager
	emotionalStateManager  rule_system.EmotionalStateManager
	promptCreator          rule_system.PromptRuleSystemManager
	session                *db.Session
	reqChan                chan *npcdata.ActionRequest
	respChan               chan *humanize_protobuf.GetConversationInformationResponse
	stopChan               chan bool
	errorChan              chan error
	responseQueue          ResponseQueue
	currentNpcs            []*npcdata.NpcData
	conversationEscalation int
	mu                     sync.Mutex
	lastActiveTime         time.Time
	idGen                  id_gen.ULIDGenerator
	stateUpdateCallback    map[string]func() (*humanize_protobuf.MessageResponseData, error)
	chatGprClient          ChatGptClient
}

func (g *GameLoopManagerImpl) Commit(commitToken string) (*humanize_protobuf.MessageResponseData, error) {
	if x, found := g.stateUpdateCallback[commitToken]; found {
		return x()
	} else {
		return nil, errors.New("could not get the state commit loop")
	}
}

func (g *GameLoopManagerImpl) IsAlive(duration time.Duration) bool {
	if time.Since(g.lastActiveTime) <= duration {
		return true
	}
	return false
}

func (g *GameLoopManagerImpl) KeepAlive() {
	g.lastActiveTime = time.Now()
}

func NewGameLoopManager(
	db db.HumanizeDB,
	manager memory_manager.MemoryManager,
	stateManager rule_system.EmotionalStateManager,
	promptCreator rule_system.PromptRuleSystemManager,
	session *db.Session,
	chatGptClient ChatGptClient,
) GameLoopManager {
	return &GameLoopManagerImpl{
		mu:                    sync.Mutex{},
		db:                    db,
		memoryManager:         manager,
		emotionalStateManager: stateManager,
		promptCreator:         promptCreator,
		session:               session,
		responseQueue:         CreateResponseQueue(100),
		stopChan:              make(chan bool, 0),
		errorChan:             make(chan error, 0),
		currentNpcs:           make([]*npcdata.NpcData, 0),
		lastActiveTime:        time.Now(),
		idGen:                 id_gen.NewULIDGenerator(),
		stateUpdateCallback:   make(map[string]func() (*humanize_protobuf.MessageResponseData, error), 0),
		chatGprClient:         chatGptClient,
	}
}

func (g *GameLoopManagerImpl) LoadNpcs() error {
	g.mu.Lock()
	defer g.mu.Unlock()
	for _, entityId := range strings.Split(g.session.NpcEntityIds, ",") {
		entity, err := g.db.GetEntity(entityId)
		if err != nil {
			return err
		}
		npcEmotionalState, err := g.emotionalStateManager.GetEmotionalState(entityId)
		if err != nil {
			return err
		}
		genConfig, err := g.db.GetGenerationConfig(entity.GenerationConfigId)
		if err != nil {
			return err
		}
		//actuationRuleSet, err := g.db.GetActuationRuleSet(entity.ActuationRuleSetId, nil)
		//if err != nil {
		//	return err
		//}
		personality, err := g.db.GetPersonality(entity.PersonalityId, nil)
		if err != nil {
			return err
		}
		// Todo should be cached (memory of each npc at least 3 turns back)
		npcData := &npcdata.NpcData{
			EmotionalState:    npcEmotionalState,
			GptConfig:         genConfig,
			Personality:       personality,
			LastInputTime:     time.Now(),
			NpcRequestChannel: make(chan *npcdata.ActionRequest, 0),
			NpcStopChannel:    make(chan bool, 0),
			Entity:            entity,
			IsPaused:          false,
		}
		g.currentNpcs = append(g.currentNpcs, npcData)
	}
	return nil
}

func (g *GameLoopManagerImpl) StartGameLoop(startNarrative bool) error {
	// The main game loop process.
	start := func(
		stopChan chan bool,
		errorChan chan error,
	) {
		startNpcThoughtProcess := func(npc *npcdata.NpcData) {
			for {
				select {
				case stop := <-npc.NpcStopChannel:
					fmt.Println(stop)
					return
				default:
					select {
					case <-npc.NpcRequestChannel:
						break
					default:
						//TODO plug in every 5 minutes query knowledge graph and calculate current needs primer
						//TODO take current most satisfactory path and create a dialog flow intent for it

						time.Sleep(time.Duration(rand.Int31n(30)) * time.Second)
						//TODO intercept the response from gpt, send it to dialog flow to capture if the need was met.
						//TODO if need was met, then re=qurey knowledge graph and find new need.

						if startNarrative {
							// In the future start this to override any actual dialog.
						} else {
							if err := g.npcSpeakWithTarget(npc); err != nil {
								logrus.Error("npc could not auto speak")
							}
						}

						break
					}
				}
			}
		}

		for _, npc := range g.currentNpcs {
			npc := npc
			go func() {
				startNpcThoughtProcess(npc)
			}()
		}
	}

	// Process will keep running after function scope ends. To be closed via sending a bool to stopChan.
	go func() {
		start(g.stopChan, g.errorChan)
	}()

	return nil
}

func (g *GameLoopManagerImpl) ProcessResult(
	sessionId string,
	askerNpc *npcdata.NpcData,
	responderNpc *npcdata.NpcData,
	responderResponse *humanize_protobuf.HumanizeResponse,
	askerResponse *humanize_protobuf.HumanizeResponse,
	addToQueue bool,
	askForCommit bool,
) (*humanize_protobuf.MessageResponseData, error) {
	// First run actuation rule set.

	// Next run emotional rules and get updated state.
	processEmotionalRules := func(
		data *npcdata.NpcData, response *humanize_protobuf.HumanizeResponse,
	) (*humanize_protobuf.EmotionalState, string, error) {
		// Next process the state update based upon what has happened
		emotionRuleTriggers := make([]string, 0)
		if len(response.NpcTextEmotion) > 0 {
			emotionRuleTriggers = append(emotionRuleTriggers, "NPC_EMOTION:"+response.NpcTextEmotion)
		}
		if response.FearTriggered == true {
			emotionRuleTriggers = append(emotionRuleTriggers, "FEAR_TRIGGERED")
		}
		if response.HobbiesTriggered == true {
			emotionRuleTriggers = append(emotionRuleTriggers, "HOBBIES_TRIGGERED")
		}
		updatedState, triggeredActions, err := g.emotionalStateManager.ProcessEmotionalStateRules(
			*data, emotionRuleTriggers,
		)
		if err != nil {
			return nil, "", err
		}
		triggeredAction := ""
		if len(triggeredActions) > 0 {
			triggeredAction = triggeredActions[0]
		}
		return updatedState, triggeredAction, nil
	}

	var messageResponseData *humanize_protobuf.MessageResponseData
	var updatedAskerState *humanize_protobuf.EmotionalState = nil
	var updatedResponderState *humanize_protobuf.EmotionalState = nil
	updateStateId := g.idGen.GetULIDfromtime()
	getMessageResponseData := func() *humanize_protobuf.MessageResponseData {
		if askerNpc == nil {
			//meaning that the asker is the player...
			messageResponseData = &humanize_protobuf.MessageResponseData{
				AskerName:                     g.session.SpeakerName,
				AskerId:                       g.session.SpeakerEntityId,
				Message:                       askerResponse.Message,
				AskerEmotionalStateUpdate:     nil,
				AskerEmotion:                  askerResponse.NpcTextEmotion,
				Response:                      responderResponse.Response,
				ResponderName:                 responderNpc.Entity.Name,
				ResponderId:                   responderNpc.Entity.EntityId,
				ResponderEmotionalStateUpdate: updatedResponderState,
				ResponderTriggeredEmotion:     responderResponse.NpcTextEmotion,
				CommitToken:                   updateStateId,
			}
		} else if responderResponse == nil {
			messageResponseData = &humanize_protobuf.MessageResponseData{
				AskerName:                 askerNpc.Entity.Name,
				AskerId:                   askerNpc.Entity.EntityId,
				Message:                   askerResponse.Message,
				AskerEmotionalStateUpdate: updatedAskerState,
				AskerEmotion:              askerResponse.NpcTextEmotion,
				CommitToken:               updateStateId,
			}
		} else {
			messageResponseData = &humanize_protobuf.MessageResponseData{
				AskerName:                     askerNpc.Entity.Name,
				AskerId:                       askerNpc.Entity.EntityId,
				Message:                       askerResponse.Message,
				AskerEmotionalStateUpdate:     updatedAskerState,
				AskerEmotion:                  askerResponse.NpcTextEmotion,
				Response:                      responderResponse.Response,
				ResponderName:                 responderNpc.Entity.Name,
				ResponderId:                   responderNpc.Entity.EntityId,
				ResponderEmotionalStateUpdate: updatedResponderState,
				CommitToken:                   updateStateId,
			}
		}
		return messageResponseData
	}

	// Saves message and updates the emotional rules.
	commitDuration := time.Second * 20
	timeStateShouldHaveBeenDone := time.Now()
	updateState := func() (*humanize_protobuf.MessageResponseData, error) {
		if time.Since(timeStateShouldHaveBeenDone) > commitDuration {
			return nil, errors.New("commit time out")
		}
		var err error
		if askerResponse != nil && addToQueue && askerNpc != nil {
			updatedAskerState, _, err = processEmotionalRules(askerNpc, askerResponse)
			if err != nil || updatedAskerState == nil {
				return nil, err
			}
		}
		if responderResponse != nil {
			updatedResponderState, _, err = processEmotionalRules(responderNpc, responderResponse)
			if err != nil || updatedResponderState == nil {
				return nil, err
			}
		}
		if askerNpc == nil {
			//meaning that the asker is the player...
			err = g.memoryManager.CreateMessage(
				sessionId,
				g.session.SpeakerName,
				g.session.SpeakerEntityId,
				responderNpc.Entity.Name,
				responderNpc.Entity.EntityId,
				askerResponse.Message,
				responderResponse.Response,
				askerResponse.NpcTextEmotion,
				responderResponse.NpcTextEmotion,
			)
		} else if responderResponse == nil {
			err = g.memoryManager.CreateMessage(
				sessionId,
				askerNpc.Entity.Name,
				askerNpc.Entity.EntityId,
				"",
				"",
				askerResponse.Message,
				"",
				askerResponse.NpcTextEmotion,
				"",
			)
		} else {
			err = g.memoryManager.CreateMessage(
				sessionId,
				askerNpc.Entity.Name,
				askerNpc.Entity.EntityId,
				responderNpc.Entity.Name,
				responderNpc.Entity.EntityId,
				askerResponse.Message,
				responderResponse.Response,
				askerResponse.NpcTextEmotion,
				responderResponse.NpcTextEmotion,
			)
		}

		return getMessageResponseData(), nil
	}

	if askForCommit {
		g.stateUpdateCallback[updateStateId] = updateState
	} else {
		_, err := updateState()
		if err != nil {
			return nil, err
		}
	}

	getMessageResponseData()

	if addToQueue {
		err := g.responseQueue.Insert(messageResponseData)
		if err != nil {
			return nil, err
		}
	}
	return messageResponseData, nil
}

func (g *GameLoopManagerImpl) findNpcById(entityId string) *npcdata.NpcData {
	if g.currentNpcs != nil {
		for _, npc := range g.currentNpcs {
			if npc.Entity.EntityId == entityId {
				return npc
			}
		}
	}
	return nil
}

func (g *GameLoopManagerImpl) SendMessage(
	req *humanize_protobuf.SendTextRequest,
) (*humanize_protobuf.GetConversationInformationResponse, error) {
	// First obtain the target npc information
	g.KeepAlive()
	npcInformation := g.findNpcById(req.DirectedEntityId)
	finalMessageThatIsSent := req.Message
	if req.IsAction {
		mappedAction, err := g.mapActionText(
			req.Message,
			g.session.SpeakerName,
			npcInformation.Entity.Name,
		)
		finalMessageThatIsSent = mappedAction
		if err != nil {
			errMessage := "invalid action"
			return &humanize_protobuf.GetConversationInformationResponse{
				Status:       humanize_protobuf.GetConversationInformationResponse_FAILED,
				ErrorMessage: errMessage,
			}, nil
		}
	}
	// Now construct the prompt and send the text
	memLog, err := g.memoryManager.GetConversationMemoryLog(
		g.session.SessionID,
		int(npcInformation.GptConfig.MaxMemoryLogEntries),
		finalMessageThatIsSent, true, true,
	)
	if err != nil {
		errMessage := "could not get the conversation memory log"
		return &humanize_protobuf.GetConversationInformationResponse{
			Status:       humanize_protobuf.GetConversationInformationResponse_FAILED,
			ErrorMessage: errMessage,
		}, nil
	}
	// Generate the prompt, here we decide if the npc is to
	prompt, err := g.promptCreator.GenerateCharacterPrompt(
		*g.session, memLog, g.session.SpeakerName, npcInformation.Entity.Name, npcInformation,
	)
	if err != nil {
		errMessage := "could not create the character prompt"
		return &humanize_protobuf.GetConversationInformationResponse{
			Status:       humanize_protobuf.GetConversationInformationResponse_FAILED,
			ErrorMessage: errMessage,
		}, nil
	}
	response, err := g.chatGprClient.SendPrompt(
		npcInformation.Entity.Name,
		g.session.SpeakerName,
		prompt.PromptText,
		humanize_protobuf.HumanizeRequest_REQUEST_TYPE_SEND_MESSAGE_TO_NPC,
		req.Message,
	)
	if err != nil {
		errMessage := "could not get a response from gpt"
		return &humanize_protobuf.GetConversationInformationResponse{
			Status:       humanize_protobuf.GetConversationInformationResponse_FAILED,
			ErrorMessage: errMessage,
		}, nil
	}
	response.Message = req.Message
	addToQueue := true
	if !g.session.IsAsyncSession {
		addToQueue = false
	}
	askerResponse := &humanize_protobuf.HumanizeResponse{
		Message: req.Message,
	}
	messageResponseData, err := g.ProcessResult(g.session.SessionID, nil,
		npcInformation, response, askerResponse, addToQueue, g.session.WaitForCommit)
	if err != nil {
		errMessage := "could not process the result"
		return &humanize_protobuf.GetConversationInformationResponse{
			Status:       humanize_protobuf.GetConversationInformationResponse_FAILED,
			ErrorMessage: errMessage,
		}, nil
	}
	if !g.session.IsAsyncSession {
		syncResponse := &humanize_protobuf.GetConversationInformationResponse{
			Status:              humanize_protobuf.GetConversationInformationResponse_SUCCESS,
			MessageResponseData: make(map[string]*humanize_protobuf.MessageResponseData, 0),
		}
		syncResponse.MessageResponseData[time.Now().String()] = messageResponseData
		return syncResponse, nil
	}
	return &humanize_protobuf.GetConversationInformationResponse{
		Status: humanize_protobuf.GetConversationInformationResponse_SUCCESS,
	}, nil
}

func (g *GameLoopManagerImpl) RetrieveInformation(
	request *humanize_protobuf.GetConversationInformationRequest,
) (*humanize_protobuf.GetConversationInformationResponse, error) {
	if g.responseQueue.IsEmpty() {
		return &humanize_protobuf.GetConversationInformationResponse{
			Status:              humanize_protobuf.GetConversationInformationResponse_SUCCESS,
			MessageResponseData: nil,
		}, nil
	}
	resp := &humanize_protobuf.GetConversationInformationResponse{
		MessageResponseData: make(map[string]*humanize_protobuf.MessageResponseData, 0),
	}
	for {
		if g.responseQueue.IsEmpty() {
			break
		}
		listItem, err := g.responseQueue.Remove()
		if err != nil {
			return &humanize_protobuf.GetConversationInformationResponse{
				Status:              humanize_protobuf.GetConversationInformationResponse_FAILED,
				ErrorMessage:        "couldn't remove response from the queue ",
				MessageResponseData: nil,
			}, err
		}
		resp.MessageResponseData[time.Now().String()] = listItem
	}

	return resp, nil
}

func (g *GameLoopManagerImpl) StopGameLoop() {
	for _, npc := range g.currentNpcs {
		select {
		case npc.NpcStopChannel <- true:
			break
		}
	}
}

func (g *GameLoopManagerImpl) mapActionText(
	action, speakerName, charName string,
) (string, error) {
	switch action {
	case "KICK":
		return "*" + speakerName + " kicks " + charName + "*", nil
	}
	return "", errors.New("invalid action")
}

func (g *GameLoopManagerImpl) npcAskAQuestionWithDecisionPrompt(npc *npcdata.NpcData) error {
	return g.npcSpeakWithPlayer(npc)
}

func (g *GameLoopManagerImpl) npcSpeakWithPlayer(npc *npcdata.NpcData) error {
	// interjection <- effects just their emotional state
	memLog, err := g.memoryManager.GetConversationMemoryLog(
		g.session.SessionID,
		int(npc.GptConfig.MaxMemoryLogEntries),
		"", false, true,
	)
	if err != nil {
		return err
	}
	// Generate the prompt, here we decide if the npc is
	autonomousPrompt, err := g.promptCreator.GenerateCharacterPrompt(
		*g.session, memLog, g.session.SpeakerName, npc.Entity.Name, npc)
	if err != nil {
		return err
	}
	// Send with python client and process the result
	response, err := g.chatGprClient.SendPrompt(
		npc.Entity.Name,
		g.session.SpeakerName,
		autonomousPrompt.PromptText,
		humanize_protobuf.HumanizeRequest_REQUEST_TYPE_AUTONOMOUS,
		"",
	)
	if err != nil {
		return err
	}
	response.Message = response.Response // as it is aunto
	if err != nil {
		return err
	}
	_, err = g.ProcessResult(
		g.session.SessionID, npc, nil,
		nil, response, true, g.session.WaitForCommit,
	)
	if err != nil {
		return err
	}

	return nil
}

func (g *GameLoopManagerImpl) npcSpeakWithTarget(npc *npcdata.NpcData) error {
	// target a random npc.
	if len(g.currentNpcs) < 2 {
		return g.npcSpeakWithPlayer(npc)
	}
	selectedNpc := g.currentNpcs[rand.Int31n(int32(len(g.currentNpcs)))]
	if selectedNpc.Entity.EntityId == npc.Entity.EntityId {
		for _, otherNpc := range g.currentNpcs {
			if otherNpc.Entity.EntityId != npc.Entity.EntityId {
				selectedNpc = otherNpc
			}
		}
	}
	// directed interjection <- effects npc and targeted state.
	memLog, err := g.memoryManager.GetConversationMemoryLog(
		g.session.SessionID,
		int(npc.GptConfig.MaxMemoryLogEntries),
		"", false, true,
	)
	if err != nil {
		return err
	}
	// Generate the prompt, here we decide if the npc is to
	autonomousPrompt, err := g.promptCreator.GenerateCharacterPrompt(
		*g.session, memLog, selectedNpc.Entity.Name, npc.Entity.Name, npc)
	if err != nil {
		return err
	}
	// Send with python client and process the result
	askerResponse, err := g.chatGprClient.SendPrompt(
		npc.Entity.Name,
		selectedNpc.Entity.Name,
		autonomousPrompt.PromptText,
		humanize_protobuf.HumanizeRequest_REQUEST_TYPE_AUTONOMOUS,
		"",
	)
	if err != nil {
		return err
	}
	// Reprocess the memory log.
	memLog, err = g.memoryManager.GetConversationMemoryLog(
		g.session.SessionID,
		int(npc.GptConfig.MaxMemoryLogEntries),
		askerResponse.Response, false, true,
	)
	if err != nil {
		return err
	}
	// generate the question prompt to the targeted npc
	questionPrompt, err := g.promptCreator.GenerateCharacterPrompt(
		*g.session, memLog, npc.Entity.Name, selectedNpc.Entity.Name, selectedNpc)
	if err != nil {
		return err
	}
	// Send with python client and process the result
	response, err := g.chatGprClient.SendPrompt(
		selectedNpc.Entity.Name,
		npc.Entity.Name,
		questionPrompt.PromptText,
		humanize_protobuf.HumanizeRequest_REQUEST_TYPE_SEND_MESSAGE_TO_NPC,
		askerResponse.Response,
	)
	askerResponse.Message = askerResponse.Response
	response.Message = askerResponse.Response
	if err != nil {
		return err
	}
	_, err = g.ProcessResult(
		g.session.SessionID, npc, selectedNpc,
		response, askerResponse, true, g.session.WaitForCommit,
	)
	if err != nil {
		return err
	}
	return nil
}
