package quest_generation

import (
	"gorm.io/gorm"
	"log"
	"testserver/conversation_code/npc_data"
	"testserver/conversation_code/rule_system"
	"testserver/db"
)

type QuestGenerator interface {
	GenerateQuest(npcData []npc_data.NpcData, session db.Session) error
}

type QuestGeneratorImpl struct {
	db            gorm.DB
	logger        *log.Logger
	promptCreator rule_system.PromptRuleSystemManager
}

type QuestSettings struct {
	NumScenarios            *string
	NumDecisionsPerScenario *string
	Topic                   *string
}

func NewQuestGenerator(
	db gorm.DB, logger *log.Logger,
	promptRuleSystemManager rule_system.PromptRuleSystemManager,
) QuestGenerator {
	return &QuestGeneratorImpl{
		db:            db,
		logger:        logger,
		promptCreator: promptRuleSystemManager,
	}
}

func (q QuestGeneratorImpl) GenerateQuest(npcData []npc_data.NpcData, session db.Session) error {
	// Function that combines npc data along with topic, story and environment (future)
	// It will construct a prompt outlying crossover of personalities backstory and topic
	// It will generate x amount of decisions for an npc and then parse that into a quest for the npc to autonomously play out
	//	promptTemplate := `You are a quest generator. Based upon the following topic: "{{.Topic}}" Generate
	//ne {{.NumScenarios}} different scenarios. Each scenario will have {{.NumDecisionsPerScenario}} possible outcomes.. $
	//
	//To help you, here are the following backgrounds fo the different characters: `
	//
	return nil
}
