package conversation_code

import (
	"context"
	"fmt"
	"testserver/conversation_code/rule_system"
	"testserver/db"
	"testserver/db/id_gen"
	humanize_protobuf "testserver/src/generated/humanize-protobuf"
)

type MiddlewareApiImplementation struct {
	idGen                 id_gen.ULIDGenerator
	db                    db.HumanizeDB
	connectionManager     ConnectionManager
	emotionalStateManager rule_system.EmotionalStateManager
	humanize_protobuf.UnimplementedConnectServer
}

func NewMiddlewareApiImplementation(
	db db.HumanizeDB, generator id_gen.ULIDGenerator,
) (*MiddlewareApiImplementation, error) {
	emotionalStateManager := rule_system.NewEmotionalStateManager(db)
	promptCreation := rule_system.NewActuationRuleBasedSystemManager(db, emotionalStateManager)
	memoryManager := NewMemoryManager(db)
	gameLoopRegistry := NewGameLoopManagerRegistry()
	chatGptClient := NewChatGptClient()
	conversationManager, err := NewConnectionManager(
		db,
		emotionalStateManager,
		generator,
		promptCreation,
		memoryManager,
		gameLoopRegistry,
		chatGptClient,
	)
	if err != nil {
		fmt.Println("couldn't connect to python client")
		return nil, err
	}
	return &MiddlewareApiImplementation{
		db:                    db,
		connectionManager:     conversationManager,
		idGen:                 id_gen.NewULIDGenerator(),
		emotionalStateManager: emotionalStateManager,
	}, nil
}

func (m *MiddlewareApiImplementation) Connect(
	ctx context.Context, req *humanize_protobuf.MiddleWareConnectReq,
) (*humanize_protobuf.MiddleWareConnectResp, error) {
	return m.connectionManager.Connect(req)
}

func (m *MiddlewareApiImplementation) SendText(
	ctx context.Context, req *humanize_protobuf.SendTextRequest,
) (*humanize_protobuf.GetConversationInformationResponse, error) {
	return m.connectionManager.SendMessage(req)
}

func (m *MiddlewareApiImplementation) Commit(
	ctx context.Context, req *humanize_protobuf.CommitRequest,
) (*humanize_protobuf.CommitResponse, error) {
	return m.connectionManager.Commit(req)
}
func (m *MiddlewareApiImplementation) GetConversationInformation(
	ctx context.Context, req *humanize_protobuf.GetConversationInformationRequest,
) (*humanize_protobuf.GetConversationInformationResponse, error) {
	return m.connectionManager.RetrieveInformation(req)
}
