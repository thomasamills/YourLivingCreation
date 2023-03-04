package apis

import (
	"context"
	"fmt"
	humanize_client "testserver/connection_manager"
	"testserver/db"
	"testserver/game_loop"
	"testserver/id_gen"
	"testserver/memory_manager"
	emotional_state_manager "testserver/rule_system"
	humanize_protobuf "testserver/src/generated/humanize-protobuf"
)

type MiddlewareApiImplementation struct {
	idGen                 id_gen.ULIDGenerator
	db                    db.HumanizeDB
	connectionManager     humanize_client.ConnectionManager
	emotionalStateManager emotional_state_manager.EmotionalStateManager
	humanize_protobuf.UnimplementedConnectServer
}

func NewMiddlewareApiImplementation(
	db db.HumanizeDB, generator id_gen.ULIDGenerator,
) (*MiddlewareApiImplementation, error) {
	emotionalStateManager := emotional_state_manager.NewEmotionalStateManager(db)
	promptCreation := emotional_state_manager.NewActuationRuleBasedSystemManager(db, emotionalStateManager)
	memoryManager := memory_manager.NewMemoryManager(db)
	gameLoopRegistry := humanize_client.NewGameLoopManagerRegistry()
	chatGptClient := game_loop.NewChatGptClient()
	conversationManager, err := humanize_client.NewConnectionManager(
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
