package db

import (
	"fmt"

	"gorm.io/gorm"
)

func (h *HumanizeDbImpl) CreateEntity(
	id, name, sessionId,
	personalityId,
	generationConfigId, promptSetId,
	promptSegmentSetId, promptId,
	actuationSetId, automationPromptSetId,
	automationPromptSetSegmentId string,
) bool {
	entity := &Entity{
		EntityId:                     id,
		Name:                         name,
		SessionId:                    sessionId,
		PersonalityId:                personalityId,
		GenerationConfigId:           generationConfigId,
		PromptSetId:                  promptSetId,
		PromptId:                     promptId,
		PromptSegmentSetId:           promptSegmentSetId,
		ActuationRuleSetId:           actuationSetId,
		AutonomousPromptSegmentSetId: automationPromptSetSegmentId,
		AutonomousPromptSetId:        automationPromptSetId,
	}
	_ = h.mainDB.Create(entity)
	return true
}

func (h *HumanizeDbImpl) GetEntity(entityId string) (*Entity, error) {
	result := &Entity{}
	err := h.mainDB.
		Where("entity_id = ?", entityId).
		First(&result).
		Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Println(err.Error())
		}
		return nil, err
	}

	return result, nil
}

func (h *HumanizeDbImpl) UpdateEntityNeeds(
	entityId string,
	needs []string,
) (*Entity, error) {
	return nil, nil

}
